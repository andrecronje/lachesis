package main

import (
	"github.com/ethereum/go-ethereum/common"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-lachesis/crypto"
	"github.com/Fantom-foundation/go-lachesis/gossip"
)

var validatorFlag = cli.StringFlag{
	Name:  "validator",
	Usage: "Address of a validator to create events from",
	Value: "no",
}

func getValidatorAddr(ctx *cli.Context, cfg *gossip.EmitterConfig) common.Address {
	// Extract the current validator address, new flag overriding legacy one
	validator := cfg.Validator
	switch {
	case ctx.GlobalIsSet(validatorFlag.Name):
		validatorStr := ctx.GlobalString(validatorFlag.Name)
		if validatorStr != "no" && validatorStr != "0" {
			validator = common.HexToAddress(validatorStr)
		}
	case ctx.GlobalIsSet(FakeNetFlag.Name):
		key := getFakeValidator(ctx)
		if key != nil {
			validator = crypto.PubkeyToAddress(key.PublicKey)
		}
	}
	return validator
}

// setValidator retrieves the validator address either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setValidator(ctx *cli.Context, cfg *gossip.EmitterConfig) {
	cfg.Validator = getValidatorAddr(ctx, cfg)
}
