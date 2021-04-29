package gossip

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"

	"github.com/Fantom-foundation/go-lachesis/lachesis"
	"github.com/Fantom-foundation/go-lachesis/logger"
)

func init() {
	// log.Root().SetHandler(log.LvlFilterHandler(log.LvlTrace, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}

var testAccount, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

// Tests that handshake failures are detected and reported correctly.
func TestStatusMsgErrors62(t *testing.T) {
	logger.SetTestMode(t)
	testStatusMsgErrors(t, lachesis62)
}

func testStatusMsgErrors(t *testing.T, protocol int) {
	pm, _ := newTestProtocolManagerMust(t, 5, 5, nil, nil)
	var (
		genesis   = pm.engine.GetGenesisHash()
		networkID = lachesis.FakeNetworkID
	)
	pm.downloader.Terminate() // disable downloader so test would be deterministic
	defer pm.Stop()

	tests := []struct {
		code      uint64
		data      interface{}
		wantError error
	}{
		{
			code: EvmTxMsg, data: []interface{}{},
			wantError: errResp(ErrNoStatusMsg, "first msg has code 2 (!= 0)"),
		},
		{
			code: EthStatusMsg, data: ethStatusData{ProtocolVersion: 10, NetworkID: networkID, Genesis: genesis},
			wantError: errResp(ErrProtocolVersionMismatch, "10 (!= %d)", protocol),
		},
		{
			code: EthStatusMsg, data: ethStatusData{ProtocolVersion: uint32(protocol), NetworkID: 999, Genesis: genesis},
			wantError: errResp(ErrNetworkIDMismatch, "999 (!= %d)", networkID),
		},
		{
			code: EthStatusMsg, data: ethStatusData{ProtocolVersion: uint32(protocol), NetworkID: networkID, Genesis: common.Hash{3}},
			wantError: errResp(ErrGenesisMismatch, "0300000000000000 (!= %x)", genesis.Bytes()[:8]),
		},
	}

	for i, test := range tests {
		p, errc := newTestPeer("peer", protocol, pm, false)
		// The send call might hang until reset because
		// the protocol might not read the payload.
		go p2p.Send(p.app, test.code, test.data)

		select {
		case err := <-errc:
			if err == nil {
				t.Fatalf("test %d: protocol returned nil error, want %q", i, test.wantError)
			} else if err.Error() != test.wantError.Error() {
				t.Fatalf("test %d: wrong error: got %q, want %q", i, err, test.wantError)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("protocol did not shut down within 2 seconds")
		}
		p.close()
	}
}

// This test checks that received transactions are added to the local pool.
func TestRecvTransactions62(t *testing.T) {
	logger.SetTestMode(t)
	testRecvTransactions(t, lachesis62)
}

func testRecvTransactions(t *testing.T, protocol int) {
	txAdded := make(chan []*types.Transaction)
	pm, _ := newTestProtocolManagerMust(t, 5, 5, txAdded, nil)
	pm.synced = 1             // mark synced to accept transactions
	pm.downloader.Terminate() // disable downloader so test would be deterministic
	p, _ := newTestPeer("peer", protocol, pm, true)
	defer pm.Stop()
	defer p.close()

	tx := newTestTransaction(testAccount, 0, 0)
	if err := p2p.Send(p.app, EvmTxMsg, []interface{}{tx}); err != nil {
		t.Fatalf("send error: %v", err)
	}
	select {
	case added := <-txAdded:
		if len(added) != 1 {
			t.Fatalf("wrong number of added transactions: got %d, want 1", len(added))
		} else if added[0].Hash() != tx.Hash() {
			t.Fatalf("added wrong tx hash: got %v, want %v", added[0].Hash(), tx.Hash())
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("no NewTxsNotify received within 2 seconds")
	}
}

// This test checks that pending transactions are sent.
func TestSendTransactions62(t *testing.T) {
	logger.SetTestMode(t)
	testSendTransactions(t, lachesis62)
}

func testSendTransactions(t *testing.T, protocol int) {
	pm, _ := newTestProtocolManagerMust(t, 5, 5, nil, nil)
	defer pm.Stop()
	pm.downloader.Terminate() // disable downloader so test would be deterministic

	// Fill the pool with big transactions.
	const txsize = txsyncPackSize / 10
	alltxs := make([]*types.Transaction, 100)
	for nonce := range alltxs {
		alltxs[nonce] = newTestTransaction(testAccount, uint64(nonce), txsize)
	}
	pm.txpool.AddRemotes(alltxs)

	// Connect several peers. They should all receive the pending transactions.
	var wg sync.WaitGroup
	checktxs := func(p *testPeer) {
		defer wg.Done()
		defer p.close()
		seen := make(map[common.Hash]bool)
		for _, tx := range alltxs {
			seen[tx.Hash()] = false
		}
		for n := 0; n < len(alltxs) && !t.Failed(); {
			var txs []*types.Transaction
			msg, err := p.app.ReadMsg()
			if err != nil {
				t.Fatalf("%v: read error: %v", p.Peer, err)
			} else if msg.Code != EvmTxMsg {
				t.Fatalf("%v: got code %d, want TxMsg", p.Peer, msg.Code)
			}
			if err := msg.Decode(&txs); err != nil {
				t.Fatalf("%v: %v", p.Peer, err)
			}
			for _, tx := range txs {
				hash := tx.Hash()
				seentx, want := seen[hash]
				if seentx {
					t.Fatalf("%v: got tx more than once: %x", p.Peer, hash)
				}
				if !want {
					t.Fatalf("%v: got unexpected tx: %x", p.Peer, hash)
				}
				seen[hash] = true
				n++
			}
		}
	}
	for i := 0; i < 3; i++ {
		p, _ := newTestPeer(fmt.Sprintf("peer #%d", i), protocol, pm, true)
		wg.Add(1)
		go checktxs(p)
	}
	wg.Wait()
}
