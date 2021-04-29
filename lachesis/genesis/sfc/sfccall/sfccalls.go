package sfccall

import (
	"github.com/Fantom-foundation/go-lachesis/gossip/sfc202"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-opera/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

const ContractABI = sfc202.ContractABI

var (
	sAbi, _ = abi.JSON(strings.NewReader(ContractABI))
)

// Methods

func CalcValidatorRewards(id idx.StakerID, start idx.Epoch, maxEpochs idx.Epoch) []byte {
	data, _ := sAbi.Pack("calcValidatorRewards", utils.U64toBig(uint64(id)), utils.U64toBig(uint64(start)), utils.U64toBig(uint64(maxEpochs)))
	return data
}

func CalcDelegationRewards(delegator common.Address, toStakerID idx.StakerID, start idx.Epoch, maxEpochs idx.Epoch) []byte {
	data, _ := sAbi.Pack("calcValidatorRewards", delegator, utils.U64toBig(uint64(toStakerID)), utils.U64toBig(uint64(start)), utils.U64toBig(uint64(maxEpochs)))
	return data
}
