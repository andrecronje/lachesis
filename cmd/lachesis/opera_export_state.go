package main

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"path"
	"time"

	ointer "github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/drivertype"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesis/driver"
	"github.com/Fantom-foundation/go-opera/opera/genesis/driverauth"
	"github.com/Fantom-foundation/go-opera/opera/genesis/evmwriter"
	"github.com/Fantom-foundation/go-opera/opera/genesis/gpos"
	"github.com/Fantom-foundation/go-opera/opera/genesis/netinit"
	osfc "github.com/Fantom-foundation/go-opera/opera/genesis/sfc"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	futils "github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/lachesis-base/hash"
	oidx "github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb/leveldb"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	ethparams "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/gossip/sfc204"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/inter/sfctype"
	"github.com/Fantom-foundation/go-lachesis/lachesis"
	"github.com/Fantom-foundation/go-lachesis/lachesis/genesis/sfc"
	"github.com/Fantom-foundation/go-lachesis/lachesis/genesis/sfc/sfcpos"
	"github.com/Fantom-foundation/go-lachesis/poset"
)

var (
	emptyCodeHash = common.BytesToHash(crypto.Keccak256(nil))
	emptyRoot     = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
	emptyNodeHash = common.Hash{}
)

type evmCaller struct {
	statedb     *state.StateDB
	header      *evmcore.EvmHeader
	chainConfig *ethparams.ChainConfig
}

func (c *evmCaller) CodeAt(_ context.Context, contract common.Address, _ *big.Int) ([]byte, error) {
	return c.statedb.GetCode(contract), nil
}

func (c *evmCaller) CallContract(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	msg := types.NewMessage(call.From, call.To, 0, new(big.Int), 1e12, new(big.Int), call.Data, false)

	// Create a new context to be used in the EVM environment
	econtext := evmcore.NewEVMContext(msg, c.header, nil, nil)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(econtext, c.statedb, c.chainConfig, vm.Config{})

	gp := new(evmcore.GasPool).AddGas(math.MaxUint64)
	result, err := evmcore.ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, err
	}
	if result.Err != nil {
		return nil, fmt.Errorf("revert err: %v", result.Err)
	}
	return result.Return(), nil
}

func addDelegation(genStore *genesisstore.Store, addr common.Address, to oidx.ValidatorID, added genesis.Delegation) {
	delegation := genStore.GetDelegation(addr, to)
	delegation.Stake.Add(delegation.Stake, added.Stake)
	delegation.Rewards.Add(delegation.Rewards, added.Rewards)
	delegation.LockedStake.Add(delegation.LockedStake, added.LockedStake)
	delegation.EarlyUnlockPenalty.Add(delegation.EarlyUnlockPenalty, added.EarlyUnlockPenalty)
	if delegation.LockupFromEpoch == 0 {
		delegation.LockupFromEpoch = added.LockupFromEpoch
	}
	if delegation.LockupDuration == 0 {
		delegation.LockupDuration = added.LockupDuration
	}
	if delegation.LockupEndTime == 0 {
		delegation.LockupEndTime = added.LockupEndTime
	}
	genStore.SetDelegation(addr, to, delegation)
}

func addBalance(genStore *genesisstore.Store, statedb *state.StateDB, addr common.Address, added *big.Int) {
	acc := genStore.GetEvmAccount(addr)
	if acc.Balance.Sign() == 0 {
		acc.Balance = statedb.GetBalance(addr)
	}
	acc.Balance = new(big.Int).Add(acc.Balance, added)
	genStore.SetEvmAccount(addr, acc)
}

type withdrawalRequestID struct {
	from common.Address
	wrID string
}

type rawEvmItemsToEthdb struct {
	*genesisstore.Store
}

// Put inserts the given value into the key-value data store.
func (db *rawEvmItemsToEthdb) Put(key []byte, value []byte) error {
	db.Store.SetRawEvmItem(key, value)
	return nil
}

// Delete removes the key from the key-value data store.
func (db *rawEvmItemsToEthdb) Delete(key []byte) error {
	return errors.New("not supported")
}

func ExportState(path string, gdb *gossip.Store, cdb *poset.Store, net *lachesis.Config) (*genesisstore.Store, error) {
	_ = os.RemoveAll(path)

	err := os.MkdirAll(path, 0700)
	if err != nil {
		return nil, err
	}
	db, err := leveldb.New(path, 2*opt.MiB, 0, nil, nil)
	if err != nil {
		return nil, err
	}

	genStore := genesisstore.NewStore(db)

	start, reported := time.Now(), time.Time{}

	// export blocks, transactions and receipts
	var lastBlock *inter.Block
	gdb.ForEachBlock(func(index idx.Block, block *inter.Block) {
		lastBlock = block
		receipts := gdb.App().GetReceipts(index)
		receiptsStorage := make([]*types.ReceiptForStorage, len(receipts))
		for i, r := range receipts {
			receiptsStorage[i] = (*types.ReceiptForStorage)(r)
		}
		genStore.SetBlock(oidx.Block(index), genesis.Block{
			Time:        ointer.Timestamp(block.Time),
			Atropos:     hash.Event(block.Atropos),
			Txs:         gdb.GetBlockTransactions(block),
			InternalTxs: types.Transactions{},
			Root:        hash.Hash(block.Root),
			Receipts:    receiptsStorage,
		})
		if time.Since(reported) >= statsReportLimit {
			log.Info("Exporting blocks", "last", lastBlock.Index, "elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	})
	log.Info("Exported blocks", "last", lastBlock.Index, "elapsed", common.PrettyDuration(time.Since(start)))

	// export EVM state
	log.Info("Exporting EVM state", "root", lastBlock.Root.String())
	statedb, _ := gdb.App().StateDB(lastBlock.Root)

	stateTr, err := statedb.Database().OpenTrie(lastBlock.Root)
	if err != nil {
		return genStore, errors.Wrap(err, "failed to open EVM trie")
	}

	// export EVM tries
	destDb := &rawEvmItemsToEthdb{genStore}
	stateIt := stateTr.NodeIterator(nil)
	for stateIt.Next(true) {
		if stateIt.Leaf() {
			addrHash := common.BytesToHash(stateIt.LeafKey())

			var account state.Account
			if err := rlp.Decode(bytes.NewReader(stateIt.LeafBlob()), &account); err != nil {
				return genStore, errors.Wrap(err, "failed to decode account")
			}

			codeHash := common.BytesToHash(account.CodeHash)
			if codeHash != emptyCodeHash {
				code, err := statedb.Database().ContractCode(addrHash, codeHash)
				if err != nil {
					return genStore, errors.Wrap(err, "failed to open EVM trie")
				}
				rawdb.WriteCode(destDb, codeHash, code)
			}

			if account.Root != emptyRoot {
				dataTr, err := statedb.Database().OpenStorageTrie(addrHash, account.Root)
				if err != nil {
					return genStore, errors.Wrap(err, "failed to open EVM trie")
				}
				dataIt := dataTr.NodeIterator(nil)
				for dataIt.Next(true) {
					if dataIt.Leaf() || dataIt.Hash() == emptyNodeHash {
						continue
					}
					nodeBody, err := gdb.App().EvmTable().Get(dataIt.Hash().Bytes())
					if err != nil {
						return genStore, errors.Wrap(err, "failed to get storage trie node")
					}
					rawdb.WriteTrieNode(destDb, dataIt.Hash(), nodeBody)
				}
			}
		} else if stateIt.Hash() != emptyNodeHash {
			nodeBody, err := gdb.App().EvmTable().Get(stateIt.Hash().Bytes())
			if err != nil {
				return genStore, errors.Wrap(err, "failed to get trie node")
			}
			rawdb.WriteTrieNode(destDb, stateIt.Hash(), nodeBody)
		}
	}

	// SFC 3.x
	genStore.SetEvmAccount(osfc.ContractAddress, genesis.Account{
		Code:         osfc.GetContractBin(),
		Balance:      new(big.Int),
		Nonce:        0,
		SelfDestruct: true, // erase SFCv2 storage
	})
	// NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// EvmWriter
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})
	log.Info("Exported EVM state", "elapsed", common.PrettyDuration(time.Since(start)))

	// export network rules
	var rules = opera.TestNetRules()
	if net.NetworkID == lachesis.MainNetworkID {
		rules = opera.MainNetRules()
	} else {
		rules.NetworkID = net.NetworkID
		rules.Name = net.Name
	}
	genStore.SetRules(rules)

	// export metadata
	metadata := genesisstore.Metadata{}

	lastStakerID := idx.StakerID(statedb.GetState(sfc.ContractAddress, sfcpos.StakersLastID()).Big().Uint64())
	metadata.Validators = make(gpos.Validators, lastStakerID)

	caller := &evmCaller{
		statedb:     statedb,
		header:      evmcore.ToEvmHeader(lastBlock),
		chainConfig: net.EvmChainConfig(),
	}

	sfcCaller, err := sfc204.NewContractCaller(sfc.ContractAddress, caller)
	if err != nil {
		return genStore, err
	}

	callsOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: futils.U64toBig(uint64(lastBlock.Index)),
		Context:     context.TODO(),
	}

	// export validators
	log.Info("Exporting validators")

	epoch := cdb.GetEpoch().EpochN
	for stakerID := idx.StakerID(1); stakerID <= lastStakerID; stakerID++ {
		vpos := sfcpos.Staker(stakerID)
		status := statedb.GetState(sfc.ContractAddress, vpos.Status()).Big().Uint64()
		stake := statedb.GetState(sfc.ContractAddress, vpos.StakeAmount()).Big()
		createdTime := inter.Timestamp(statedb.GetState(sfc.ContractAddress, vpos.CreatedTime()).Big().Uint64() * uint64(time.Second))
		createdEpoch := idx.Epoch(statedb.GetState(sfc.ContractAddress, vpos.CreatedEpoch()).Big().Uint64())
		deactivatedTime := inter.Timestamp(statedb.GetState(sfc.ContractAddress, vpos.DeactivatedTime()).Big().Uint64() * uint64(time.Second))
		deactivatedEpoch := idx.Epoch(statedb.GetState(sfc.ContractAddress, vpos.DeactivatedEpoch()).Big().Uint64())
		addr := common.BytesToAddress(statedb.GetState(sfc.ContractAddress, vpos.Address()).Bytes()[12:])

		// retrieve rewards
		pendingRewards, _, _, _ := sfcCaller.CalcValidatorRewards(callsOpts, futils.U64toBig(uint64(stakerID)), futils.U64toBig(0), futils.U64toBig(uint64(epoch)))
		if pendingRewards == nil {
			pendingRewards = new(big.Int)
		}

		oDeactivatedEpoch := deactivatedEpoch
		oDeactivatedTime := deactivatedTime
		oStatus := status

		if stake.Sign() != 0 {
			if oDeactivatedEpoch != 0 && status != sfctype.ForkBit {
				// finalize withdrawal balance if validator isn't a cheater
				addBalance(genStore, statedb, addr, stake)
			} else {
				// retrieve lockup info
				info, _ := sfcCaller.LockedStakes(callsOpts, futils.U64toBig(uint64(stakerID)))
				if info.Duration == nil {
					info.Duration = new(big.Int)
				}
				if info.EndTime == nil {
					info.EndTime = new(big.Int)
				}
				if info.FromEpoch == nil {
					info.FromEpoch = new(big.Int)
				}
				earlyUnlockPenalty, _ := sfcCaller.CalcValidatorLockupPenalty(callsOpts, futils.U64toBig(uint64(stakerID)), stake)
				if earlyUnlockPenalty == nil {
					earlyUnlockPenalty = new(big.Int)
				}
				lockedStake := stake
				if info.Duration.Sign() == 0 {
					lockedStake = new(big.Int)
				}
				// add validator self-delegation
				addDelegation(genStore, addr, oidx.ValidatorID(stakerID), genesis.Delegation{
					Stake:              stake,
					Rewards:            pendingRewards,
					LockedStake:        lockedStake,
					LockupFromEpoch:    oidx.Epoch(info.FromEpoch.Uint64()),
					LockupEndTime:      ointer.Timestamp(info.EndTime.Uint64()),
					LockupDuration:     ointer.Timestamp(info.Duration.Uint64()),
					EarlyUnlockPenalty: earlyUnlockPenalty,
				})
			}
		} else {
			// recover creation time/epoch/address from logs and genesis data
			if v, ok := net.Genesis.Alloc.Validators.Map()[stakerID]; ok {
				addr = v.Address
				createdEpoch = 0
				createdTime = net.Genesis.Time
			} else {
				// event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
				creationLog, err := gdb.App().EvmLogs().Find([][]common.Hash{{sfcpos.Topics.CreatedStake}, {futils.U64to256(uint64(stakerID))}})
				if err != nil {
					return genStore, err
				}
				for _, l := range creationLog {
					if l.Address != sfc.ContractAddress {
						continue
					}
					addr = common.BytesToAddress(l.Topics[2][12:])
					block := gdb.GetBlock(idx.Block(l.BlockNumber))
					createdTime = block.Time
				}
				var e = idx.Epoch(0)
				for ; e < epoch; e++ {
					stats := gdb.GetEpochStats(e)
					if createdTime <= stats.End {
						createdEpoch = e
						break
					}
				}
				if e == epoch {
					createdEpoch = epoch
				}
			}
			// recover deactivation time/epoch from logs
			{
				// event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
				withdrawalLog, err := gdb.App().EvmLogs().Find([][]common.Hash{{sfcpos.Topics.WithdrawnStake}, {futils.U64to256(uint64(stakerID))}})
				if err != nil {
					return genStore, err
				}
				for _, l := range withdrawalLog {
					if l.Address != sfc.ContractAddress {
						continue
					}
					block := gdb.GetBlock(idx.Block(l.BlockNumber))
					oDeactivatedTime = block.Time
				}
				var e = idx.Epoch(0)
				for ; e < epoch; e++ {
					stats := gdb.GetEpochStats(e)
					if oDeactivatedTime <= stats.End {
						oDeactivatedEpoch = e
						break
					}
				}
				if e == epoch {
					oDeactivatedEpoch = epoch
				}
			}
		}

		if oDeactivatedEpoch != 0 {
			// withdrawn
			oStatus = 1
		}
		if status == sfctype.OfflineBit {
			oStatus = 1 << 3
		}
		if status == sfctype.ForkBit {
			oStatus = drivertype.DoublesignBit
		}
		if oStatus != 0 && oDeactivatedEpoch == 0 {
			// find highest epoch where validator was active
			for e := epoch; e > createdEpoch; e-- {
				if !gdb.App().HasEpochValidator(e, stakerID) {
					oDeactivatedEpoch = e - 1
					oDeactivatedTime = gdb.GetEpochStats(e - 1).End
				}
			}
		}

		metadata.Validators[stakerID-1] = gpos.Validator{
			ID:               oidx.ValidatorID(stakerID),
			Address:          addr,
			PubKey:           validatorpk.PubKey{},
			CreationTime:     ointer.Timestamp(createdTime),
			CreationEpoch:    oidx.Epoch(createdEpoch),
			DeactivatedTime:  ointer.Timestamp(oDeactivatedTime),
			DeactivatedEpoch: oidx.Epoch(oDeactivatedEpoch),
			Status:           oStatus,
		}
	}

	log.Info("Recovering validators public keys")
	isEpochInterested := func(e idx.Epoch, metadata *genesisstore.Metadata) bool {
		for _, v := range metadata.Validators {
			if v.PubKey.Empty() && oidx.Epoch(e) >= v.CreationEpoch && (v.DeactivatedEpoch == 0 || oidx.Epoch(e) <= v.DeactivatedEpoch) {
				return true
			}
		}
		return false
	}
	// find latest secp256k1 public key
	for e := epoch; e > 0; e-- {
		if !isEpochInterested(e, &metadata) {
			continue
		}
		gdb.ForEachEpochEvent(e, func(event *inter.Event) bool {
			if event.Lamport > 1000 {
				return false
			}
			if !metadata.Validators[event.Creator-1].PubKey.Empty() {
				return true
			}
			if cdb.GetEventConfirmedOn(event.Hash()) == 0 {
				// cannot use non-confirmed events for the sake of determinism
				return true
			}
			pk := event.RecoverPubkey()
			if pk != nil {
				metadata.Validators[event.Creator-1].PubKey.Type = validatorpk.Types.Secp256k1
				metadata.Validators[event.Creator-1].PubKey.Raw = crypto.FromECDSAPub(pk)
				return isEpochInterested(e, &metadata)
			}
			return true
		})
	}

	log.Info("Exported validators", "elapsed", common.PrettyDuration(time.Since(start)))

	// export delegations
	log.Info("Exporting delegations")
	validatorsMap := metadata.Validators.Map()
	gdb.App().ForEachSfcDelegation(func(it sfctype.SfcDelegationAndID) {
		// retrieve stake amount
		delegation, _ := sfcCaller.Delegations(callsOpts, it.ID.Delegator, futils.U64toBig(uint64(it.ID.StakerID)))
		if delegation.Amount == nil {
			return
		}
		if delegation.Amount.Sign() == 0 {
			return
		}

		// retrieve rewards
		var pendingRewards *big.Int
		if delegation.DeactivatedTime.Sign() == 0 {
			pendingRewards, _, _, _ = sfcCaller.CalcDelegationRewards(callsOpts, it.ID.Delegator, futils.U64toBig(uint64(it.ID.StakerID)), futils.U64toBig(0), futils.U64toBig(uint64(epoch)))
		}
		if pendingRewards == nil {
			pendingRewards = new(big.Int)
		}

		v := validatorsMap[oidx.ValidatorID(it.ID.StakerID)]
		if delegation.DeactivatedTime.Sign() != 0 && v.Status != drivertype.DoublesignBit {
			// finalize withdrawal balance if validator isn't a cheater
			addBalance(genStore, statedb, it.ID.Delegator, delegation.Amount)
		} else {
			// retrieve lockup info
			info, _ := sfcCaller.LockedDelegations(callsOpts, it.ID.Delegator, futils.U64toBig(uint64(it.ID.StakerID)))
			if info.Duration == nil {
				info.Duration = new(big.Int)
			}
			if info.EndTime == nil {
				info.EndTime = new(big.Int)
			}
			if info.FromEpoch == nil {
				info.FromEpoch = new(big.Int)
			}
			earlyUnlockPenalty, _ := sfcCaller.DelegationEarlyWithdrawalPenalty(callsOpts, it.ID.Delegator, futils.U64toBig(uint64(it.ID.StakerID)))
			if earlyUnlockPenalty == nil {
				earlyUnlockPenalty = new(big.Int)
			}
			lockedStake := delegation.Amount
			if info.Duration.Sign() == 0 {
				lockedStake = new(big.Int)
			}
			// add delegation
			addDelegation(genStore, it.ID.Delegator, oidx.ValidatorID(it.ID.StakerID), genesis.Delegation{
				Stake:              delegation.Amount,
				Rewards:            pendingRewards,
				LockedStake:        lockedStake,
				LockupFromEpoch:    oidx.Epoch(info.FromEpoch.Uint64()),
				LockupEndTime:      ointer.Timestamp(info.EndTime.Uint64()),
				LockupDuration:     ointer.Timestamp(info.Duration.Uint64()),
				EarlyUnlockPenalty: earlyUnlockPenalty,
			})
		}
	})
	log.Info("Exported delegations", "elapsed", common.PrettyDuration(time.Since(start)))

	withdrawalRequestLogs, err := gdb.App().EvmLogs().Find([][]common.Hash{{sfcpos.Topics.CreatedWithdrawRequest}})
	if err != nil {
		return genStore, err
	}
	paidWrs := map[withdrawalRequestID]bool{}
	for _, l := range withdrawalRequestLogs {
		if l.Address != sfc.ContractAddress {
			continue
		}
		receiver := common.BytesToAddress(l.Topics[2].Bytes()[12:])
		wrID := new(big.Int).SetBytes(l.Data[0:32])
		fullWrID := withdrawalRequestID{receiver, wrID.String()}
		if paidWrs[fullWrID] {
			continue
		}
		res, _ := sfcCaller.WithdrawalRequests(callsOpts, receiver, wrID)
		if res.Amount == nil || res.Amount.Sign() == 0 {
			continue
		}
		stakerID := oidx.ValidatorID(res.StakerID.Uint64())
		v := validatorsMap[stakerID]
		if v.Status != drivertype.DoublesignBit {
			// finalize withdrawal balance if validator isn't a cheater
			addBalance(genStore, statedb, receiver, res.Amount)
		} else {
			addDelegation(genStore, receiver, stakerID, genesis.Delegation{
				Stake:              res.Amount,
				Rewards:            new(big.Int),
				LockedStake:        new(big.Int),
				LockupFromEpoch:    0,
				LockupEndTime:      0,
				LockupDuration:     0,
				EarlyUnlockPenalty: new(big.Int),
			})
		}
		paidWrs[fullWrID] = true
	}

	metadata.ExtraData = []byte{}
	metadata.FirstEpoch = oidx.Epoch(epoch + 1)
	metadata.Time = ointer.Timestamp(lastBlock.Time + inter.Timestamp(time.Second))
	metadata.PrevEpochTime = ointer.Timestamp(gdb.GetEpochStats(epoch - 1).End)
	metadata.DriverOwner, _ = sfcCaller.Owner(callsOpts)
	metadata.TotalSupply = gdb.App().GetTotalSupply()
	genStore.SetMetadata(metadata)

	return genStore, nil
}

func ExportEncodedState(encodedPath, tmpPath string, gdb *gossip.Store, cdb *poset.Store, net *lachesis.Config) (*genesisstore.Store, error) {
	_ = os.RemoveAll(tmpPath)

	log.Info("Exporting go-opera genesis state", "path", tmpPath)
	genStore, err := ExportState(tmpPath, gdb, cdb, net)
	if err != nil {
		return genStore, err
	}

	log.Info("Encoding go-opera genesis state", "path", encodedPath)
	fh, err := os.Create(encodedPath)
	if err != nil {
		return genStore, err
	}
	defer fh.Close()
	err = genesisstore.WriteGenesisStore(fh, genStore)
	if err != nil {
		return genStore, err
	}
	return genStore, nil
}

func exportState(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	cfg := makeAllConfigs(ctx)

	gdb, cdb := makeStores(cfg.Node.DataDir, &cfg.Lachesis)
	defer gdb.Close()
	defer cdb.Close()

	genFilePath := ctx.Args().First()

	start := time.Now()

	tmpPath := path.Join(cfg.Node.DataDir, "tmp", "export")
	_ = os.RemoveAll(tmpPath)

	genStore, err := ExportEncodedState(genFilePath, tmpPath, gdb, cdb, &cfg.Lachesis.Net)
	if genStore != nil {
		defer genStore.Close()
		defer os.RemoveAll(tmpPath)
	}
	if err != nil {
		utils.Fatalf("Failed to export state: %v\n", err)
	}

	log.Info("Exported go-opera genesis state", "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}
