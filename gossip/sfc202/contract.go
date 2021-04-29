// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfc202

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedDelegationReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedValidatorReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dagSfcAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedWithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"LockingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"LockingStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minVersion\",\"type\":\"uint256\"}],\"name\":\"NetworkUpgradeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"PartialWithdrawnByRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UnstashedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldStakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"UpdatedGasPowerAllocationRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minGasPrice\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinGasPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UpdatedOfflinePenaltyThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"}],\"name\":\"UpdatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"UpdatedStakerMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSfcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"UpdatedStakerSfcAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnStake\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minVersion\",\"type\":\"uint256\"}],\"name\":\"_activateNetworkUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"name\":\"_sfcAddressToStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"_syncDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_syncStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_updateBaseRewardPerSec\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"_updateGasPowerAllocationRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minGasPrice\",\"type\":\"uint256\"}],\"name\":\"_updateMinGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"_updateOfflinePenaltyThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"_upgradeDelegationStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_upgradeStakerStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcDelegationCompoundRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcValidatorCompoundRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimValidatorCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStakeWithAddresses\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegationEarlyWithdrawalPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"e\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"epochValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txRewardWeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstLockedUpEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"getDelegationRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"getValidatorRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"isDelegationLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"staker\",\"type\":\"uint256\"}],\"name\":\"isStakeLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"legacyDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"lockUpDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"lockUpStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxStakerMetadataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"partialWithdrawByRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegationPartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawStakePartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedDelegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedStakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNum\",\"type\":\"uint256\"}],\"name\":\"startLockedUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurntLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockedRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"updateStakerMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"updateStakerSfcAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040819052600080546001600160a01b031916339081178255918291907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350615b7e806100536000396000f3fe6080604052600436106105945760003560e01c8063846ebb77116102e0578063c3d74f1a11610184578063dd099bb6116100e1578063f2fde38b11610095578063f8b18d8a1161006f578063f8b18d8a146113de578063f99837e614611408578063fd5e6dd11461143857610594565b8063f2fde38b14611357578063f3ae5b1a1461138a578063f5a83c7d146113b457610594565b8063df4f49d4116100c6578063df4f49d414611303578063eac3baf21461132d578063ec6a7f1c1461090d57610594565b8063dd099bb6146112a0578063df0e307a146112d957610594565b8063cc8c212011610138578063cfd5fa0c1161011d578063cfd5fa0c146111f2578063d845fc901461122b578063dc599bb11461127057610594565b8063cc8c212014611122578063cda5826a146111c857610594565b8063c4b5dd7e11610169578063c4b5dd7e14610599578063c9400d4f146110f8578063cb1c4e671461059957610594565b8063c3d74f1a146110b0578063c41b6405146110e357610594565b8063a4b89fab1161023d578063b1e64339116101f1578063bb03a4bd116101cb578063bb03a4bd14611048578063bed9d8611461107e578063c312eb071461109357610594565b8063b1e6433914610f95578063b42cb58d14610fbf578063b9029d5014610ff257610594565b8063a778651511610222578063a778651514610f41578063aa34eb4514610f56578063ab2273c014610f8057610594565b8063a4b89fab14610efc578063a70da4d214610f2c57610594565b806390475ae4116102945780639864183d116102795780639864183d14610e0357806398ec2de514610e48578063a289ad6e14610ee757610594565b806390475ae414610d0e57806396060e7114610dcd57610594565b80638da5cb5b116102c55780638da5cb5b14610c9e5780638e431b8d14610ccf5780638f32d59b14610cf957610594565b8063846ebb7714610c56578063876f7e2a14610c8957610594565b80633fee10a81161044757806366799a54116103a457806375b9d3d8116103585780637cacb1d6116103325780637cacb1d614610bee5780637f664d8714610c0357806381d9dc7a14610c4157610594565b806375b9d3d814610b765780637667180814610baf5780637b015db914610bc457610594565b80636f498663116103895780636f49866314610ad4578063715018a614610b0d5780637424036214610b2257610594565b806366799a5414610a865780636e1a767a14610abf57610594565b80635573184d116103fb5780635e2308d2116103e05780635e2308d21461078257806360c7e37f1461059957806363321e2714610a5357610594565b80635573184d146109e75780635b81b88614610a2057610594565b80634e5a23281161042c5780634e5a23281461093757806354d77ed21461064657806354fd4d501461099d57610594565b80633fee10a81461090d5780634bd202dc1461092257610594565b80632265f284116104f55780632e5f75ef116104a957806333a149121161048e57806333a1491214610830578063375b3c0a146108e35780633d0317fe146108f857610594565b80632e5f75ef146107eb57806330fa99291461081b57610594565b80632709275e116104da5780632709275e1461078257806328dca8ff14610797578063295cccba146107c157610594565b80632265f2841461073d57806326682c711461075257610594565b806319ddb54f1161054c5780631d58179c116105315780631d58179c146106465780631e8a69561461065b578063223fae09146106cc57610594565b806319ddb54f146105995780631c3c60c81461061457610594565b80630a29180c1161057d5780630a29180c146105d55780630d4955e3146105ea5780630d7b2609146105ff57610594565b8063029859921461059957806308728f6e146105c0575b600080fd5b3480156105a557600080fd5b506105ae6114be565b60408051918252519081900360200190f35b3480156105cc57600080fd5b506105ae6114cb565b3480156105e157600080fd5b506105ae6114d1565b3480156105f657600080fd5b506105ae6114d7565b34801561060b57600080fd5b506105ae6114df565b34801561062057600080fd5b506106446004803603604081101561063757600080fd5b50803590602001356114e6565b005b34801561065257600080fd5b506105ae61157e565b34801561066757600080fd5b506106856004803603602081101561067e57600080fd5b5035611583565b60408051998a5260208a0198909852888801969096526060880194909452608087019290925260a086015260c085015260e084015261010083015251908190036101200190f35b3480156106d857600080fd5b50610705600480360360408110156106ef57600080fd5b506001600160a01b0381351690602001356115d1565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b34801561074957600080fd5b506105ae611620565b34801561075e57600080fd5b506106446004803603604081101561077557600080fd5b5080359060200135611627565b34801561078e57600080fd5b506105ae611924565b3480156107a357600080fd5b50610644600480360360208110156107ba57600080fd5b5035611934565b3480156107cd57600080fd5b50610644600480360360208110156107e457600080fd5b50356119ec565b3480156107f757600080fd5b506106446004803603604081101561080e57600080fd5b50803590602001356119fa565b34801561082757600080fd5b506105ae611a92565b34801561083c57600080fd5b506106446004803603602081101561085357600080fd5b81019060208101813564010000000081111561086e57600080fd5b82018360208201111561088057600080fd5b803590602001918460018302840111640100000000831117156108a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a98945050505050565b3480156108ef57600080fd5b506105ae611b5a565b34801561090457600080fd5b506105ae611b69565b34801561091957600080fd5b506105ae611b6f565b34801561092e57600080fd5b506105ae611b76565b34801561094357600080fd5b506109706004803603604081101561095a57600080fd5b506001600160a01b038135169060200135611b7c565b60408051958652602086019490945284840192909252606084015215156080830152519081900360a00190f35b3480156109a957600080fd5b506109b2611bb9565b604080517fffffff00000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b3480156109f357600080fd5b506105ae60048036036040811015610a0a57600080fd5b506001600160a01b038135169060200135611bdd565b348015610a2c57600080fd5b5061070560048036036020811015610a4357600080fd5b50356001600160a01b0316611c47565b348015610a5f57600080fd5b506105ae60048036036020811015610a7657600080fd5b50356001600160a01b0316611c84565b348015610a9257600080fd5b506105ae60048036036040811015610aa957600080fd5b506001600160a01b038135169060200135611ca3565b348015610acb57600080fd5b506105ae611cc0565b348015610ae057600080fd5b506105ae60048036036040811015610af757600080fd5b506001600160a01b038135169060200135611cc6565b348015610b1957600080fd5b50610644611ce3565b348015610b2e57600080fd5b50610b5860048036036060811015610b4557600080fd5b5080359060208101359060400135611d93565b60408051938452602084019290925282820152519081900360600190f35b348015610b8257600080fd5b5061064460048036036040811015610b9957600080fd5b506001600160a01b038135169060200135611dd3565b348015610bbb57600080fd5b506105ae611e3f565b348015610bd057600080fd5b5061064460048036036020811015610be757600080fd5b5035611e48565b348015610bfa57600080fd5b506105ae611ed7565b348015610c0f57600080fd5b50610c2d60048036036020811015610c2657600080fd5b5035611edd565b604080519115158252519081900360200190f35b348015610c4d57600080fd5b506105ae611f11565b348015610c6257600080fd5b5061064460048036036020811015610c7957600080fd5b50356001600160a01b0316611f17565b348015610c9557600080fd5b50610644611f77565b348015610caa57600080fd5b50610cb361208b565b604080516001600160a01b039092168252519081900360200190f35b348015610cdb57600080fd5b506105ae60048036036020811015610cf257600080fd5b503561209a565b348015610d0557600080fd5b50610c2d6120e6565b61064460048036036060811015610d2457600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135640100000000811115610d5857600080fd5b820183602082011115610d6a57600080fd5b80359060200191846001830284011164010000000083111715610d8c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506120f7945050505050565b348015610dd957600080fd5b50610b5860048036036060811015610df057600080fd5b5080359060208101359060400135612179565b348015610e0f57600080fd5b50610b5860048036036080811015610e2657600080fd5b506001600160a01b038135169060208101359060408101359060600135612196565b348015610e5457600080fd5b50610e7260048036036020811015610e6b57600080fd5b50356121da565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610eac578181015183820152602001610e94565b50505050905090810190601f168015610ed95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610ef357600080fd5b506105ae612275565b348015610f0857600080fd5b5061064460048036036040811015610f1f57600080fd5b508035906020013561227b565b348015610f3857600080fd5b506105ae6124f8565b348015610f4d57600080fd5b506105ae6124fe565b348015610f6257600080fd5b5061064460048036036020811015610f7957600080fd5b503561250b565b348015610f8c57600080fd5b506105ae61259a565b348015610fa157600080fd5b5061064460048036036020811015610fb857600080fd5b50356125a0565b348015610fcb57600080fd5b506105ae60048036036020811015610fe257600080fd5b50356001600160a01b0316612714565b348015610ffe57600080fd5b506110226004803603604081101561101557600080fd5b5080359060200135612769565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561105457600080fd5b506106446004803603606081101561106b57600080fd5b508035906020810135906040013561279c565b34801561108a57600080fd5b50610644612a95565b610644600480360360208110156110a957600080fd5b5035612d80565b3480156110bc57600080fd5b50610644600480360360208110156110d357600080fd5b50356001600160a01b0316612d8a565b3480156110ef57600080fd5b50610644612f8d565b34801561110457600080fd5b506106446004803603602081101561111b57600080fd5b503561305b565b6106446004803603602081101561113857600080fd5b81019060208101813564010000000081111561115357600080fd5b82018360208201111561116557600080fd5b8035906020019184600183028401116401000000008311171561118757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613172945050505050565b3480156111d457600080fd5b50610644600480360360208110156111eb57600080fd5b503561317e565b3480156111fe57600080fd5b50610c2d6004803603604081101561121557600080fd5b506001600160a01b038135169060200135613189565b34801561123757600080fd5b50610b586004803603608081101561124e57600080fd5b506001600160a01b0381351690602081013590604081013590606001356131eb565b34801561127c57600080fd5b506106446004803603604081101561129357600080fd5b5080359060200135613209565b3480156112ac57600080fd5b50610b58600480360360408110156112c357600080fd5b506001600160a01b038135169060200135613219565b3480156112e557600080fd5b50610644600480360360208110156112fc57600080fd5b5035613245565b34801561130f57600080fd5b50610b586004803603602081101561132657600080fd5b503561355e565b34801561133957600080fd5b506106446004803603602081101561135057600080fd5b503561357f565b34801561136357600080fd5b506106446004803603602081101561137a57600080fd5b50356001600160a01b03166135de565b34801561139657600080fd5b50610644600480360360208110156113ad57600080fd5b5035613640565b3480156113c057600080fd5b50610644600480360360208110156113d757600080fd5b5035613826565b3480156113ea57600080fd5b506106446004803603602081101561140157600080fd5b50356138b5565b34801561141457600080fd5b506106446004803603604081101561142b57600080fd5b5080359060200135613d0c565b34801561144457600080fd5b506114626004803603602081101561145b57600080fd5b5035613d18565b604080519a8b5260208b0199909952898901979097526060890195909552608088019390935260a087019190915260c086015260e08501526001600160a01b039081166101008501521661012083015251908190036101400190f35b670de0b6b3a76400005b90565b60235481565b60285481565b6301e1338090565b6212750090565b6114ee6120e6565b61153f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080518381526020810183905281517f95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c929181900390910190a15050565b600390565b601f6020528060005260406000206000915090508060010154908060020154908060030154908060040154908060050154908060060154908060070154908060080154908060090154905089565b602e602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154908060030154908060040154908060050154908060060154905087565b62e4e1c090565b33600061163382612714565b905061163e81613d75565b61164781613de1565b61165081611edd565b156116a2576040805162461bcd60e51b815260206004820152600f60248201527f7374616b65206973206c6f636b65640000000000000000000000000000000000604482015290519081900360640190fd5b6116aa6114be565b8310156116fe576040805162461bcd60e51b815260206004820152601060248201527f746f6f20736d616c6c20616d6f756e7400000000000000000000000000000000604482015290519081900360640190fd5b600081815260208052604090206005015480611718611b5a565b8501111561176d576040805162461bcd60e51b815260206004820152601c60248201527f6d757374206c65617665206174206c65617374206d696e5374616b6500000000604482015290519081900360640190fd5b60008281526020805260409020600701548482039061178b82613e47565b10156117de576040805162461bcd60e51b815260206004820152601460248201527f746f6f206d7563682064656c65676174696f6e73000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152602d6020908152604080832089845290915290206003015415611857576040805162461bcd60e51b815260206004820152601360248201527f7772494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b600083815260208080526040808320600501805489900390556001600160a01b0387168352602d8252808320898452909152902083815560030185905561189c611e3f565b6001600160a01b0385166000818152602d602090815260408083208b8452825280832060018101959095554260029095019490945583518a8152908101919091528083018890529151859282917fde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d9181900360600190a461191c8361357f565b505050505050565b600060646301c9c3805b04905090565b60008181526020805260409020600901546001600160a01b0316156119a0576040805162461bcd60e51b815260206004820152600f60248201527f616c726561647920757064617465640000000000000000000000000000000000604482015290519081900360640190fd5b6119a981613e74565b6000908152602080526040902060088101546009909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055565b6119f7816000613ed6565b50565b611a026120e6565b611a53576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080518381526020810183905281517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34929181900390910190a15050565b60265481565b6000611aa333612714565b9050611aae81613e74565b611ab661259a565b82511115611b0b576040805162461bcd60e51b815260206004820152601060248201527f746f6f20626967206d6574616461746100000000000000000000000000000000604482015290519081900360640190fd5b6000818152602b602090815260409091208351611b2a9285019061595d565b5060405181907fb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd8390600090a25050565b6a02a055184a310c1260000090565b60245481565b62093a8090565b60255481565b602d602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401549293919290919060ff1685565b7f323032000000000000000000000000000000000000000000000000000000000090565b600080611bea8484613189565b90506000611c33620f4240611c05611c00611e3f565b613ffc565b6001600160a01b03881660009081526031602090815260408083208a84529091529020600201548590614013565b60600151620f424003925050505b92915050565b6029602052600090815260409020805460018201546002830154600384015460048501546005860154600690960154949593949293919290919087565b6001600160a01b0381166000908152602160205260409020545b919050565b603260209081526000928352604080842090915290825290205481565b602f5481565b602c60209081526000928352604080842090915290825290205481565b611ceb6120e6565b611d3c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000806000611da06159db565b600080611db08989896001614121565b6040830151602084015193519093019092019b909a509098509650505050505050565b611ddd828261427b565b6001600160a01b0382166000818152602e6020908152604080832085845282529182902060040154825190815291518493849390927f19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff92918290030190a45050565b601e5460010190565b611e506120e6565b611ea1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b601e5481565b60008181526030602052604081206001015415801590611c4157505060009081526030602052604090206001015442111590565b60225481565b6001600160a01b038116600090815260296020526040902060040154611f6e5760405162461bcd60e51b8152600401808060200182810382526021815260200180615ae76021913960400191505060405180910390fd5b6119f7816142f3565b336000818152602c60209081526040808320838052909152902054819080611fe6576040805162461bcd60e51b815260206004820152600a60248201527f6e6f207265776172647300000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038084166000908152602c60209081526040808320838052909152808220829055519184169183156108fc0291849190818181858888f1935050505015801561203a573d6000803e3d6000fd5b50816001600160a01b0316836001600160a01b03167f80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126836040518082815260200191505060405180910390a3505050565b6000546001600160a01b031690565b6000806120a683611edd565b905060006120d5620f42406120bc611c00611e3f565b6000878152603060205260409020600201548590614013565b60600151620f424003949350505050565b6000546001600160a01b0316331490565b6001600160a01b0383161580159061211757506001600160a01b03821615155b612168576040805162461bcd60e51b815260206004820152600f60248201527f696e76616c696420616464726573730000000000000000000000000000000000604482015290519081900360640190fd5b612174838334846143bb565b505050565b60008060006121866159db565b600080611db08989896000614121565b60008060006121a36159db565b6000806121b48a8a8a8a6001614623565b6040830151602084015193519093019092019750955093505050505b9450945094915050565b602b6020908152600091825260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909183018282801561226d5780601f106122425761010080835404028352916020019161226d565b820191906000526020600020905b81548152906001019060200180831161225057829003601f168201915b505050505081565b60335481565b612286611c00611e3f565b6122d7576040805162461bcd60e51b815260206004820152601960248201527f6665617475726520776173206e6f742061637469766174656400000000000000604482015290519081900360640190fd5b336122e1816142f3565b6122eb818361427b565b6122f4826147f2565b6122fc6114df565b8310158015612312575061230e6114d7565b8311155b612363576040805162461bcd60e51b815260206004820152601260248201527f696e636f7272656374206475726174696f6e0000000000000000000000000000604482015290519081900360640190fd5b6000612375428563ffffffff61485b16565b6000848152603060205260409020600101549091508111156123c85760405162461bcd60e51b8152600401808060200182810382526022815260200180615ac56022913960400191505060405180910390fd5b6123d28284613189565b15612424576040805162461bcd60e51b815260206004820152601160248201527f616c7265616479206c6f636b6564207570000000000000000000000000000000604482015290519081900360640190fd5b61242e82846148b5565b6001600160a01b03821660009081526032602090815260408083208684529091528082209190915580516060810190915280612468611e3f565b8152602080820184905260409182018790526001600160a01b038516600081815260318352838120888252835283902084518155918401516001830155929091015160029091015583907f823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d6124db611e3f565b60408051918252602082018690528051918290030190a350505050565b60275481565b6000606462e4e1c061192e565b6125136120e6565b612564576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6040805182815290517f35feeeac858525cae277d98c1c4792d0550aeab30f107addc09d8d5279faa53f9181900360200190a150565b61010090565b336125aa816142f3565b6001600160a01b0381166000908152602e6020908152604080832085845290915290206125d7828461494e565b6125e182846149d1565b6125e9611e3f565b60028201554260038201556004810154600084815260208052604090206005015415612642576000848152602080526040902060070154612630908263ffffffff614a4d16565b60008581526020805260409020600701555b600061264e8486613189565b1561269a5761265f84868485614a8f565b905081811061266f575060001981015b6001600160a01b03841660009081526032602090815260408083208884529091529020805482900390555b600483018054829003905560268054829003905560405185906001600160a01b038616907f912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f50646990600090a36126ee8486611dd3565b60008581526020805260409020600501541561270d5761270d8561357f565b5050505050565b6001600160a01b0381166000908152602160205260408120548061273c576000915050611c9e565b60008181526020805260409020600901546001600160a01b03848116911614611c41576000915050611c9e565b6000918252601f602090815260408084209284529190529020805460018201546002830154600390930154919390929190565b336127a6816142f3565b6001600160a01b0381166000908152602e6020908152604080832086845290915290206127d3828561494e565b6127dd82856149d1565b6127e56114be565b831015612839576040805162461bcd60e51b815260206004820152601060248201527f746f6f20736d616c6c20616d6f756e7400000000000000000000000000000000604482015290519081900360640190fd5b6004810154806128476114be565b850111156128865760405162461bcd60e51b8152600401808060200182810382526021815260200180615b086021913960400191505060405180910390fd5b6001600160a01b0383166000908152602d60209081526040808320898452909152902060030154156128ff576040805162461bcd60e51b815260206004820152601360248201527f7772494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b600061290b8487613189565b156129575761291c84878785614a8f565b905084811061292c575060001984015b6001600160a01b03841660009081526032602090815260408083208984529091529020805482900390555b60048301805486900390556000868152602080526040902060050154156129ab576000868152602080526040902060070154612999908663ffffffff614a4d16565b60008781526020805260409020600701555b6001600160a01b0384166000908152602d602090815260408083208a845290915290208681558186036003909101556129e2611e3f565b6001600160a01b0385166000818152602d602090815260408083208c8452825291829020600180820195909555426002820155600401805460ff19168517905581518b81529081019390935282810188905251889282917fde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d9181900360600190a4612a6d8487611dd3565b600086815260208052604090206005015415612a8c57612a8c8661357f565b50505050505050565b336000612aa182612714565b6000818152602080526040902060040154909150612b06576040805162461bcd60e51b815260206004820152601960248201527f7374616b6572207761736e277420646561637469766174656400000000000000604482015290519081900360640190fd5b612b0e611b6f565b600082815260208052604090206004015401421015612b74576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b612b7c61157e565b600082815260208052604090206003015401612b96611e3f565b1015612be9576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b600081815260208080526040808320600881018054600583018054845488865560018087018a9055600287018a9055600387018a9055600487018a905592899055600686018990556007860189905573ffffffffffffffffffffffffffffffffffffffff1980851690955560099095018054909416909355602b9095529285206001600160a01b039093169490939092908216151590612c899084615a03565b6001600160a01b038088166000908152602160205260408082208290559187168152908120558115612cc657600086815260208052604090208290555b60238054600019019055602454612ce3908563ffffffff614a4d16565b60245580612d27576040516001600160a01b0388169085156108fc029086906000818181858888f19350505050158015612d21573d6000803e3d6000fd5b50612d2b565b8392505b602854612d3e908463ffffffff61485b16565b60285560408051848152905187917f8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6919081900360200190a250505050505050565b6119f73382614ad2565b336001600160a01b038216811415612de9576040805162461bcd60e51b815260206004820152601060248201527f7468652073616d65206164647265737300000000000000000000000000000000604482015290519081900360640190fd5b6000612df482612714565b9050612dff81613e74565b6001600160a01b0383166000908152602160205260409020541580612e3b57506001600160a01b03831660009081526021602052604090205481145b612e8c576040805162461bcd60e51b815260206004820152601460248201527f6164647265737320616c72656164792075736564000000000000000000000000604482015290519081900360640190fd5b60008181526020808052604080832060098101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03898116918217909255878216808752602186528487208790559086528386208790556008909201541684528184208590558352602c825280832083805290915290205415612f47576001600160a01b038281166000908152602c60208181526040808420848052808352818520958916855292825280842084805282528320845490555290555b826001600160a01b0316826001600160a01b0316827f7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b20694660405160405180910390a4505050565b336000612f9982612714565b9050612fa481613d75565b612fad81613de1565b612fb681611edd565b15613008576040805162461bcd60e51b815260206004820152600f60248201527f7374616b65206973206c6f636b65640000000000000000000000000000000000604482015290519081900360640190fd5b613010611e3f565b6000828152602080526040808220600381019390935542600490930192909255905182917ff7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc6091a25050565b6130636120e6565b6130b4576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b601e54811161310a576040805162461bcd60e51b815260206004820152601760248201527f63616e277420737461727420696e207468652070617374000000000000000000604482015290519081900360640190fd5b602f54158061311c5750601e54602f54115b61316d576040805162461bcd60e51b815260206004820152601360248201527f6665617475726520776173207374617274656400000000000000000000000000604482015290519081900360640190fd5b602f55565b6119f7333334846143bb565b6119f7816001613ed6565b6001600160a01b0382166000908152603160209081526040808320848452909152812060010154158015906131e457506001600160a01b03831660009081526031602090815260408083208584529091529020600101544211155b9392505050565b60008060006131f86159db565b6000806121b48a8a8a8a6000614623565b61321582826001614dca565b5050565b603160209081526000928352604080842090915290825290208054600182015460029092015490919083565b3361324f816142f3565b613257615a47565b506001600160a01b0381166000908152602e60209081526040808320858452825291829020825160e0810184528154815260018201549281019290925260028101549282019290925260038201546060820181905260048301546080830152600583015460a083015260069092015460c08201529061331d576040805162461bcd60e51b815260206004820152601d60248201527f64656c65676174696f6e207761736e2774206465616374697661746564000000604482015290519081900360640190fd5b6000838152602080526040902060050154156133fe5761333b611b6f565b816060015101421015613395576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b61339d61157e565b8160400151016133ab611e3f565b10156133fe576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6000838152602080805260408083205460808501516001600160a01b038716808652602e85528386208987528552838620868155600181810188905560028083018990556003830189905560048301899055600583018990556006909201889055828852603187528588208b89528752858820888155808201899055909101879055908652603285528386208987529094529184208490556025805460001901905560265492161515916134b8908263ffffffff614a4d16565b602655816134fc576040516001600160a01b0386169082156108fc029083906000818181858888f193505050501580156134f6573d6000803e3d6000fd5b50613500565b8092505b602754613513908463ffffffff61485b16565b60275560408051848152905187916001600160a01b038816917f87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f9181900360200190a3505050505050565b60306020526000908152604090208054600182015460029092015490919083565b61358881613e74565b600081815260208080526040918290206005810154600790910154835191825291810191909152815183927f509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c928290030190a250565b6135e66120e6565b613637576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6119f781614f39565b61364b611c00611e3f565b61369c576040805162461bcd60e51b815260206004820152601960248201527f6665617475726520776173206e6f742061637469766174656400000000000000604482015290519081900360640190fd5b60006136a733612714565b90506136b2816147f2565b6136ba6114df565b82101580156136d057506136cc6114d7565b8211155b613721576040805162461bcd60e51b815260206004820152601260248201527f696e636f7272656374206475726174696f6e0000000000000000000000000000604482015290519081900360640190fd5b6000613733428463ffffffff61485b16565b905061373e82611edd565b15613790576040805162461bcd60e51b815260206004820152601160248201527f616c7265616479206c6f636b6564207570000000000000000000000000000000604482015290519081900360640190fd5b61379982614fe6565b60405180606001604052806137ac611e3f565b81526020808201849052604091820186905260008581526030825282902083518155908301516001820155910151600290910155817f71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b61380a611e3f565b60408051918252602082018590528051918290030190a2505050565b61382e6120e6565b61387f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6040805182815290517fa3deceaa35ccc5aa4f1e61ffe83a006792b8989d4e80dd2c8aa07ba8a923cde19181900360200190a150565b336000818152602d602090815260408083208584529091529020600201548190613926576040805162461bcd60e51b815260206004820152601560248201527f7265717565737420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152602d6020908152604080832086845290915290206004810154905460ff909116908180156139725750600081815260208052604090206005015415155b15613a8b5761397f611b6f565b6001600160a01b0385166000908152602d60209081526040808320898452909152902060020154014210156139fb576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b613a0361157e565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206001015401613a33611e3f565b1015613a86576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b613b9f565b81613b9f57613a98611b6f565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206002015401421015613b14576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b613b1c61157e565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206001015401613b4c611e3f565b1015613b9f576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b600081815260208080526040808320546001600160a01b0388168452602d83528184208985529092528220600381018054848355600180840186905560028401869055918590556004909201805460ff191690559091161515908415613c1a57602654613c12908263ffffffff614a4d16565b602655613c31565b602454613c2d908263ffffffff614a4d16565b6024555b81613c72576040516001600160a01b0387169082156108fc029083906000818181858888f19350505050158015613c6c573d6000803e3d6000fd5b50613c76565b8092505b8415613c9757602754613c8f908463ffffffff61485b16565b602755613cae565b602854613caa908463ffffffff61485b16565b6028555b604080518981528615156020820152808201859052905185916001600160a01b03808a1692908b16917fd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd919081900360600190a45050505050505050565b61321582826000614dca565b60208052600090815260409020805460018201546002830154600384015460048501546005860154600687015460078801546008890154600990990154979896979596949593949293919290916001600160a01b0391821691168a565b613d7e81613e74565b6000818152602080526040902060040154156119f7576040805162461bcd60e51b815260206004820152601560248201527f7374616b65722069732064656163746976617465640000000000000000000000604482015290519081900360640190fd5b601e546000828152602080526040902060060154146119f7576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c207265776172647320636c61696d6564000000000000000000604482015290519081900360640190fd5b6000611c41620f4240613e68613e5b611620565b859063ffffffff61506116565b9063ffffffff6150ba16565b60008181526020805260409020600501546119f7576040805162461bcd60e51b815260206004820152601460248201527f7374616b657220646f65736e2774206578697374000000000000000000000000604482015290519081900360640190fd5b336000613ee282612714565b9050613eed81613e74565b613ef56159db565b600080613f058460008989614121565b925092509250600083604001518460200151856000015101019050613f41602060008781526020019081526020016000206006015484846150fc565b6000858152602080526040902060060182905560608401516033805490910190558615613f7757613f7285826151fc565b613faf565b6040516001600160a01b0387169082156108fc029083906000818181858888f19350505050158015613fad573d6000803e3d6000fd5b505b6040805182815260208101859052808201849052905186917f2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c919081900360600190a25050505050505050565b600080602f54118015611c41575050602f54111590565b61401b6159db565b60405180608001604052806000815260200160008152602001600081526020016000815250905083156140f1576000614052611924565b620f424003905060006140766140666114d7565b613e68848763ffffffff61506116565b905084156140c757600083526140a1620f4240613e68614094611924565b8a9063ffffffff61506116565b60208401526140bd620f4240613e68898463ffffffff61506116565b60408401526140ea565b6140d9620f4240613e68614094611924565b835260006020840181905260408401525b5050614103565b84815260006020820181905260408201525b6040810151602082015182518703030360608201525b949350505050565b6141296159db565b6000858152602080526040812060060154819061414a908790600101615288565b600088815260208052604090206006015490965086116141925750506040805160808101825260008082526020820181905291810182905260608101829052915084906121d0565b61419a6159db565b60405180608001604052806000815260200160008152602001600081526020016000815250905060008790505b601e5481111580156141da575086880181105b1561425357600086156141f7575060408201516020830151835101015b6141ff6159db565b6142128b8461420c6124fe565b8561529d565b8051855101855260208082015190860180519091019052604080820151908601805190910190526060908101519085018051909101905250506001016141c7565b60008882116142645750600061426b565b5060001981015b9199979850909695505050505050565b6001600160a01b0382166000908152602e60209081526040808320848452909152902060040154613215576040805162461bcd60e51b815260206004820152601860248201527f64656c65676174696f6e20646f65736e27742065786973740000000000000000604482015290519081900360640190fd5b6001600160a01b038116600090815260296020526040902060040154156119f7576001600160a01b03166000818152602960208181526040808420602e8352818520600680830180548852918552928620825481556001808401805491830191909155600280850180549184019190915560038086018054918501919091556004808701805491860191909155600580880180549187019190915586549590980194909455998952969095529186905592859055928490559383905590829055918190559055565b6001600160a01b0384166000908152602160205260409020541580156143f757506001600160a01b038316600090815260216020526040902054155b614448576040805162461bcd60e51b815260206004820152601560248201527f7374616b657220616c7265616479206578697374730000000000000000000000604482015290519081900360640190fd5b614450611b5a565b8210156144a4576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b60228054600101908190556001600160a01b03808616600090815260216020908152604080832085905592871682528282208490558382528052206005018390556144ed611e3f565b600082815260208052604090206001808201929092554260028201556008810180546001600160a01b03808a1673ffffffffffffffffffffffffffffffffffffffff199283161790925560098301805492891692909116919091179055601e5460069091015560238054909101905560245461456f908463ffffffff61485b16565b6024556040805184815290516001600160a01b0387169183917f0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d399181900360200190a38151156145c2576145c282611a98565b836001600160a01b0316856001600160a01b03161461270d57836001600160a01b0316856001600160a01b0316827f7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b20694660405160405180910390a45050505050565b61462b6159db565b600080614636615a47565b506001600160a01b0388166000908152602e602090815260408083208a8452825291829020825160e08101845281548152600180830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a0840181905260069091015460c08401526146b791899101615288565b965080606001516000146146c757fe5b868160a0015110614701575050604080516080810182526000808252602082018190529181018290526060810182905292508591506147e7565b6147096159db565b60405180608001604052806000815260200160008152602001600081526020016000815250905060008890505b601e548111158015614749575087890181105b156147c35760008715614766575060408201516020830151835101015b61476e6159db565b6147828d8d8561477c6124fe565b866153cd565b805185510185526020808201519086018051909101905260408082015190860180519091019052606090810151908501805190910190525050600101614736565b60008982116147d4575060006147db565b5060001981015b91955088945090925050505b955095509592505050565b6147fb81613d75565b6000818152602080526040902054156119f7576040805162461bcd60e51b815260206004820152601760248201527f7374616b65722073686f756c6420626520616374697665000000000000000000604482015290519081900360640190fd5b6000828201838110156131e4576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0382166000818152602e60209081526040808320858452825280832060050154938352603182528083208584529091529020600101546148fb82615526565b1015612174576040805162461bcd60e51b815260206004820152601e60248201527f6e6f7420616c6c206c6f636b7570207265776172647320636c61696d65640000604482015290519081900360640190fd5b614958828261427b565b6001600160a01b0382166000908152602e6020908152604080832084845290915290206003015415613215576040805162461bcd60e51b815260206004820152601960248201527f64656c65676174696f6e20697320646561637469766174656400000000000000604482015290519081900360640190fd5b601e546001600160a01b0383166000908152602e6020908152604080832085845290915290206005015414613215576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c207265776172647320636c61696d6564000000000000000000604482015290519081900360640190fd5b60006131e483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061553b565b6001600160a01b0384166000908152603260209081526040808320868452909152812054614ac9908390613e68908663ffffffff61506116565b95945050505050565b614adb826142f3565b614ae4816147f2565b614aec6114be565b341015614b40576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152602e6020908152604080832084845290915290206004015415614bb9576040805162461bcd60e51b815260206004820152601960248201527f64656c65676174696f6e20616c72656164792065786973747300000000000000604482015290519081900360640190fd5b6001600160a01b03821660009081526021602052604090205415614c24576040805162461bcd60e51b815260206004820152600f60248201527f616c7265616479207374616b696e670000000000000000000000000000000000604482015290519081900360640190fd5b6000818152602080526040902060070154614c45903463ffffffff61485b16565b6000828152602080526040902060050154614c5f90613e47565b1015614cb2576040805162461bcd60e51b815260206004820152601a60248201527f7374616b65722773206c696d6974206973206578636565646564000000000000604482015290519081900360640190fd5b614cba615a47565b614cc2611e3f565b8152426020808301918252346080840181815260c08501868152601e5460a087019081526001600160a01b0389166000908152602e865260408082208a8352875280822089518155975160018901558089015160028901556060890151600389015593516004880155905160058701559051600690950194909455918052912060070154614d559163ffffffff61485b16565b6000838152602080526040902060070155602580546001019055602654614d82903463ffffffff61485b16565b60265560408051348152905183916001600160a01b038616917ffd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be9181900360200190a3505050565b33614dd4816142f3565b6001600160a01b0381166000908152602e602090815260408083208684529091529020614e01828561494e565b614e096159db565b600080614e1a858860008b8a614623565b925092509250600083604001518460200151856000015101019050614e44856005015484846150fc565b600585018290556040848101516020808701516001600160a01b038a166000908152603283528481208d825290925292902080546002909304909101909101905560608401516033805490910190558615614ea957614ea48689836155d2565b614ee1565b6040516001600160a01b0387169082156108fc029083906000818181858888f19350505050158015614edf573d6000803e3d6000fd5b505b6040805182815260208101859052808201849052905189916001600160a01b038916917f2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce9181900360600190a3505050505050505050565b6001600160a01b038116614f7e5760405162461bcd60e51b8152600401808060200182810382526026815260200180615a9f6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6000818152602080805260408083206006015460309092529091206001015461500e82615526565b1015613215576040805162461bcd60e51b815260206004820152601e60248201527f6e6f7420616c6c206c6f636b7570207265776172647320636c61696d65640000604482015290519081900360640190fd5b60008261507057506000611c41565b8282028284828161507d57fe5b04146131e45760405162461bcd60e51b8152600401808060200182810382526021815260200180615b296021913960400191505060405180910390fd5b60006131e483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250615764565b818310615150576040805162461bcd60e51b815260206004820152601560248201527f65706f636820697320616c726561647920706169640000000000000000000000604482015290519081900360640190fd5b601e548211156151a7576040805162461bcd60e51b815260206004820152600c60248201527f6675747572652065706f63680000000000000000000000000000000000000000604482015290519081900360640190fd5b81811015612174576040805162461bcd60e51b815260206004820152601160248201527f6e6f2065706f63687320636c61696d6564000000000000000000000000000000604482015290519081900360640190fd5b600082815260208052604081206005015461521d908363ffffffff61485b16565b60008481526020805260409020600501819055602454909150615246908363ffffffff61485b16565b6024556040805182815260208101849052815185927fa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9928290030190a2505050565b600082615296575080611c41565b5090919050565b6152a56159db565b6000806152b287876157c9565b6000878152601f602090815260408083208b8452909152812054919250906152e0908663ffffffff61485b16565b6000888152601f602090815260408083208c8452909152812060010154919250615310838363ffffffff61485b16565b905080615348576040518060800160405280600081526020016000815260200160008152602001600081525095505050505050614119565b6000615371615364620f4240613e68868d63ffffffff61506116565b859063ffffffff61485b16565b905061538782613e68878463ffffffff61506116565b95505050505050600061539a87876158b9565b90506153c2826153a988613ffc565b60008a8152603060205260409020600201548490614013565b979650505050505050565b6153d56159db565b6000806153e287876157c9565b6000878152601f602090815260408083208b84529091528120805460019091015492935091615411908761485b565b90506000615425838363ffffffff61485b16565b90508061545d576040518060800160405280600081526020016000815260200160008152602001600081525095505050505050614ac9565b6001600160a01b038b166000908152602e602090815260408083208d8452909152812060040154615494908963ffffffff61485b16565b905060006154b2620f4240613e68613e5b828e63ffffffff614a4d16565b90506154c883613e68888463ffffffff61506116565b965050505050505060006154dd8888886158f6565b905061551a826154ec88613ffc565b6001600160a01b038b1660009081526031602090815260408083208d84529091529020600201548490614013565b98975050505050505050565b6000908152601f602052604090206001015490565b600081848411156155ca5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561558f578181015183820152602001615577565b50505050905090810190601f1680156155bc5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60008281526020805260409020600701546155f3908263ffffffff61485b16565b600083815260208052604090206005015461560d90613e47565b1015615660576040805162461bcd60e51b815260206004820152601a60248201527f7374616b65722773206c696d6974206973206578636565646564000000000000604482015290519081900360640190fd5b6001600160a01b0383166000908152602e60209081526040808320858452909152812060040154615697908363ffffffff61485b16565b6001600160a01b0385166000908152602e6020908152604080832087845282528083206004018490559080529020600701549091506156dc908363ffffffff61485b16565b6000848152602080526040902060070155602654615700908363ffffffff61485b16565b6026556040805182815260208101849052815185926001600160a01b038816927f4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1929081900390910190a36157558484611dd3565b61575e8361357f565b50505050565b600081836157b35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561558f578181015183820152602001615577565b5060008385816157bf57fe5b0495945050505050565b6000818152601f602090815260408083206004810154868552928190529083206002810154600590920154600390910154848315615843576000878152601f6020526040812060068101546002909101546158299163ffffffff61506116565b905061583f86613e68838863ffffffff61506116565b9150505b6000821561589c576000888152601f6020526040902060030154615873908590613e68908663ffffffff61506116565b9050615899620f4240613e68615887611924565b8490620f42400363ffffffff61506116565b90505b6158ac828263ffffffff61485b16565b9998505050505050505050565b60008281526030602052604081205482108015906131e457506158db82615526565b60008481526030602052604090206001015411905092915050565b6001600160a01b03831660009081526031602090815260408083208584529091528120548210801590614119575061592d82615526565b6001600160a01b038516600090815260316020908152604080832087845290915290206001015411949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061599e57805160ff19168380011785556159cb565b828001600101855582156159cb579182015b828111156159cb5782518255916020019190600101906159b0565b506159d7929150615a84565b5090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b50805460018160011615610100020316600290046000825580601f10615a2957506119f7565b601f0160209004906000526020600020908101906119f79190615a84565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6114c891905b808211156159d75760008155600101615a8a56fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573737374616b65722773206c6f636b696e672077696c6c2066696e697368206669727374646f65736e2774206578697374206f7220616c72656164792075706772616465646d757374206c65617665206174206c65617374206d696e44656c65676174696f6e536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a723158201943f7c33c86200cbacc225c431ca52b91a5305e67be36c01f1d1862f7b4606d64736f6c634300050c0032"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Contract *ContractCaller) SfcAddressToStakerID(opts *bind.CallOpts, sfcAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "_sfcAddressToStakerID", sfcAddress)
	return *ret0, err
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Contract *ContractSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.SfcAddressToStakerID(&_Contract.CallOpts, sfcAddress)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Contract *ContractCallerSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.SfcAddressToStakerID(&_Contract.CallOpts, sfcAddress)
}

// CalcDelegationCompoundRewards is a free data retrieval call binding the contract method 0x9864183d.
//
// Solidity: function calcDelegationCompoundRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcDelegationCompoundRewards(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "calcDelegationCompoundRewards", delegator, toStakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcDelegationCompoundRewards is a free data retrieval call binding the contract method 0x9864183d.
//
// Solidity: function calcDelegationCompoundRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcDelegationCompoundRewards(delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationCompoundRewards(&_Contract.CallOpts, delegator, toStakerID, _fromEpoch, maxEpochs)
}

// CalcDelegationCompoundRewards is a free data retrieval call binding the contract method 0x9864183d.
//
// Solidity: function calcDelegationCompoundRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcDelegationCompoundRewards(delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationCompoundRewards(&_Contract.CallOpts, delegator, toStakerID, _fromEpoch, maxEpochs)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "calcDelegationRewards", delegator, toStakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, _fromEpoch, maxEpochs)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorCompoundRewards is a free data retrieval call binding the contract method 0x74240362.
//
// Solidity: function calcValidatorCompoundRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcValidatorCompoundRewards(opts *bind.CallOpts, stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "calcValidatorCompoundRewards", stakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcValidatorCompoundRewards is a free data retrieval call binding the contract method 0x74240362.
//
// Solidity: function calcValidatorCompoundRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcValidatorCompoundRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorCompoundRewards(&_Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorCompoundRewards is a free data retrieval call binding the contract method 0x74240362.
//
// Solidity: function calcValidatorCompoundRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcValidatorCompoundRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorCompoundRewards(&_Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "calcValidatorRewards", stakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "contractCommission")
	return *ret0, err
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "currentSealedEpoch")
	return *ret0, err
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) DelegationEarlyWithdrawalPenalty(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "delegationEarlyWithdrawalPenalty", arg0, arg1)
	return *ret0, err
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) DelegationEarlyWithdrawalPenalty(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.DelegationEarlyWithdrawalPenalty(&_Contract.CallOpts, arg0, arg1)
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) DelegationEarlyWithdrawalPenalty(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.DelegationEarlyWithdrawalPenalty(&_Contract.CallOpts, arg0, arg1)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "delegationLockPeriodEpochs")
	return *ret0, err
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "delegationLockPeriodTime")
	return *ret0, err
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address , uint256 ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCaller) Delegations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	ret := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "delegations", arg0, arg1)
	return *ret, err
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address , uint256 ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractSession) Delegations(arg0 common.Address, arg1 *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, arg0, arg1)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address , uint256 ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCallerSession) Delegations(arg0 common.Address, arg1 *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, arg0, arg1)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Contract *ContractCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "delegationsNum")
	return *ret0, err
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Contract *ContractSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Contract *ContractCallerSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "delegationsTotalAmount")
	return *ret0, err
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Contract *ContractCaller) EpochSnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	ret := new(struct {
		EndTime                *big.Int
		Duration               *big.Int
		EpochFee               *big.Int
		TotalBaseRewardWeight  *big.Int
		TotalTxRewardWeight    *big.Int
		BaseRewardPerSecond    *big.Int
		StakeTotalAmount       *big.Int
		DelegationsTotalAmount *big.Int
		TotalSupply            *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "epochSnapshots", arg0)
	return *ret, err
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Contract *ContractSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _Contract.Contract.EpochSnapshots(&_Contract.CallOpts, arg0)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Contract *ContractCallerSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _Contract.Contract.EpochSnapshots(&_Contract.CallOpts, arg0)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Contract *ContractCaller) EpochValidator(opts *bind.CallOpts, e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	ret := new(struct {
		StakeAmount      *big.Int
		DelegatedMe      *big.Int
		BaseRewardWeight *big.Int
		TxRewardWeight   *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "epochValidator", e, v)
	return *ret, err
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Contract *ContractSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _Contract.Contract.EpochValidator(&_Contract.CallOpts, e, v)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Contract *ContractCallerSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _Contract.Contract.EpochValidator(&_Contract.CallOpts, e, v)
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Contract *ContractCaller) FirstLockedUpEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "firstLockedUpEpoch")
	return *ret0, err
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Contract *ContractSession) FirstLockedUpEpoch() (*big.Int, error) {
	return _Contract.Contract.FirstLockedUpEpoch(&_Contract.CallOpts)
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) FirstLockedUpEpoch() (*big.Int, error) {
	return _Contract.Contract.FirstLockedUpEpoch(&_Contract.CallOpts)
}

// GetDelegationRewardRatio is a free data retrieval call binding the contract method 0x5573184d.
//
// Solidity: function getDelegationRewardRatio(address delegator, uint256 toStakerID) view returns(uint256)
func (_Contract *ContractCaller) GetDelegationRewardRatio(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getDelegationRewardRatio", delegator, toStakerID)
	return *ret0, err
}

// GetDelegationRewardRatio is a free data retrieval call binding the contract method 0x5573184d.
//
// Solidity: function getDelegationRewardRatio(address delegator, uint256 toStakerID) view returns(uint256)
func (_Contract *ContractSession) GetDelegationRewardRatio(delegator common.Address, toStakerID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDelegationRewardRatio(&_Contract.CallOpts, delegator, toStakerID)
}

// GetDelegationRewardRatio is a free data retrieval call binding the contract method 0x5573184d.
//
// Solidity: function getDelegationRewardRatio(address delegator, uint256 toStakerID) view returns(uint256)
func (_Contract *ContractCallerSession) GetDelegationRewardRatio(delegator common.Address, toStakerID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDelegationRewardRatio(&_Contract.CallOpts, delegator, toStakerID)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Contract *ContractCaller) GetStakerID(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getStakerID", addr)
	return *ret0, err
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Contract *ContractSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Contract *ContractCallerSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, addr)
}

// GetValidatorRewardRatio is a free data retrieval call binding the contract method 0x8e431b8d.
//
// Solidity: function getValidatorRewardRatio(uint256 stakerID) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorRewardRatio(opts *bind.CallOpts, stakerID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getValidatorRewardRatio", stakerID)
	return *ret0, err
}

// GetValidatorRewardRatio is a free data retrieval call binding the contract method 0x8e431b8d.
//
// Solidity: function getValidatorRewardRatio(uint256 stakerID) view returns(uint256)
func (_Contract *ContractSession) GetValidatorRewardRatio(stakerID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorRewardRatio(&_Contract.CallOpts, stakerID)
}

// GetValidatorRewardRatio is a free data retrieval call binding the contract method 0x8e431b8d.
//
// Solidity: function getValidatorRewardRatio(uint256 stakerID) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorRewardRatio(stakerID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetValidatorRewardRatio(&_Contract.CallOpts, stakerID)
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCaller) IsDelegationLockedUp(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isDelegationLockedUp", delegator, toStakerID)
	return *ret0, err
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCallerSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 staker) view returns(bool)
func (_Contract *ContractCaller) IsStakeLockedUp(opts *bind.CallOpts, staker *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isStakeLockedUp", staker)
	return *ret0, err
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 staker) view returns(bool)
func (_Contract *ContractSession) IsStakeLockedUp(staker *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, staker)
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 staker) view returns(bool)
func (_Contract *ContractCallerSession) IsStakeLockedUp(staker *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, staker)
}

// LegacyDelegations is a free data retrieval call binding the contract method 0x5b81b886.
//
// Solidity: function legacyDelegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCaller) LegacyDelegations(opts *bind.CallOpts, arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	ret := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "legacyDelegations", arg0)
	return *ret, err
}

// LegacyDelegations is a free data retrieval call binding the contract method 0x5b81b886.
//
// Solidity: function legacyDelegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractSession) LegacyDelegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.LegacyDelegations(&_Contract.CallOpts, arg0)
}

// LegacyDelegations is a free data retrieval call binding the contract method 0x5b81b886.
//
// Solidity: function legacyDelegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCallerSession) LegacyDelegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.LegacyDelegations(&_Contract.CallOpts, arg0)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedDelegations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	ret := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "lockedDelegations", arg0, arg1)
	return *ret, err
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedDelegations(arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, arg0, arg1)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedDelegations(arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, arg0, arg1)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedStakes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	ret := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "lockedStakes", arg0)
	return *ret, err
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedStakes(arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, arg0)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedStakes(arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, arg0)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "maxDelegatedRatio")
	return *ret0, err
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MaxLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "maxLockupDuration")
	return *ret0, err
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Contract *ContractCaller) MaxStakerMetadataSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "maxStakerMetadataSize")
	return *ret0, err
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Contract *ContractSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _Contract.Contract.MaxStakerMetadataSize(&_Contract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _Contract.Contract.MaxStakerMetadataSize(&_Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minDelegationDecrease")
	return *ret0, err
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minDelegationIncrease")
	return *ret0, err
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MinLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minLockupDuration")
	return *ret0, err
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minStakeDecrease")
	return *ret0, err
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "minStakeIncrease")
	return *ret0, err
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Contract *ContractCaller) RewardsStash(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "rewardsStash", arg0, arg1)
	return *ret0, err
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Contract *ContractSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, arg0, arg1)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Contract *ContractCallerSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, arg0, arg1)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) SlashedDelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "slashedDelegationsTotalAmount")
	return *ret0, err
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Contract *ContractSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.SlashedDelegationsTotalAmount(&_Contract.CallOpts)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.SlashedDelegationsTotalAmount(&_Contract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) SlashedStakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "slashedStakeTotalAmount")
	return *ret0, err
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Contract *ContractSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.SlashedStakeTotalAmount(&_Contract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.SlashedStakeTotalAmount(&_Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakeLockPeriodEpochs")
	return *ret0, err
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakeLockPeriodTime")
	return *ret0, err
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakeTotalAmount")
	return *ret0, err
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Contract *ContractCaller) StakerMetadata(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakerMetadata", arg0)
	return *ret0, err
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Contract *ContractSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.StakerMetadata(&_Contract.CallOpts, arg0)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Contract *ContractCallerSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.StakerMetadata(&_Contract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	ret := new(struct {
		Status           *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		StakeAmount      *big.Int
		PaidUntilEpoch   *big.Int
		DelegatedMe      *big.Int
		DagAddress       common.Address
		SfcAddress       common.Address
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "stakers", arg0)
	return *ret, err
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCallerSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, arg0)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakersLastID")
	return *ret0, err
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCallerSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "stakersNum")
	return *ret0, err
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCallerSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Contract *ContractCaller) TotalBurntLockupRewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "totalBurntLockupRewards")
	return *ret0, err
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Contract *ContractSession) TotalBurntLockupRewards() (*big.Int, error) {
	return _Contract.Contract.TotalBurntLockupRewards(&_Contract.CallOpts)
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Contract *ContractCallerSession) TotalBurntLockupRewards() (*big.Int, error) {
	return _Contract.Contract.TotalBurntLockupRewards(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCaller) UnlockedRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "unlockedRewardRatio")
	return *ret0, err
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "validatorCommission")
	return *ret0, err
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) ([3]byte, error) {
	var (
		ret0 = new([3]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCallerSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Contract *ContractCaller) WithdrawalRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	ret := new(struct {
		StakerID   *big.Int
		Epoch      *big.Int
		Time       *big.Int
		Amount     *big.Int
		Delegation bool
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "withdrawalRequests", arg0, arg1)
	return *ret, err
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Contract *ContractSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _Contract.Contract.WithdrawalRequests(&_Contract.CallOpts, arg0, arg1)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Contract *ContractCallerSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _Contract.Contract.WithdrawalRequests(&_Contract.CallOpts, arg0, arg1)
}

// ActivateNetworkUpgrade is a paid mutator transaction binding the contract method 0xf5a83c7d.
//
// Solidity: function _activateNetworkUpgrade(uint256 minVersion) returns()
func (_Contract *ContractTransactor) ActivateNetworkUpgrade(opts *bind.TransactOpts, minVersion *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_activateNetworkUpgrade", minVersion)
}

// ActivateNetworkUpgrade is a paid mutator transaction binding the contract method 0xf5a83c7d.
//
// Solidity: function _activateNetworkUpgrade(uint256 minVersion) returns()
func (_Contract *ContractSession) ActivateNetworkUpgrade(minVersion *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ActivateNetworkUpgrade(&_Contract.TransactOpts, minVersion)
}

// ActivateNetworkUpgrade is a paid mutator transaction binding the contract method 0xf5a83c7d.
//
// Solidity: function _activateNetworkUpgrade(uint256 minVersion) returns()
func (_Contract *ContractTransactorSession) ActivateNetworkUpgrade(minVersion *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ActivateNetworkUpgrade(&_Contract.TransactOpts, minVersion)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x75b9d3d8.
//
// Solidity: function _syncDelegation(address delegator, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) SyncDelegation(opts *bind.TransactOpts, delegator common.Address, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_syncDelegation", delegator, toStakerID)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x75b9d3d8.
//
// Solidity: function _syncDelegation(address delegator, uint256 toStakerID) returns()
func (_Contract *ContractSession) SyncDelegation(delegator common.Address, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SyncDelegation(&_Contract.TransactOpts, delegator, toStakerID)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x75b9d3d8.
//
// Solidity: function _syncDelegation(address delegator, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) SyncDelegation(delegator common.Address, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SyncDelegation(&_Contract.TransactOpts, delegator, toStakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Contract *ContractTransactor) SyncStaker(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_syncStaker", stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Contract *ContractSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SyncStaker(&_Contract.TransactOpts, stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Contract *ContractTransactorSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SyncStaker(&_Contract.TransactOpts, stakerID)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x7b015db9.
//
// Solidity: function _updateBaseRewardPerSec(uint256 value) returns()
func (_Contract *ContractTransactor) UpdateBaseRewardPerSec(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_updateBaseRewardPerSec", value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x7b015db9.
//
// Solidity: function _updateBaseRewardPerSec(uint256 value) returns()
func (_Contract *ContractSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSec(&_Contract.TransactOpts, value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x7b015db9.
//
// Solidity: function _updateBaseRewardPerSec(uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSec(&_Contract.TransactOpts, value)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x1c3c60c8.
//
// Solidity: function _updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Contract *ContractTransactor) UpdateGasPowerAllocationRate(opts *bind.TransactOpts, short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_updateGasPowerAllocationRate", short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x1c3c60c8.
//
// Solidity: function _updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Contract *ContractSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateGasPowerAllocationRate(&_Contract.TransactOpts, short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x1c3c60c8.
//
// Solidity: function _updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Contract *ContractTransactorSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateGasPowerAllocationRate(&_Contract.TransactOpts, short, long)
}

// UpdateMinGasPrice is a paid mutator transaction binding the contract method 0xaa34eb45.
//
// Solidity: function _updateMinGasPrice(uint256 minGasPrice) returns()
func (_Contract *ContractTransactor) UpdateMinGasPrice(opts *bind.TransactOpts, minGasPrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_updateMinGasPrice", minGasPrice)
}

// UpdateMinGasPrice is a paid mutator transaction binding the contract method 0xaa34eb45.
//
// Solidity: function _updateMinGasPrice(uint256 minGasPrice) returns()
func (_Contract *ContractSession) UpdateMinGasPrice(minGasPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinGasPrice(&_Contract.TransactOpts, minGasPrice)
}

// UpdateMinGasPrice is a paid mutator transaction binding the contract method 0xaa34eb45.
//
// Solidity: function _updateMinGasPrice(uint256 minGasPrice) returns()
func (_Contract *ContractTransactorSession) UpdateMinGasPrice(minGasPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinGasPrice(&_Contract.TransactOpts, minGasPrice)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x2e5f75ef.
//
// Solidity: function _updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 period) returns()
func (_Contract *ContractTransactor) UpdateOfflinePenaltyThreshold(opts *bind.TransactOpts, blocksNum *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_updateOfflinePenaltyThreshold", blocksNum, period)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x2e5f75ef.
//
// Solidity: function _updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 period) returns()
func (_Contract *ContractSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, period)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x2e5f75ef.
//
// Solidity: function _updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 period) returns()
func (_Contract *ContractTransactorSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, period)
}

// UpgradeDelegationStorage is a paid mutator transaction binding the contract method 0x846ebb77.
//
// Solidity: function _upgradeDelegationStorage(address delegator) returns()
func (_Contract *ContractTransactor) UpgradeDelegationStorage(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_upgradeDelegationStorage", delegator)
}

// UpgradeDelegationStorage is a paid mutator transaction binding the contract method 0x846ebb77.
//
// Solidity: function _upgradeDelegationStorage(address delegator) returns()
func (_Contract *ContractSession) UpgradeDelegationStorage(delegator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeDelegationStorage(&_Contract.TransactOpts, delegator)
}

// UpgradeDelegationStorage is a paid mutator transaction binding the contract method 0x846ebb77.
//
// Solidity: function _upgradeDelegationStorage(address delegator) returns()
func (_Contract *ContractTransactorSession) UpgradeDelegationStorage(delegator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeDelegationStorage(&_Contract.TransactOpts, delegator)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Contract *ContractTransactor) UpgradeStakerStorage(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_upgradeStakerStorage", stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Contract *ContractSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeStakerStorage(&_Contract.TransactOpts, stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Contract *ContractTransactorSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeStakerStorage(&_Contract.TransactOpts, stakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationCompoundRewards(opts *bind.TransactOpts, maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationCompoundRewards", maxEpochs, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationCompoundRewards(maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, maxEpochs, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationCompoundRewards(maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, maxEpochs, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationRewards", maxEpochs, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationRewards(maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, maxEpochs, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationRewards(maxEpochs *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, maxEpochs, toStakerID)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 maxEpochs) returns()
func (_Contract *ContractTransactor) ClaimValidatorCompoundRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorCompoundRewards", maxEpochs)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 maxEpochs) returns()
func (_Contract *ContractSession) ClaimValidatorCompoundRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, maxEpochs)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 maxEpochs) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorCompoundRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Contract *ContractTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorRewards", maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Contract *ContractSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, maxEpochs)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Contract *ContractTransactor) CreateDelegation(opts *bind.TransactOpts, to *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createDelegation", to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Contract *ContractSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Contract *ContractTransactorSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, to)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Contract *ContractTransactor) CreateStake(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createStake", metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Contract *ContractSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateStake(&_Contract.TransactOpts, metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Contract *ContractTransactorSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateStake(&_Contract.TransactOpts, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Contract *ContractTransactor) CreateStakeWithAddresses(opts *bind.TransactOpts, dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createStakeWithAddresses", dagAddress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Contract *ContractSession) CreateStakeWithAddresses(dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateStakeWithAddresses(&_Contract.TransactOpts, dagAddress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Contract *ContractTransactorSession) CreateStakeWithAddresses(dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateStakeWithAddresses(&_Contract.TransactOpts, dagAddress, sfcAddress, metadata)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) LockUpDelegation(opts *bind.TransactOpts, lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpDelegation", lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactor) LockUpStake(opts *bind.TransactOpts, lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpStake", lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactorSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Contract *ContractTransactor) PartialWithdrawByRequest(opts *bind.TransactOpts, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "partialWithdrawByRequest", wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Contract *ContractSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Contract *ContractTransactorSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, wrID)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegation", toStakerID)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegation(toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, toStakerID)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegation(toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, toStakerID)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegationPartial(opts *bind.TransactOpts, wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegationPartial", wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStakePartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStakePartial", wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Contract *ContractTransactor) StartLockedUp(opts *bind.TransactOpts, epochNum *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "startLockedUp", epochNum)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Contract *ContractSession) StartLockedUp(epochNum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StartLockedUp(&_Contract.TransactOpts, epochNum)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Contract *ContractTransactorSession) StartLockedUp(epochNum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StartLockedUp(&_Contract.TransactOpts, epochNum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Contract *ContractTransactor) UnstashRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unstashRewards")
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Contract *ContractSession) UnstashRewards() (*types.Transaction, error) {
	return _Contract.Contract.UnstashRewards(&_Contract.TransactOpts)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Contract *ContractTransactorSession) UnstashRewards() (*types.Transaction, error) {
	return _Contract.Contract.UnstashRewards(&_Contract.TransactOpts)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Contract *ContractTransactor) UpdateStakerMetadata(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateStakerMetadata", metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Contract *ContractSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakerMetadata(&_Contract.TransactOpts, metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Contract *ContractTransactorSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakerMetadata(&_Contract.TransactOpts, metadata)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Contract *ContractTransactor) UpdateStakerSfcAddress(opts *bind.TransactOpts, newSfcAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateStakerSfcAddress", newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Contract *ContractSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakerSfcAddress(&_Contract.TransactOpts, newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Contract *ContractTransactorSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakerSfcAddress(&_Contract.TransactOpts, newSfcAddress)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractTransactor) WithdrawDelegation(opts *bind.TransactOpts, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawDelegation", toStakerID)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractSession) WithdrawDelegation(toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, toStakerID)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) WithdrawDelegation(toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, toStakerID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// ContractClaimedDelegationRewardIterator is returned from FilterClaimedDelegationReward and is used to iterate over the raw logs and unpacked data for ClaimedDelegationReward events raised by the Contract contract.
type ContractClaimedDelegationRewardIterator struct {
	Event *ContractClaimedDelegationReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedDelegationRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimedDelegationReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimedDelegationReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedDelegationRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedDelegationRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimedDelegationReward represents a ClaimedDelegationReward event raised by the Contract contract.
type ContractClaimedDelegationReward struct {
	From       common.Address
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedDelegationReward is a free log retrieval operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) FilterClaimedDelegationReward(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*ContractClaimedDelegationRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedDelegationRewardIterator{contract: _Contract.contract, event: "ClaimedDelegationReward", logs: logs, sub: sub}, nil
}

// WatchClaimedDelegationReward is a free log subscription operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) WatchClaimedDelegationReward(opts *bind.WatchOpts, sink chan<- *ContractClaimedDelegationReward, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimedDelegationReward)
				if err := _Contract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedDelegationReward is a log parse operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) ParseClaimedDelegationReward(log types.Log) (*ContractClaimedDelegationReward, error) {
	event := new(ContractClaimedDelegationReward)
	if err := _Contract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractClaimedValidatorRewardIterator is returned from FilterClaimedValidatorReward and is used to iterate over the raw logs and unpacked data for ClaimedValidatorReward events raised by the Contract contract.
type ContractClaimedValidatorRewardIterator struct {
	Event *ContractClaimedValidatorReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedValidatorRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimedValidatorReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimedValidatorReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedValidatorRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedValidatorRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimedValidatorReward represents a ClaimedValidatorReward event raised by the Contract contract.
type ContractClaimedValidatorReward struct {
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedValidatorReward is a free log retrieval operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) FilterClaimedValidatorReward(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractClaimedValidatorRewardIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedValidatorRewardIterator{contract: _Contract.contract, event: "ClaimedValidatorReward", logs: logs, sub: sub}, nil
}

// WatchClaimedValidatorReward is a free log subscription operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) WatchClaimedValidatorReward(opts *bind.WatchOpts, sink chan<- *ContractClaimedValidatorReward, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimedValidatorReward)
				if err := _Contract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedValidatorReward is a log parse operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Contract *ContractFilterer) ParseClaimedValidatorReward(log types.Log) (*ContractClaimedValidatorReward, error) {
	event := new(ContractClaimedValidatorReward)
	if err := _Contract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractCreatedDelegationIterator is returned from FilterCreatedDelegation and is used to iterate over the raw logs and unpacked data for CreatedDelegation events raised by the Contract contract.
type ContractCreatedDelegationIterator struct {
	Event *ContractCreatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedDelegation represents a CreatedDelegation event raised by the Contract contract.
type ContractCreatedDelegation struct {
	Delegator  common.Address
	ToStakerID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedDelegation is a free log retrieval operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Contract *ContractFilterer) FilterCreatedDelegation(opts *bind.FilterOpts, delegator []common.Address, toStakerID []*big.Int) (*ContractCreatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedDelegationIterator{contract: _Contract.contract, event: "CreatedDelegation", logs: logs, sub: sub}, nil
}

// WatchCreatedDelegation is a free log subscription operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Contract *ContractFilterer) WatchCreatedDelegation(opts *bind.WatchOpts, sink chan<- *ContractCreatedDelegation, delegator []common.Address, toStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedDelegation)
				if err := _Contract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedDelegation is a log parse operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Contract *ContractFilterer) ParseCreatedDelegation(log types.Log) (*ContractCreatedDelegation, error) {
	event := new(ContractCreatedDelegation)
	if err := _Contract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractCreatedStakeIterator is returned from FilterCreatedStake and is used to iterate over the raw logs and unpacked data for CreatedStake events raised by the Contract contract.
type ContractCreatedStakeIterator struct {
	Event *ContractCreatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedStake represents a CreatedStake event raised by the Contract contract.
type ContractCreatedStake struct {
	StakerID      *big.Int
	DagSfcAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreatedStake is a free log retrieval operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Contract *ContractFilterer) FilterCreatedStake(opts *bind.FilterOpts, stakerID []*big.Int, dagSfcAddress []common.Address) (*ContractCreatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedStakeIterator{contract: _Contract.contract, event: "CreatedStake", logs: logs, sub: sub}, nil
}

// WatchCreatedStake is a free log subscription operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Contract *ContractFilterer) WatchCreatedStake(opts *bind.WatchOpts, sink chan<- *ContractCreatedStake, stakerID []*big.Int, dagSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedStake)
				if err := _Contract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedStake is a log parse operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Contract *ContractFilterer) ParseCreatedStake(log types.Log) (*ContractCreatedStake, error) {
	event := new(ContractCreatedStake)
	if err := _Contract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractCreatedWithdrawRequestIterator is returned from FilterCreatedWithdrawRequest and is used to iterate over the raw logs and unpacked data for CreatedWithdrawRequest events raised by the Contract contract.
type ContractCreatedWithdrawRequestIterator struct {
	Event *ContractCreatedWithdrawRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedWithdrawRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedWithdrawRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedWithdrawRequest represents a CreatedWithdrawRequest event raised by the Contract contract.
type ContractCreatedWithdrawRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedWithdrawRequest is a free log retrieval operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Contract *ContractFilterer) FilterCreatedWithdrawRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*ContractCreatedWithdrawRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedWithdrawRequestIterator{contract: _Contract.contract, event: "CreatedWithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchCreatedWithdrawRequest is a free log subscription operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Contract *ContractFilterer) WatchCreatedWithdrawRequest(opts *bind.WatchOpts, sink chan<- *ContractCreatedWithdrawRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedWithdrawRequest)
				if err := _Contract.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedWithdrawRequest is a log parse operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Contract *ContractFilterer) ParseCreatedWithdrawRequest(log types.Log) (*ContractCreatedWithdrawRequest, error) {
	event := new(ContractCreatedWithdrawRequest)
	if err := _Contract.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractDeactivatedDelegationIterator is returned from FilterDeactivatedDelegation and is used to iterate over the raw logs and unpacked data for DeactivatedDelegation events raised by the Contract contract.
type ContractDeactivatedDelegationIterator struct {
	Event *ContractDeactivatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeactivatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeactivatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeactivatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeactivatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeactivatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeactivatedDelegation represents a DeactivatedDelegation event raised by the Contract contract.
type ContractDeactivatedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedDelegation is a free log retrieval operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) FilterDeactivatedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*ContractDeactivatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeactivatedDelegationIterator{contract: _Contract.contract, event: "DeactivatedDelegation", logs: logs, sub: sub}, nil
}

// WatchDeactivatedDelegation is a free log subscription operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) WatchDeactivatedDelegation(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeactivatedDelegation)
				if err := _Contract.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedDelegation is a log parse operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) ParseDeactivatedDelegation(log types.Log) (*ContractDeactivatedDelegation, error) {
	event := new(ContractDeactivatedDelegation)
	if err := _Contract.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractDeactivatedStakeIterator is returned from FilterDeactivatedStake and is used to iterate over the raw logs and unpacked data for DeactivatedStake events raised by the Contract contract.
type ContractDeactivatedStakeIterator struct {
	Event *ContractDeactivatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeactivatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeactivatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeactivatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeactivatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeactivatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeactivatedStake represents a DeactivatedStake event raised by the Contract contract.
type ContractDeactivatedStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedStake is a free log retrieval operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) FilterDeactivatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractDeactivatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeactivatedStakeIterator{contract: _Contract.contract, event: "DeactivatedStake", logs: logs, sub: sub}, nil
}

// WatchDeactivatedStake is a free log subscription operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) WatchDeactivatedStake(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeactivatedStake)
				if err := _Contract.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedStake is a log parse operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) ParseDeactivatedStake(log types.Log) (*ContractDeactivatedStake, error) {
	event := new(ContractDeactivatedStake)
	if err := _Contract.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractIncreasedDelegationIterator is returned from FilterIncreasedDelegation and is used to iterate over the raw logs and unpacked data for IncreasedDelegation events raised by the Contract contract.
type ContractIncreasedDelegationIterator struct {
	Event *ContractIncreasedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncreasedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncreasedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncreasedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncreasedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncreasedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncreasedDelegation represents a IncreasedDelegation event raised by the Contract contract.
type ContractIncreasedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedDelegation is a free log retrieval operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) FilterIncreasedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*ContractIncreasedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncreasedDelegationIterator{contract: _Contract.contract, event: "IncreasedDelegation", logs: logs, sub: sub}, nil
}

// WatchIncreasedDelegation is a free log subscription operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) WatchIncreasedDelegation(opts *bind.WatchOpts, sink chan<- *ContractIncreasedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncreasedDelegation)
				if err := _Contract.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasedDelegation is a log parse operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) ParseIncreasedDelegation(log types.Log) (*ContractIncreasedDelegation, error) {
	event := new(ContractIncreasedDelegation)
	if err := _Contract.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractIncreasedStakeIterator is returned from FilterIncreasedStake and is used to iterate over the raw logs and unpacked data for IncreasedStake events raised by the Contract contract.
type ContractIncreasedStakeIterator struct {
	Event *ContractIncreasedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncreasedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncreasedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncreasedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncreasedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncreasedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncreasedStake represents a IncreasedStake event raised by the Contract contract.
type ContractIncreasedStake struct {
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedStake is a free log retrieval operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) FilterIncreasedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractIncreasedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncreasedStakeIterator{contract: _Contract.contract, event: "IncreasedStake", logs: logs, sub: sub}, nil
}

// WatchIncreasedStake is a free log subscription operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) WatchIncreasedStake(opts *bind.WatchOpts, sink chan<- *ContractIncreasedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncreasedStake)
				if err := _Contract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasedStake is a log parse operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Contract *ContractFilterer) ParseIncreasedStake(log types.Log) (*ContractIncreasedStake, error) {
	event := new(ContractIncreasedStake)
	if err := _Contract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractLockingDelegationIterator is returned from FilterLockingDelegation and is used to iterate over the raw logs and unpacked data for LockingDelegation events raised by the Contract contract.
type ContractLockingDelegationIterator struct {
	Event *ContractLockingDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLockingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLockingDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLockingDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLockingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLockingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLockingDelegation represents a LockingDelegation event raised by the Contract contract.
type ContractLockingDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	FromEpoch *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockingDelegation is a free log retrieval operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) FilterLockingDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*ContractLockingDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockingDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractLockingDelegationIterator{contract: _Contract.contract, event: "LockingDelegation", logs: logs, sub: sub}, nil
}

// WatchLockingDelegation is a free log subscription operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) WatchLockingDelegation(opts *bind.WatchOpts, sink chan<- *ContractLockingDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockingDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLockingDelegation)
				if err := _Contract.contract.UnpackLog(event, "LockingDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockingDelegation is a log parse operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) ParseLockingDelegation(log types.Log) (*ContractLockingDelegation, error) {
	event := new(ContractLockingDelegation)
	if err := _Contract.contract.UnpackLog(event, "LockingDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractLockingStakeIterator is returned from FilterLockingStake and is used to iterate over the raw logs and unpacked data for LockingStake events raised by the Contract contract.
type ContractLockingStakeIterator struct {
	Event *ContractLockingStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLockingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLockingStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLockingStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLockingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLockingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLockingStake represents a LockingStake event raised by the Contract contract.
type ContractLockingStake struct {
	StakerID  *big.Int
	FromEpoch *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockingStake is a free log retrieval operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) FilterLockingStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractLockingStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockingStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractLockingStakeIterator{contract: _Contract.contract, event: "LockingStake", logs: logs, sub: sub}, nil
}

// WatchLockingStake is a free log subscription operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) WatchLockingStake(opts *bind.WatchOpts, sink chan<- *ContractLockingStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockingStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLockingStake)
				if err := _Contract.contract.UnpackLog(event, "LockingStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockingStake is a log parse operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Contract *ContractFilterer) ParseLockingStake(log types.Log) (*ContractLockingStake, error) {
	event := new(ContractLockingStake)
	if err := _Contract.contract.UnpackLog(event, "LockingStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractNetworkUpgradeActivatedIterator is returned from FilterNetworkUpgradeActivated and is used to iterate over the raw logs and unpacked data for NetworkUpgradeActivated events raised by the Contract contract.
type ContractNetworkUpgradeActivatedIterator struct {
	Event *ContractNetworkUpgradeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNetworkUpgradeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNetworkUpgradeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNetworkUpgradeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNetworkUpgradeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNetworkUpgradeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNetworkUpgradeActivated represents a NetworkUpgradeActivated event raised by the Contract contract.
type ContractNetworkUpgradeActivated struct {
	MinVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNetworkUpgradeActivated is a free log retrieval operation binding the contract event 0xa3deceaa35ccc5aa4f1e61ffe83a006792b8989d4e80dd2c8aa07ba8a923cde1.
//
// Solidity: event NetworkUpgradeActivated(uint256 minVersion)
func (_Contract *ContractFilterer) FilterNetworkUpgradeActivated(opts *bind.FilterOpts) (*ContractNetworkUpgradeActivatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NetworkUpgradeActivated")
	if err != nil {
		return nil, err
	}
	return &ContractNetworkUpgradeActivatedIterator{contract: _Contract.contract, event: "NetworkUpgradeActivated", logs: logs, sub: sub}, nil
}

// WatchNetworkUpgradeActivated is a free log subscription operation binding the contract event 0xa3deceaa35ccc5aa4f1e61ffe83a006792b8989d4e80dd2c8aa07ba8a923cde1.
//
// Solidity: event NetworkUpgradeActivated(uint256 minVersion)
func (_Contract *ContractFilterer) WatchNetworkUpgradeActivated(opts *bind.WatchOpts, sink chan<- *ContractNetworkUpgradeActivated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NetworkUpgradeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNetworkUpgradeActivated)
				if err := _Contract.contract.UnpackLog(event, "NetworkUpgradeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNetworkUpgradeActivated is a log parse operation binding the contract event 0xa3deceaa35ccc5aa4f1e61ffe83a006792b8989d4e80dd2c8aa07ba8a923cde1.
//
// Solidity: event NetworkUpgradeActivated(uint256 minVersion)
func (_Contract *ContractFilterer) ParseNetworkUpgradeActivated(log types.Log) (*ContractNetworkUpgradeActivated, error) {
	event := new(ContractNetworkUpgradeActivated)
	if err := _Contract.contract.UnpackLog(event, "NetworkUpgradeActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractPartialWithdrawnByRequestIterator is returned from FilterPartialWithdrawnByRequest and is used to iterate over the raw logs and unpacked data for PartialWithdrawnByRequest events raised by the Contract contract.
type ContractPartialWithdrawnByRequestIterator struct {
	Event *ContractPartialWithdrawnByRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPartialWithdrawnByRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPartialWithdrawnByRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPartialWithdrawnByRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPartialWithdrawnByRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPartialWithdrawnByRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPartialWithdrawnByRequest represents a PartialWithdrawnByRequest event raised by the Contract contract.
type ContractPartialWithdrawnByRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Penalty    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawnByRequest is a free log retrieval operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Contract *ContractFilterer) FilterPartialWithdrawnByRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*ContractPartialWithdrawnByRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractPartialWithdrawnByRequestIterator{contract: _Contract.contract, event: "PartialWithdrawnByRequest", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawnByRequest is a free log subscription operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Contract *ContractFilterer) WatchPartialWithdrawnByRequest(opts *bind.WatchOpts, sink chan<- *ContractPartialWithdrawnByRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPartialWithdrawnByRequest)
				if err := _Contract.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePartialWithdrawnByRequest is a log parse operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Contract *ContractFilterer) ParsePartialWithdrawnByRequest(log types.Log) (*ContractPartialWithdrawnByRequest, error) {
	event := new(ContractPartialWithdrawnByRequest)
	if err := _Contract.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractPreparedToWithdrawDelegationIterator is returned from FilterPreparedToWithdrawDelegation and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawDelegation events raised by the Contract contract.
type ContractPreparedToWithdrawDelegationIterator struct {
	Event *ContractPreparedToWithdrawDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPreparedToWithdrawDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPreparedToWithdrawDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPreparedToWithdrawDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPreparedToWithdrawDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPreparedToWithdrawDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPreparedToWithdrawDelegation represents a PreparedToWithdrawDelegation event raised by the Contract contract.
type ContractPreparedToWithdrawDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawDelegation is a free log retrieval operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) FilterPreparedToWithdrawDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*ContractPreparedToWithdrawDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractPreparedToWithdrawDelegationIterator{contract: _Contract.contract, event: "PreparedToWithdrawDelegation", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawDelegation is a free log subscription operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) WatchPreparedToWithdrawDelegation(opts *bind.WatchOpts, sink chan<- *ContractPreparedToWithdrawDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPreparedToWithdrawDelegation)
				if err := _Contract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreparedToWithdrawDelegation is a log parse operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Contract *ContractFilterer) ParsePreparedToWithdrawDelegation(log types.Log) (*ContractPreparedToWithdrawDelegation, error) {
	event := new(ContractPreparedToWithdrawDelegation)
	if err := _Contract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractPreparedToWithdrawStakeIterator is returned from FilterPreparedToWithdrawStake and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawStake events raised by the Contract contract.
type ContractPreparedToWithdrawStakeIterator struct {
	Event *ContractPreparedToWithdrawStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPreparedToWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPreparedToWithdrawStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPreparedToWithdrawStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPreparedToWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPreparedToWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPreparedToWithdrawStake represents a PreparedToWithdrawStake event raised by the Contract contract.
type ContractPreparedToWithdrawStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawStake is a free log retrieval operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) FilterPreparedToWithdrawStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractPreparedToWithdrawStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractPreparedToWithdrawStakeIterator{contract: _Contract.contract, event: "PreparedToWithdrawStake", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawStake is a free log subscription operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) WatchPreparedToWithdrawStake(opts *bind.WatchOpts, sink chan<- *ContractPreparedToWithdrawStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPreparedToWithdrawStake)
				if err := _Contract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreparedToWithdrawStake is a log parse operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Contract *ContractFilterer) ParsePreparedToWithdrawStake(log types.Log) (*ContractPreparedToWithdrawStake, error) {
	event := new(ContractPreparedToWithdrawStake)
	if err := _Contract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUnstashedRewardsIterator is returned from FilterUnstashedRewards and is used to iterate over the raw logs and unpacked data for UnstashedRewards events raised by the Contract contract.
type ContractUnstashedRewardsIterator struct {
	Event *ContractUnstashedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnstashedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnstashedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnstashedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnstashedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnstashedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnstashedRewards represents a UnstashedRewards event raised by the Contract contract.
type ContractUnstashedRewards struct {
	Auth     common.Address
	Receiver common.Address
	Rewards  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstashedRewards is a free log retrieval operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Contract *ContractFilterer) FilterUnstashedRewards(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address) (*ContractUnstashedRewardsIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnstashedRewardsIterator{contract: _Contract.contract, event: "UnstashedRewards", logs: logs, sub: sub}, nil
}

// WatchUnstashedRewards is a free log subscription operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Contract *ContractFilterer) WatchUnstashedRewards(opts *bind.WatchOpts, sink chan<- *ContractUnstashedRewards, auth []common.Address, receiver []common.Address) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnstashedRewards)
				if err := _Contract.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstashedRewards is a log parse operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Contract *ContractFilterer) ParseUnstashedRewards(log types.Log) (*ContractUnstashedRewards, error) {
	event := new(ContractUnstashedRewards)
	if err := _Contract.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the Contract contract.
type ContractUpdatedBaseRewardPerSecIterator struct {
	Event *ContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedBaseRewardPerSec)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedBaseRewardPerSec)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the Contract contract.
type ContractUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*ContractUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedBaseRewardPerSecIterator{contract: _Contract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedBaseRewardPerSec)
				if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*ContractUpdatedBaseRewardPerSec, error) {
	event := new(ContractUpdatedBaseRewardPerSec)
	if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedDelegationIterator is returned from FilterUpdatedDelegation and is used to iterate over the raw logs and unpacked data for UpdatedDelegation events raised by the Contract contract.
type ContractUpdatedDelegationIterator struct {
	Event *ContractUpdatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedDelegation represents a UpdatedDelegation event raised by the Contract contract.
type ContractUpdatedDelegation struct {
	Delegator   common.Address
	OldStakerID *big.Int
	NewStakerID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDelegation is a free log retrieval operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Contract *ContractFilterer) FilterUpdatedDelegation(opts *bind.FilterOpts, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (*ContractUpdatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedDelegationIterator{contract: _Contract.contract, event: "UpdatedDelegation", logs: logs, sub: sub}, nil
}

// WatchUpdatedDelegation is a free log subscription operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Contract *ContractFilterer) WatchUpdatedDelegation(opts *bind.WatchOpts, sink chan<- *ContractUpdatedDelegation, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedDelegation)
				if err := _Contract.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedDelegation is a log parse operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Contract *ContractFilterer) ParseUpdatedDelegation(log types.Log) (*ContractUpdatedDelegation, error) {
	event := new(ContractUpdatedDelegation)
	if err := _Contract.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedGasPowerAllocationRateIterator is returned from FilterUpdatedGasPowerAllocationRate and is used to iterate over the raw logs and unpacked data for UpdatedGasPowerAllocationRate events raised by the Contract contract.
type ContractUpdatedGasPowerAllocationRateIterator struct {
	Event *ContractUpdatedGasPowerAllocationRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedGasPowerAllocationRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedGasPowerAllocationRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedGasPowerAllocationRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedGasPowerAllocationRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedGasPowerAllocationRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedGasPowerAllocationRate represents a UpdatedGasPowerAllocationRate event raised by the Contract contract.
type ContractUpdatedGasPowerAllocationRate struct {
	Short *big.Int
	Long  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasPowerAllocationRate is a free log retrieval operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Contract *ContractFilterer) FilterUpdatedGasPowerAllocationRate(opts *bind.FilterOpts) (*ContractUpdatedGasPowerAllocationRateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedGasPowerAllocationRateIterator{contract: _Contract.contract, event: "UpdatedGasPowerAllocationRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasPowerAllocationRate is a free log subscription operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Contract *ContractFilterer) WatchUpdatedGasPowerAllocationRate(opts *bind.WatchOpts, sink chan<- *ContractUpdatedGasPowerAllocationRate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedGasPowerAllocationRate)
				if err := _Contract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGasPowerAllocationRate is a log parse operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Contract *ContractFilterer) ParseUpdatedGasPowerAllocationRate(log types.Log) (*ContractUpdatedGasPowerAllocationRate, error) {
	event := new(ContractUpdatedGasPowerAllocationRate)
	if err := _Contract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedMinGasPriceIterator is returned from FilterUpdatedMinGasPrice and is used to iterate over the raw logs and unpacked data for UpdatedMinGasPrice events raised by the Contract contract.
type ContractUpdatedMinGasPriceIterator struct {
	Event *ContractUpdatedMinGasPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedMinGasPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedMinGasPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedMinGasPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedMinGasPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedMinGasPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedMinGasPrice represents a UpdatedMinGasPrice event raised by the Contract contract.
type ContractUpdatedMinGasPrice struct {
	MinGasPrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMinGasPrice is a free log retrieval operation binding the contract event 0x35feeeac858525cae277d98c1c4792d0550aeab30f107addc09d8d5279faa53f.
//
// Solidity: event UpdatedMinGasPrice(uint256 minGasPrice)
func (_Contract *ContractFilterer) FilterUpdatedMinGasPrice(opts *bind.FilterOpts) (*ContractUpdatedMinGasPriceIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedMinGasPrice")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedMinGasPriceIterator{contract: _Contract.contract, event: "UpdatedMinGasPrice", logs: logs, sub: sub}, nil
}

// WatchUpdatedMinGasPrice is a free log subscription operation binding the contract event 0x35feeeac858525cae277d98c1c4792d0550aeab30f107addc09d8d5279faa53f.
//
// Solidity: event UpdatedMinGasPrice(uint256 minGasPrice)
func (_Contract *ContractFilterer) WatchUpdatedMinGasPrice(opts *bind.WatchOpts, sink chan<- *ContractUpdatedMinGasPrice) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedMinGasPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedMinGasPrice)
				if err := _Contract.contract.UnpackLog(event, "UpdatedMinGasPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedMinGasPrice is a log parse operation binding the contract event 0x35feeeac858525cae277d98c1c4792d0550aeab30f107addc09d8d5279faa53f.
//
// Solidity: event UpdatedMinGasPrice(uint256 minGasPrice)
func (_Contract *ContractFilterer) ParseUpdatedMinGasPrice(log types.Log) (*ContractUpdatedMinGasPrice, error) {
	event := new(ContractUpdatedMinGasPrice)
	if err := _Contract.contract.UnpackLog(event, "UpdatedMinGasPrice", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedOfflinePenaltyThresholdIterator is returned from FilterUpdatedOfflinePenaltyThreshold and is used to iterate over the raw logs and unpacked data for UpdatedOfflinePenaltyThreshold events raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThresholdIterator struct {
	Event *ContractUpdatedOfflinePenaltyThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedOfflinePenaltyThreshold represents a UpdatedOfflinePenaltyThreshold event raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThreshold struct {
	BlocksNum *big.Int
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOfflinePenaltyThreshold is a free log retrieval operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) FilterUpdatedOfflinePenaltyThreshold(opts *bind.FilterOpts) (*ContractUpdatedOfflinePenaltyThresholdIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedOfflinePenaltyThresholdIterator{contract: _Contract.contract, event: "UpdatedOfflinePenaltyThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedOfflinePenaltyThreshold is a free log subscription operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) WatchUpdatedOfflinePenaltyThreshold(opts *bind.WatchOpts, sink chan<- *ContractUpdatedOfflinePenaltyThreshold) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedOfflinePenaltyThreshold)
				if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedOfflinePenaltyThreshold is a log parse operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) ParseUpdatedOfflinePenaltyThreshold(log types.Log) (*ContractUpdatedOfflinePenaltyThreshold, error) {
	event := new(ContractUpdatedOfflinePenaltyThreshold)
	if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedStakeIterator is returned from FilterUpdatedStake and is used to iterate over the raw logs and unpacked data for UpdatedStake events raised by the Contract contract.
type ContractUpdatedStakeIterator struct {
	Event *ContractUpdatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedStake represents a UpdatedStake event raised by the Contract contract.
type ContractUpdatedStake struct {
	StakerID    *big.Int
	Amount      *big.Int
	DelegatedMe *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStake is a free log retrieval operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Contract *ContractFilterer) FilterUpdatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractUpdatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedStakeIterator{contract: _Contract.contract, event: "UpdatedStake", logs: logs, sub: sub}, nil
}

// WatchUpdatedStake is a free log subscription operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Contract *ContractFilterer) WatchUpdatedStake(opts *bind.WatchOpts, sink chan<- *ContractUpdatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedStake)
				if err := _Contract.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStake is a log parse operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Contract *ContractFilterer) ParseUpdatedStake(log types.Log) (*ContractUpdatedStake, error) {
	event := new(ContractUpdatedStake)
	if err := _Contract.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedStakerMetadataIterator is returned from FilterUpdatedStakerMetadata and is used to iterate over the raw logs and unpacked data for UpdatedStakerMetadata events raised by the Contract contract.
type ContractUpdatedStakerMetadataIterator struct {
	Event *ContractUpdatedStakerMetadata // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedStakerMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedStakerMetadata)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedStakerMetadata)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedStakerMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedStakerMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedStakerMetadata represents a UpdatedStakerMetadata event raised by the Contract contract.
type ContractUpdatedStakerMetadata struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerMetadata is a free log retrieval operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Contract *ContractFilterer) FilterUpdatedStakerMetadata(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractUpdatedStakerMetadataIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedStakerMetadataIterator{contract: _Contract.contract, event: "UpdatedStakerMetadata", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerMetadata is a free log subscription operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Contract *ContractFilterer) WatchUpdatedStakerMetadata(opts *bind.WatchOpts, sink chan<- *ContractUpdatedStakerMetadata, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedStakerMetadata)
				if err := _Contract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStakerMetadata is a log parse operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Contract *ContractFilterer) ParseUpdatedStakerMetadata(log types.Log) (*ContractUpdatedStakerMetadata, error) {
	event := new(ContractUpdatedStakerMetadata)
	if err := _Contract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractUpdatedStakerSfcAddressIterator is returned from FilterUpdatedStakerSfcAddress and is used to iterate over the raw logs and unpacked data for UpdatedStakerSfcAddress events raised by the Contract contract.
type ContractUpdatedStakerSfcAddressIterator struct {
	Event *ContractUpdatedStakerSfcAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedStakerSfcAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedStakerSfcAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedStakerSfcAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedStakerSfcAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedStakerSfcAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedStakerSfcAddress represents a UpdatedStakerSfcAddress event raised by the Contract contract.
type ContractUpdatedStakerSfcAddress struct {
	StakerID      *big.Int
	OldSfcAddress common.Address
	NewSfcAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerSfcAddress is a free log retrieval operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Contract *ContractFilterer) FilterUpdatedStakerSfcAddress(opts *bind.FilterOpts, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (*ContractUpdatedStakerSfcAddressIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedStakerSfcAddressIterator{contract: _Contract.contract, event: "UpdatedStakerSfcAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerSfcAddress is a free log subscription operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Contract *ContractFilterer) WatchUpdatedStakerSfcAddress(opts *bind.WatchOpts, sink chan<- *ContractUpdatedStakerSfcAddress, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedStakerSfcAddress)
				if err := _Contract.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStakerSfcAddress is a log parse operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Contract *ContractFilterer) ParseUpdatedStakerSfcAddress(log types.Log) (*ContractUpdatedStakerSfcAddress, error) {
	event := new(ContractUpdatedStakerSfcAddress)
	if err := _Contract.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractWithdrawnDelegationIterator is returned from FilterWithdrawnDelegation and is used to iterate over the raw logs and unpacked data for WithdrawnDelegation events raised by the Contract contract.
type ContractWithdrawnDelegationIterator struct {
	Event *ContractWithdrawnDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWithdrawnDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawnDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWithdrawnDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWithdrawnDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawnDelegation represents a WithdrawnDelegation event raised by the Contract contract.
type ContractWithdrawnDelegation struct {
	Delegator  common.Address
	ToStakerID *big.Int
	Penalty    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnDelegation is a free log retrieval operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 penalty)
func (_Contract *ContractFilterer) FilterWithdrawnDelegation(opts *bind.FilterOpts, delegator []common.Address, toStakerID []*big.Int) (*ContractWithdrawnDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "WithdrawnDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnDelegationIterator{contract: _Contract.contract, event: "WithdrawnDelegation", logs: logs, sub: sub}, nil
}

// WatchWithdrawnDelegation is a free log subscription operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 penalty)
func (_Contract *ContractFilterer) WatchWithdrawnDelegation(opts *bind.WatchOpts, sink chan<- *ContractWithdrawnDelegation, delegator []common.Address, toStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "WithdrawnDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawnDelegation)
				if err := _Contract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnDelegation is a log parse operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 penalty)
func (_Contract *ContractFilterer) ParseWithdrawnDelegation(log types.Log) (*ContractWithdrawnDelegation, error) {
	event := new(ContractWithdrawnDelegation)
	if err := _Contract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractWithdrawnStakeIterator is returned from FilterWithdrawnStake and is used to iterate over the raw logs and unpacked data for WithdrawnStake events raised by the Contract contract.
type ContractWithdrawnStakeIterator struct {
	Event *ContractWithdrawnStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWithdrawnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawnStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWithdrawnStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWithdrawnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawnStake represents a WithdrawnStake event raised by the Contract contract.
type ContractWithdrawnStake struct {
	StakerID *big.Int
	Penalty  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnStake is a free log retrieval operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Contract *ContractFilterer) FilterWithdrawnStake(opts *bind.FilterOpts, stakerID []*big.Int) (*ContractWithdrawnStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnStakeIterator{contract: _Contract.contract, event: "WithdrawnStake", logs: logs, sub: sub}, nil
}

// WatchWithdrawnStake is a free log subscription operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Contract *ContractFilterer) WatchWithdrawnStake(opts *bind.WatchOpts, sink chan<- *ContractWithdrawnStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawnStake)
				if err := _Contract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnStake is a log parse operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Contract *ContractFilterer) ParseWithdrawnStake(log types.Log) (*ContractWithdrawnStake, error) {
	event := new(ContractWithdrawnStake)
	if err := _Contract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
		return nil, err
	}
	return event, nil
}
