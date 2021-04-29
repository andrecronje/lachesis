package integration

import (
	"time"

	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"

	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/lachesis"
)

// NewIntegration creates gossip service for the integration test
func NewIntegration(ctx *adapters.ServiceContext, network lachesis.Config, stack *node.Node) *gossip.Service {
	gossipCfg := gossip.DefaultConfig(network)

	engine, _, _, gdb := MakeEngine(ctx.Config.DataDir, &gossipCfg)

	coinbase := SetAccountKey(
		stack.AccountManager(),
		ctx.Config.PrivateKey,
		"fakepassword",
	)

	gossipCfg.Emitter.Validator = coinbase.Address
	gossipCfg.Emitter.EmitIntervals.Max = 3 * time.Second
	gossipCfg.Emitter.EmitIntervals.SelfForkProtection = 0

	svc, err := gossip.NewService(stack, &gossipCfg, gdb, engine)
	if err != nil {
		panic(err)
	}

	return svc
}
