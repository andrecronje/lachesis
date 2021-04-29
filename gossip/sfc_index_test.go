package gossip

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-lachesis/gossip/sfc110"
	"github.com/Fantom-foundation/go-lachesis/gossip/sfc202"
	"github.com/Fantom-foundation/go-lachesis/gossip/sfcproxy"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/lachesis/genesis/sfc"
	"github.com/Fantom-foundation/go-lachesis/lachesis/genesis/sfc/sfcpos"
	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/Fantom-foundation/go-lachesis/utils"
)

type commonSfc interface {
	CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error)
	CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error)
}

func TestSFC(t *testing.T) {
	logger.SetTestMode(t)

	env := newTestEnv()
	defer env.Close()

	sfcProxy, err := sfcproxy.NewContract(sfc.ContractAddress, env)
	require.NoError(t, err)

	var (
		sfc11 *sfc110.Contract
		sfc22 *sfc202.Contract

		prev struct {
			epoch  *big.Int
			reward *big.Int
		}
	)

	_ = true &&

		t.Run("Genesis v1.0.0", func(t *testing.T) {
			// nothing to do
		}) &&

		t.Run("Some transfers I", func(t *testing.T) {
			cicleTransfers(t, env, 3)
		}) &&

		t.Run("Upgrade to v1.1.0-rc1", func(t *testing.T) {
			require := require.New(t)

			r := env.ApplyBlock(nextEpoch,
				env.Contract(1, utils.ToFtm(0), sfc110.ContractBin),
			)
			newImpl := r[0].ContractAddress

			admin := env.Payer(1)
			tx, err := sfcProxy.ContractTransactor.UpgradeTo(admin, newImpl)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)

			impl, err := sfcProxy.Implementation(env.ReadOnly())
			require.NoError(err)
			require.Equal(newImpl, impl, "SFC-proxy: implementation address")

			sfc11, err = sfc110.NewContract(sfc.ContractAddress, env)
			require.NoError(err)

			epoch, err := sfc11.ContractCaller.CurrentEpoch(env.ReadOnly())
			require.NoError(err)
			require.Equal(0, epoch.Cmp(big.NewInt(2)), "current epoch")
		}) &&

		t.Run("Upgrade stakers storage", func(t *testing.T) {
			require := require.New(t)

			stakers, err := sfc11.StakersLastID(env.ReadOnly())
			require.NoError(err)
			txs := make([]*eth.Transaction, 0, int(stakers.Int64()))
			for i := stakers.Int64(); i > 0; i-- {
				tx, err := sfc11.UpgradeStakerStorage(env.Payer(int(i)), big.NewInt(i))
				require.NoError(err)
				txs = append(txs, tx)
			}
			env.ApplyBlock(sameEpoch, txs...)

		}) &&

		t.Run("Some transfers II", func(t *testing.T) {
			cicleTransfers(t, env, 3)
		}) &&

		t.Run("Create staker 4", func(t *testing.T) {
			require := require.New(t)

			newStake := utils.ToFtm(genesisStake / 2)
			minStake, err := sfc11.MinStake(env.ReadOnly())
			require.NoError(err)
			require.Greater(newStake.Cmp(minStake), 0,
				fmt.Sprintf("newStake(%s) < minStake(%s)", newStake, minStake))

			env.ApplyBlock(sameEpoch,
				env.Transfer(1, 4, big.NewInt(0).Add(newStake, utils.ToFtm(10))),
			)
			tx, err := sfc11.CreateStake(env.Payer(4, newStake), nil)
			require.NoError(err)
			env.ApplyBlock(nextEpoch, tx)
			newId, err := sfc11.SfcAddressToStakerID(env.ReadOnly(), env.Address(4))
			require.NoError(err)
			env.AddValidator(idx.StakerID(newId.Uint64()))
		}) &&

		t.Run("Create delegator 5", func(t *testing.T) {
			require := require.New(t)

			newDelegation := utils.ToFtm(genesisStake / 2)
			env.ApplyBlock(sameEpoch,
				env.Transfer(1, 5, big.NewInt(0).Add(newDelegation, utils.ToFtm(10))),
			)

			staker, err := sfc11.SfcAddressToStakerID(env.ReadOnly(), env.Address(4))
			require.NoError(err)

			tx, err := sfc11.CreateDelegation(env.Payer(5, newDelegation), staker)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)
			env.AddDelegator(env.Address(5))
		}) &&

		t.Run("Check if locking is not set", func(t *testing.T) {
			require := require.New(t)

			require.Zero(
				utils.H256toU64(env.State().GetState(
					sfc.ContractAddress,
					sfcpos.FirstLockedUpEpoch())),
			)
		}) &&

		t.Run("Check rewards before sfc2", func(t *testing.T) {
			env.ApplyBlock(nextEpoch) // clear epoch
			env.ApplyBlock(nextEpoch)
			const validatorCommission int64 = 15
			rewards := requireRewards(t, env, sfc11, []int64{200, 200, 200, 100 + validatorCommission, (100 - validatorCommission)})
			prev.reward = rewards[0]
		}) &&

		t.Run("Upgrade to v2.0.2-rc.2", func(t *testing.T) {
			require := require.New(t)

			r := env.ApplyBlock(nextEpoch,
				env.Contract(1, utils.ToFtm(0), sfc202.ContractBin),
			)
			newImpl := r[0].ContractAddress

			admin := env.Payer(1)
			tx, err := sfcProxy.ContractTransactor.UpgradeTo(admin, newImpl)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)

			impl, err := sfcProxy.Implementation(env.ReadOnly())
			require.NoError(err)
			require.Equal(newImpl, impl, "SFC-proxy: implementation address")

			sfc22, err = sfc202.NewContract(sfc.ContractAddress, env)
			require.NoError(err)

			epoch, err := sfc22.ContractCaller.CurrentEpoch(env.ReadOnly())
			require.NoError(err)
			require.Equal(0, epoch.Cmp(big.NewInt(6)), "current epoch: %d", epoch.Uint64())
			prev.epoch = epoch
		}) &&

		t.Run("Upgrade delegation storage", func(t *testing.T) {
			require := require.New(t)

			var txs []*eth.Transaction
			for _, delegator := range env.delegators {
				tx, err := sfc22.UpgradeDelegationStorage(env.Payer(5), delegator)
				require.NoError(err)
				txs = append(txs, tx)
			}
			env.ApplyBlock(sameEpoch, txs...)
		}) &&

		t.Run("Check if locking is false", func(t *testing.T) {
			require := require.New(t)

			require.Zero(
				utils.H256toU64(env.State().GetState(
					sfc.ContractAddress,
					sfcpos.FirstLockedUpEpoch())),
			)
		}) &&

		t.Run("Some transfers III", func(t *testing.T) {
			cicleTransfers(t, env, 3)
		}) &&

		t.Run("Enable lockedup functionality", func(t *testing.T) {
			require := require.New(t)

			tx, err := sfc22.StartLockedUp(env.Payer(1), prev.epoch)
			require.NoError(err)
			env.ApplyBlock(nextEpoch, tx)

			epoch, err := sfc22.FirstLockedUpEpoch(env.ReadOnly())
			require.NoError(err)
			require.Equal(0, epoch.Cmp(prev.epoch), "1st locked-up epoch")

			raw := new(big.Int).SetBytes(env.State().GetState(
				sfc.ContractAddress,
				sfcpos.FirstLockedUpEpoch()).Bytes())
			require.Equal(0, epoch.Cmp(raw), "raw 1st locked-up epoch")

			raw = new(big.Int).SetBytes(env.State().GetState(
				sfc.ContractAddress,
				sfcpos.CurrentSealedEpoch()).Bytes())
			require.Equal(0, epoch.Cmp(raw), "raw last sealed epoch")
		}) &&

		t.Run("Check if locking is true", func(t *testing.T) {
			require := require.New(t)

			epoch, err := sfc22.CurrentSealedEpoch(env.ReadOnly())
			require.NoError(err)

			require.GreaterOrEqual(
				epoch.Uint64(),
				utils.H256toU64(env.State().GetState(
					sfc.ContractAddress,
					sfcpos.FirstLockedUpEpoch())),
			)
		}) &&

		t.Run("Check rewards after sfc2", func(t *testing.T) {
			require := require.New(t)

			env.ApplyBlock(nextEpoch) // clear epoch
			env.ApplyBlock(nextEpoch)

			rewards := requireRewards(t, env, sfc22, []int64{200, 200, 200, 100 + 15, 85})
			expected := new(big.Int).Div(prev.reward, big.NewInt(10))
			expected = new(big.Int).Mul(expected, big.NewInt(3))
			require.Equal(0, rewards[0].Cmp(expected), "%s != 0.3*%s", rewards[0], prev.reward)
			prev.reward = expected
		}) &&

		t.Run("Lockup stake 4 on half period", func(t *testing.T) {
			require := require.New(t)

			lockDuration, err := sfc22.MaxLockupDuration(env.ReadOnly())
			require.NoError(err)

			lockDuration = big.NewInt(0).Div(lockDuration, big.NewInt(2))
			tx, err := sfc22.LockUpStake(env.Payer(4), lockDuration)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)
		}) &&

		t.Run("Check rewards after stake lock", func(t *testing.T) {
			require := require.New(t)

			env.ApplyBlock(nextEpoch) // clear epoch
			env.ApplyBlock(nextEpoch)
			const (
				unlockedRewardRatio int64 = 30
				lockedRewardRatio   int64 = 70
			)
			rewards := requireRewards(t, env, sfc22, []int64{
				200 * unlockedRewardRatio,
				200 * unlockedRewardRatio,
				200 * unlockedRewardRatio,
				115 * (unlockedRewardRatio + lockedRewardRatio/2),
				85 * unlockedRewardRatio})
			require.Equal(0, rewards[0].Cmp(prev.reward), "%s != %s", rewards[0], prev.reward)
		}) &&

		t.Run("Lockup delegation 5-to-4 on quarter period", func(t *testing.T) {
			require := require.New(t)

			staker, err := sfc22.SfcAddressToStakerID(env.ReadOnly(), env.Address(4))
			require.NoError(err)

			lockDuration, err := sfc22.MaxLockupDuration(env.ReadOnly())
			require.NoError(err)

			lockDuration = big.NewInt(0).Div(lockDuration, big.NewInt(4))
			tx, err := sfc22.LockUpDelegation(env.Payer(5), lockDuration, staker)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)
		}) &&

		t.Run("Delegation 5-to-4 is locked up", func(t *testing.T) {
			require := require.New(t)

			env.ApplyBlock(nextEpoch) // clear epoch

			staker, err := sfc22.SfcAddressToStakerID(env.ReadOnly(), env.Address(4))
			require.NoError(err)
			isLockedUp, err := sfc22.IsDelegationLockedUp(env.ReadOnly(), env.Address(5), staker)
			require.NoError(err)

			require.True(isLockedUp)
		}) &&

		t.Run("Check rewards after delegation lock", func(t *testing.T) {
			require := require.New(t)

			env.ApplyBlock(nextEpoch) // clear epoch
			env.ApplyBlock(nextEpoch)

			const (
				unlockedRewardRatio int64 = 30 * 2
				lockedRewardRatio   int64 = 70 * 2
			)
			rewards := requireRewards(t, env, sfc22, []int64{
				200 * unlockedRewardRatio,
				200 * unlockedRewardRatio,
				200 * unlockedRewardRatio,
				115 * (unlockedRewardRatio + lockedRewardRatio/2),
				85 * (unlockedRewardRatio + lockedRewardRatio/4)})
			require.Equal(0, rewards[0].Cmp(prev.reward), "%s != %s", rewards[0], prev.reward)
		}) &&

		t.Run("MinGasPrice up", func(t *testing.T) {
			require := require.New(t)

			admin := env.Payer(1)

			gp := env.App.MinGasPrice()
			gpExp := new(big.Int).Mul(gp, big.NewInt(10))
			tx, err := sfc22.UpdateMinGasPrice(admin, gpExp)
			require.NoError(err)
			env.ApplyBlock(sameEpoch, tx)

			env.ApplyBlock(nextEpoch)
			gpGot := env.App.MinGasPrice()
			require.Equal(gpExp, gpGot, "%s != %s", gpExp, gpGot)
		})

}

func requireRewards(
	t *testing.T, env *testEnv, sfc commonSfc, stakes []int64,
) (
	rewards []*big.Int,
) {
	require := require.New(t)

	epoch, err := sfc.CurrentSealedEpoch(env.ReadOnly())
	require.NoError(err)

	validators := env.Validators()
	rewards = make([]*big.Int, len(validators)+len(env.delegators))
	for i, id := range validators {
		staker := big.NewInt(int64(id))
		rewards[i], _, _, err = sfc.CalcValidatorRewards(env.ReadOnly(), staker, epoch, big.NewInt(1))
		require.NoError(err)
		t.Logf("validator reward %d: %s", i+1, rewards[i])
	}

	for i, addr := range env.delegators {
		i += len(validators)
		switch sfc := sfc.(type) {
		case *sfc110.Contract:
			rewards[i], _, _, err = sfc.CalcDelegationRewards(env.ReadOnly(), addr, epoch, big.NewInt(1))
			require.NoError(err)
		case *sfc202.Contract:
			sum := new(big.Int)
			for _, id := range validators {
				staker := big.NewInt(int64(id))
				r, _, _, err := sfc.CalcDelegationRewards(env.ReadOnly(), addr, staker, epoch, big.NewInt(1))
				require.NoError(err)
				sum = new(big.Int).Add(sum, r)
			}
			rewards[i] = sum
		default:
			panic("unknown contract type")
		}
		t.Logf("delegator reward %d: %s", i+1, rewards[i])
	}

	for i := range validators {
		if i == 0 {
			continue
		}

		a := new(big.Int).Mul(rewards[0], big.NewInt(stakes[i]))
		b := new(big.Int).Mul(rewards[i], big.NewInt(stakes[0]))
		want := new(big.Int).Div(b, rewards[0])
		require.Equal(
			0, a.Cmp(b),
			"reward#0: %s, reward#%d: %s. Got %d:%d, want %d:%s proportion (validator)",
			rewards[0], i, rewards[i],
			stakes[0], stakes[i],
			stakes[0], want,
		)
	}

	for i := range env.delegators {
		i += len(validators)
		a := new(big.Int).Mul(rewards[0], big.NewInt(stakes[i]))
		b := new(big.Int).Mul(rewards[i], big.NewInt(stakes[0]))
		want := new(big.Int).Div(b, rewards[0])
		require.Equal(
			0, a.Cmp(b),
			"reward#0: %s, reward#%d: %s. Got %d:%d, want %d:%s proportion (delegator)",
			rewards[0], i, rewards[i],
			stakes[0], stakes[i],
			stakes[0], want,
		)
	}

	return
}

func printEpochStats(t *testing.T, env *testEnv, sfc2 *sfc202.Contract) {
	epoch, err := sfc2.CurrentSealedEpoch(env.ReadOnly())
	require.NoError(t, err)
	es, err := sfc2.EpochSnapshots(env.ReadOnly(), epoch)
	require.NoError(t, err)
	t.Logf("Epoch%sStat{dir: %s, BaseRewardPerSecond: %s, TotalBaseRewardWeight: %s}",
		epoch, es.Duration, es.BaseRewardPerSecond, es.TotalBaseRewardWeight)
}

func printValidators(t *testing.T, env *testEnv, sfc2 *sfc202.Contract) {
	require := require.New(t)

	max, err := sfc2.StakersLastID(env.ReadOnly())
	require.NoError(err)

	for id := big.NewInt(1); id.Cmp(max) <= 0; id.Add(id, big.NewInt(1)) {
		s, err := sfc2.Stakers(env.ReadOnly(), id)
		require.NoError(err)
		t.Logf("%s: %#v", id, s)
	}
}

func cicleTransfers(t *testing.T, env *testEnv, count uint64) {
	require := require.New(t)

	balances := make([]*big.Int, 3)
	for i := range balances {
		balances[i] = env.State().GetBalance(env.Address(i + 1))
	}

	for i := uint64(0); i < count; i++ {
		env.ApplyBlock(sameEpoch,
			env.Transfer(1, 2, utils.ToFtm(100)),
		)
		env.ApplyBlock(sameEpoch,
			env.Transfer(2, 3, utils.ToFtm(100)),
		)
		env.ApplyBlock(sameEpoch,
			env.Transfer(3, 1, utils.ToFtm(100)),
		)
	}

	gp := env.App.MinGasPrice()
	gas := big.NewInt(0).Mul(big.NewInt(int64(count*gasLimit)), gp)
	for i := range balances {
		require.Equal(
			big.NewInt(0).Sub(balances[i], gas),
			env.State().GetBalance(env.Address(i+1)),
			fmt.Sprintf("account%d", i),
		)
	}
}
