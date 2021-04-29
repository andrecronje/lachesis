package gossip

import (
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	confirmBlocksMeter = metrics.NewRegisteredCounter("confirm/blocks", nil)
	confirmTxnsMeter   = metrics.NewRegisteredCounter("confirm/transactions", nil)
	txTtfMeter         = metrics.NewRegisteredHistogram("tx_ttf", nil, metrics.NewUniformSample(500))
)

var txLatency = NewTxs()
