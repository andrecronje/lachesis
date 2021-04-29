package gossip

import (
	"github.com/Fantom-foundation/go-lachesis/inter"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-lachesis/app"
	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

type EvmStateReader struct {
	*ServiceFeed
	engineMu    *sync.RWMutex
	engine      Consensus
	minGasPrice func() *big.Int

	store *Store
	app   *app.Store
}

func (s *Service) GetEvmStateReader() *EvmStateReader {
	return &EvmStateReader{
		ServiceFeed: &s.feed,
		engineMu:    s.engineMu,
		engine:      s.engine,
		minGasPrice: s.MinGasPrice,
		store:       s.store,
		app:         s.app,
	}
}

func (r *EvmStateReader) MinGasPrice() *big.Int {
	return r.minGasPrice()
}

func (r *EvmStateReader) CurrentBlock() *evmcore.EvmBlock {
	r.engineMu.RLock()
	n, h := r.engine.LastBlock()
	r.engineMu.RUnlock()

	return r.getBlock(hash.Event(h), idx.Block(n), n != 0)
}

func (r *EvmStateReader) CurrentHeader() *evmcore.EvmHeader {
	r.engineMu.RLock()
	n, h := r.engine.LastBlock()
	r.engineMu.RUnlock()

	return r.getBlock(hash.Event(h), idx.Block(n), false).Header()
}

func (r *EvmStateReader) GetHeader(h common.Hash, n uint64) *evmcore.EvmHeader {
	return r.GetDagHeader(hash.Event(h), idx.Block(n))
}

func (r *EvmStateReader) GetBlock(h common.Hash, n uint64) *evmcore.EvmBlock {
	return r.GetDagBlock(hash.Event(h), idx.Block(n))
}

func (r *EvmStateReader) GetDagHeader(h hash.Event, n idx.Block) *evmcore.EvmHeader {
	return r.getBlock(h, n, false).Header()
}

func (r *EvmStateReader) GetDagBlock(h hash.Event, n idx.Block) *evmcore.EvmBlock {
	return r.getBlock(h, n, n != 0)
}

func (r *EvmStateReader) getBlock(h hash.Event, n idx.Block, readTxs bool) *evmcore.EvmBlock {
	block := r.store.GetBlock(n)
	if block == nil {
		return nil
	}
	if (h != hash.Event{}) && (h != block.Atropos) {
		return nil
	}

	transactions := r.store.GetBlockTransactions(block)

	evmHeader := evmcore.ToEvmHeader(block)
	evmBlock := &evmcore.EvmBlock{
		EvmHeader: *evmHeader,
	}

	if readTxs {
		evmBlock.Transactions = transactions
	} else {
		evmBlock.Transactions = make(types.Transactions, 0)
	}

	return evmBlock
}

func (s *Store) GetBlockTransactions(block *inter.Block) types.Transactions {
	transactions := make(types.Transactions, 0, len(block.Events)*10)
	txCount := uint(0)
	skipCount := 0
	if block.Index == 0 {
		return transactions
	}
	for _, id := range block.Events {
		e := s.GetEvent(id)
		if e == nil {
			log.Crit("Event not found", "event", id.String())
			continue
		}

		// appends txs except skipped ones
		for _, tx := range e.Transactions {
			if skipCount < len(block.SkippedTxs) && block.SkippedTxs[skipCount] == txCount {
				skipCount++
			} else {
				transactions = append(transactions, tx)
			}
			txCount++
		}
	}
	return transactions
}

func (r *EvmStateReader) StateAt(root common.Hash) (*state.StateDB, error) {
	return r.app.StateDB(root)
}
