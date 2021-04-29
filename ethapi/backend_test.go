package ethapi

//go:generate go run github.com/golang/mock/mockgen -destination=backend_mock_test.go -package=ethapi -source=backend.go Backend
//go:generate go run github.com/golang/mock/mockgen -destination=account_mock_test.go -package=ethapi -mock_names Backend=AmBackend github.com/ethereum/go-ethereum/accounts Backend,Wallet

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	notify "github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/inter/sfctype"
	"github.com/Fantom-foundation/go-lachesis/kvdb/nokeyiserr"
	"github.com/Fantom-foundation/go-lachesis/kvdb/table"
)

const (
	globalGasCap uint64 = 100000
)

func TestGetAPIs(t *testing.T) {
	b := newTestBackend(t)

	require.NotPanics(t, func() {
		res := GetAPIs(b)
		require.NotEmpty(t, res)
	})
}

func mockAmBackend(t *testing.T) *AmBackend {
	ctrl := gomock.NewController(t)
	b := NewAmBackend(ctrl)

	b.EXPECT().Wallets().
		Return([]accounts.Wallet{
			testWallet(ctrl),
			testWallet(ctrl),
			testWallet(ctrl),
		}).
		AnyTimes()

	b.EXPECT().Subscribe(gomock.Any()).
		Return(notify.NewSubscription(func(c <-chan struct{}) error { return nil })).
		AnyTimes()

	return b
}

func newTestBackend(t *testing.T) *MockBackend {
	ctrl := gomock.NewController(t)
	b := NewMockBackend(ctrl)
	initTestBackend(t, b)

	return b
}

func initTestBackend(t *testing.T, b *MockBackend) {
	b.EXPECT().GetTd(gomock.Any()).
		Return(big.NewInt(1)).
		AnyTimes()

	b.EXPECT().SuggestPrice(gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().RPCGasCap().
		Return(uint64(1)).
		AnyTimes()

	b.EXPECT().RPCTxFeeCap().
		Return(float64(1)).
		AnyTimes()

	b.EXPECT().Stats().
		Return(2, 2).
		AnyTimes()

	b.EXPECT().ProtocolVersion().
		Return(1).
		AnyTimes()

	b.EXPECT().BlockByHash(gomock.Any(), gomock.Any()).
		Return(&evmcore.EvmBlock{
			EvmHeader: evmcore.EvmHeader{
				Number:     big.NewInt(1),
				Hash:       common.Hash{2},
				ParentHash: common.Hash{3},
				Root:       common.Hash{4},
				TxHash:     common.Hash{5},
				Time:       6,
				Coinbase:   common.Address{7},
				GasLimit:   8,
				GasUsed:    9,
			},
			Transactions: types.Transactions{
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			},
		}, nil).
		AnyTimes()

	b.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).
		Return(&evmcore.EvmBlock{
			EvmHeader: evmcore.EvmHeader{
				Number:     big.NewInt(1),
				Hash:       common.Hash{2},
				ParentHash: common.Hash{3},
				Root:       common.Hash{4},
				TxHash:     common.Hash{5},
				Time:       6,
				Coinbase:   common.Address{7},
				GasLimit:   8,
				GasUsed:    9,
			},
			Transactions: types.Transactions{
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			},
		}, nil).
		AnyTimes()

	b.EXPECT().ForEachEpochEvent(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	b.EXPECT().TxPoolContent().
		Return(
			map[common.Address]types.Transactions{
				common.Address{1}: types.Transactions{
					types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
					types.NewTransaction(2, common.Address{2}, big.NewInt(2), 2, big.NewInt(0), []byte{}),
				},
			},
			map[common.Address]types.Transactions{
				common.Address{1}: types.Transactions{
					types.NewTransaction(3, common.Address{3}, big.NewInt(3), 3, big.NewInt(0), []byte{}),
					types.NewTransaction(4, common.Address{4}, big.NewInt(4), 4, big.NewInt(0), []byte{}),
				},
			},
		).
		AnyTimes()

	b.EXPECT().GetPoolTransactions().
		Return(types.Transactions{
			types.NewTransaction(3, common.Address{3}, big.NewInt(3), 3, big.NewInt(0), []byte{}),
			types.NewTransaction(4, common.Address{4}, big.NewInt(4), 4, big.NewInt(0), []byte{}),
		}, nil).
		Times(1)

	b.EXPECT().ChainConfig().
		Return(&params.ChainConfig{
			ChainID: big.NewInt(1),
		}).
		AnyTimes()

	b.EXPECT().Progress().
		Return(PeerProgress{
			CurrentEpoch:     1,
			CurrentBlock:     2,
			CurrentBlockHash: hash.Event{3},
			CurrentBlockTime: inter.Timestamp(time.Now().Add(-91 * time.Minute).UnixNano()),
			HighestBlock:     5,
			HighestEpoch:     6,
		}).
		AnyTimes()

	// Set state DB

	db1 := rawdb.NewDatabase(
		nokeyiserr.Wrap(
			table.New(
				memorydb.New(), []byte("evm1_"))))
	stateDB, _ := state.New(common.HexToHash("0x0"), state.NewDatabase(db1), nil)
	stateDB.SetNonce(common.Address{1}, 1)
	stateDB.AddBalance(common.Address{1}, big.NewInt(10))
	stateDB.SetCode(common.Address{1}, []byte{1, 2, 3})
	b.EXPECT().StateAndHeaderByNumberOrHash(gomock.Any(), gomock.Any()).
		Return(stateDB, &evmcore.EvmHeader{}, nil).
		Times(1)

	evm := vm.NewEVM(vm.Context{}, stateDB, &params.ChainConfig{}, vm.Config{})
	b.EXPECT().GetEVM(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(evm, func() error { return nil }, nil).
		Times(1)

	keyStore := keystore.NewKeyStore("/tmp", 2, 2)
	for _, ac := range keyStore.Accounts() {
		keyStore.Delete(ac, "1234")
	}
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, mockAmBackend(t), keyStore)
	b.EXPECT().AccountManager().
		Return(am).
		AnyTimes()

	b.EXPECT().GetTransaction(gomock.Any(), gomock.Any()).
		Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			uint64(1), uint64(1), nil,
		).
		Times(1)

	b.EXPECT().CurrentEpoch(gomock.Any()).
		Return(idx.Epoch(1)).
		AnyTimes()

	b.EXPECT().GetConsensusTime(gomock.Any(), gomock.Any()).
		Return(inter.Timestamp(1), nil).
		AnyTimes()

	b.EXPECT().GetEpochStats(gomock.Any(), gomock.Any()).
		Return(&sfctype.EpochStats{
			Start: 1,
			End:   2,
			Epoch: 1,
		}, nil).
		AnyTimes()

	b.EXPECT().GetEvent(gomock.Any(), gomock.Any()).
		Return(&inter.Event{
			EventHeader: inter.EventHeader{
				EventHeaderData: inter.EventHeaderData{
					Version:       1,
					Epoch:         2,
					Seq:           1,
					Frame:         1,
					IsRoot:        true,
					Creator:       1,
					PrevEpochHash: common.Hash{0},
					Parents:       nil,
					GasPowerLeft:  inter.GasPowerLeft{},
					GasPowerUsed:  0,
					Lamport:       0,
					ClaimedTime:   0,
					MedianTime:    0,
					TxHash:        common.Hash{},
					Extra:         nil,
				},
				Sig: nil,
			},
			Transactions: types.Transactions{
				types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			},
		}, nil).
		AnyTimes()

	b.EXPECT().GetEventHeader(gomock.Any(), gomock.Any()).
		Return(&inter.EventHeaderData{
			Version:       1,
			Epoch:         2,
			Seq:           1,
			Frame:         1,
			IsRoot:        true,
			Creator:       1,
			PrevEpochHash: common.Hash{0},
			Parents:       nil,
			GasPowerLeft:  inter.GasPowerLeft{},
			GasPowerUsed:  0,
			Lamport:       0,
			ClaimedTime:   0,
			MedianTime:    0,
			TxHash:        common.Hash{},
			Extra:         nil,
		}, nil).
		AnyTimes()

	b.EXPECT().GetHeads(gomock.Any(), gomock.Any()).
		Return(hash.Events{
			hash.Event{1},
		}, nil).
		AnyTimes()

	b.EXPECT().GetDelegation(gomock.Any(), gomock.Any()).
		Return(&sfctype.SfcDelegation{
			CreatedEpoch:     1,
			CreatedTime:      2,
			DeactivatedEpoch: 0,
			DeactivatedTime:  0,
			Amount:           nil,
		}, nil).
		AnyTimes()

	b.EXPECT().GetDelegationClaimedRewards(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetDelegationsOf(gomock.Any(), gomock.Any()).
		Return([]sfctype.SfcDelegationAndID{{
			Delegation: &sfctype.SfcDelegation{
				CreatedEpoch:     1,
				CreatedTime:      2,
				DeactivatedEpoch: 0,
				DeactivatedTime:  0,
				Amount:           nil,
			},
			ID: sfctype.DelegationID{
				Delegator: common.Address{1},
				StakerID:  1,
			},
		}}, nil).
		AnyTimes()

	b.EXPECT().GetDowntime(gomock.Any(), gomock.Any()).
		Return(idx.Block(1), inter.Timestamp(1), nil).
		AnyTimes()

	b.EXPECT().GetOriginationScore(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetRewardWeights(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetStaker(gomock.Any(), gomock.Any()).
		Return(&sfctype.SfcStaker{
			CreatedEpoch:     1,
			CreatedTime:      1,
			DeactivatedEpoch: 0,
			DeactivatedTime:  0,
			Address:          common.Address{1},
			StakeAmount:      big.NewInt(1),
			DelegatedMe:      big.NewInt(0),
		}, nil).
		AnyTimes()

	b.EXPECT().GetStakerPoI(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetValidationScore(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetStakerClaimedRewards(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetStakerDelegationsClaimedRewards(gomock.Any(), gomock.Any()).
		Return(big.NewInt(1), nil).
		AnyTimes()

	b.EXPECT().GetStakers(gomock.Any()).
		Return([]sfctype.SfcStakerAndID{
			{
				StakerID: 1,
				Staker: &sfctype.SfcStaker{
					CreatedEpoch:     1,
					CreatedTime:      1,
					DeactivatedEpoch: 0,
					DeactivatedTime:  0,
					Address:          common.Address{1},
					StakeAmount:      big.NewInt(1),
					DelegatedMe:      big.NewInt(0),
				},
			},
		}, nil).
		AnyTimes()

	b.EXPECT().GetStakerID(gomock.Any(), gomock.Any()).
		Return(idx.StakerID(1), nil).
		AnyTimes()

	b.EXPECT().ValidatorTimeDrifts(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(map[idx.StakerID]map[hash.Event]time.Duration{
			idx.StakerID(1): {
				hash.HexToEventHash("0x1"): time.Second,
				hash.HexToEventHash("0x2"): 2 * time.Second,
			},
			idx.StakerID(2): {
				hash.HexToEventHash("0x3"): 3 * time.Second,
				hash.HexToEventHash("0x4"): 4 * time.Second,
			},
		}, nil).
		AnyTimes()

	db2 := rawdb.NewDatabase(
		nokeyiserr.Wrap(
			table.New(
				memorydb.New(), []byte("evm2_"))))
	b.EXPECT().ChainDb().
		Return(db2).
		AnyTimes()

	b.EXPECT().ExtRPCEnabled().
		Return(false).
		AnyTimes()

	b.EXPECT().GetPoolTransaction(gomock.Any()).
		Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
		).AnyTimes()
}

func testWallet(ctrl *gomock.Controller) accounts.Wallet {
	w := NewMockWallet(ctrl)

	w.EXPECT().URL().
		Return(accounts.URL{
			Scheme: "https",
			Path:   "test.ru/test",
		}).
		AnyTimes()

	w.EXPECT().Accounts().
		Return([]accounts.Account{
			accounts.Account{
				Address: common.Address{1},
				URL:     w.URL(),
			},
		}).
		AnyTimes()

	w.EXPECT().Status().
		Return("ok", nil).
		AnyTimes()

	w.EXPECT().Contains(gomock.Any()).
		Return(true).
		AnyTimes()

	w.EXPECT().SignTextWithPassphrase(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			make([]byte, 128, 128),
			nil,
		).
		AnyTimes()

	w.EXPECT().SignTxWithPassphrase(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			nil,
		).
		AnyTimes()

	w.EXPECT().SignTx(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			types.NewTransaction(1, common.Address{1}, big.NewInt(1), 1, big.NewInt(0), []byte{}),
			nil,
		).
		AnyTimes()

	w.EXPECT().SignText(gomock.Any(), gomock.Any()).
		Return(
			make([]byte, 128, 128),
			nil,
		).
		AnyTimes()

	return w
}
