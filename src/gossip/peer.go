package gossip

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/inter/idx"
)

var (
	errClosed            = errors.New("peer set is closed")
	errAlreadyRegistered = errors.New("peer is already registered")
	errNotRegistered     = errors.New("peer is not registered")
)

const (
	maxKnownTxs    = 32768 // Maximum transactions hashes to keep in the known list (prevent DOS)
	maxKnownEvents = 1024  // Maximum event hashes to keep in the known list (prevent DOS)

	// maxQueuedTxs is the maximum number of transaction lists to queue up before
	// dropping broadcasts. This is a sensitive number as a transaction list might
	// contain a single transaction, or thousands.
	maxQueuedTxs = 128

	// maxQueuedProps is the maximum number of event propagations to queue up before
	// dropping broadcasts. There's not much point in queueing stale events, so a few
	// that might cover uncles should be enough.
	maxQueuedProps = 4

	// maxQueuedAnns is the maximum number of event announcements to queue up before
	// dropping broadcasts. Similarly to event propagations, there's no point to queue
	// above some healthy uncle limit, so use that.
	maxQueuedAnns = 4

	handshakeTimeout = 5 * time.Second
)

// PeerInfo represents a short summary of the sub-protocol metadata known
// about a connected peer.
type PeerInfo struct {
	Version     int            `json:"version"` // protocol version negotiated
	Epoch       idx.SuperFrame `json:"epoch"`
	NumOfEvents idx.Event      `json:"numOfEvents"`
}

type PeerProgress struct {
	Epoch       idx.SuperFrame
	NumOfEvents idx.Event
}

type peer struct {
	id string

	*p2p.Peer
	rw p2p.MsgReadWriter

	version  int         // Protocol version negotiated
	syncDrop *time.Timer // Timed connection dropper if sync progress isn't validated in time

	lock sync.RWMutex

	knownTxs    mapset.Set                // Set of transaction hashes known to be known by this peer
	knownEvents mapset.Set                // Set of event hashes known to be known by this peer
	queuedTxs   chan []*types.Transaction // Queue of transactions to broadcast to the peer
	queuedProps chan *inter.Event         // Queue of events to broadcast to the peer
	queuedAnns  chan *inter.Event         // Queue of events to announce to the peer
	term        chan struct{}             // Termination channel to stop the broadcaster

	progress PeerProgress
}

func (a *PeerProgress) Less(b PeerProgress) bool {
	if a.Epoch < b.Epoch {
		return true
	}
	return a.Epoch == b.Epoch && a.NumOfEvents < b.NumOfEvents
}

func newPeer(version int, p *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer:        p,
		rw:          rw,
		version:     version,
		id:          fmt.Sprintf("%x", p.ID().Bytes()[:8]),
		knownTxs:    mapset.NewSet(),
		knownEvents: mapset.NewSet(),
		queuedTxs:   make(chan []*types.Transaction, maxQueuedTxs),
		queuedProps: make(chan *inter.Event, maxQueuedProps),
		queuedAnns:  make(chan *inter.Event, maxQueuedAnns),
		term:        make(chan struct{}),
	}
}

// broadcast is a write loop that multiplexes event propagations, announcements
// and transaction broadcasts into the remote peer. The goal is to have an async
// writer that does not lock up node internals.
func (p *peer) broadcast() {
	for {
		select {
		case txs := <-p.queuedTxs:
			if err := p.SendTransactions(txs); err != nil {
				return
			}
			p.Log().Trace("Broadcast transactions", "count", len(txs))

		case event := <-p.queuedProps:
			if err := p.SendEvents([]*inter.Event{event}); err != nil {
				return
			}
			p.Log().Trace("Propagated event", "seq", event.Seq, "hash", event.Hash())

		case event := <-p.queuedAnns:
			if err := p.SendNewEventHashes([]hash.Event{event.Hash()}); err != nil {
				return
			}
			p.Log().Trace("Announced event", "seq", event.Seq, "hash", event.Hash())

		case <-p.term:
			return
		}
	}
}

// close signals the broadcast goroutine to terminate.
func (p *peer) close() {
	close(p.term)
}

// Info gathers and returns a collection of metadata known about a peer.
func (p *peer) Info() *PeerInfo {
	return &PeerInfo{
		Version:     p.version,
		Epoch:       p.progress.Epoch,
		NumOfEvents: p.progress.NumOfEvents,
	}
}

// MarkEvent marks a event as known for the peer, ensuring that the event will
// never be propagated to this particular peer.
func (p *peer) MarkEvent(hash hash.Event) {
	// If we reached the memory allowance, drop a previously known event hash
	for p.knownEvents.Cardinality() >= maxKnownEvents {
		p.knownEvents.Pop()
	}
	p.knownEvents.Add(hash)
}

// MarkTransaction marks a transaction as known for the peer, ensuring that it
// will never be propagated to this particular peer.
func (p *peer) MarkTransaction(hash common.Hash) {
	// If we reached the memory allowance, drop a previously known transaction hash
	for p.knownTxs.Cardinality() >= maxKnownTxs {
		p.knownTxs.Pop()
	}
	p.knownTxs.Add(hash)
}

// SendTransactions sends transactions to the peer and includes the hashes
// in its transaction hash set for future reference.
func (p *peer) SendTransactions(txs types.Transactions) error {
	// Mark all the transactions as known, but ensure we don't overflow our limits
	for _, tx := range txs {
		p.knownTxs.Add(tx.Hash())
	}
	for p.knownTxs.Cardinality() >= maxKnownTxs {
		p.knownTxs.Pop()
	}
	return p2p.Send(p.rw, TxMsg, txs)
}

// AsyncSendTransactions queues list of transactions propagation to a remote
// peer. If the peer's broadcast queue is full, the event is silently dropped.
func (p *peer) AsyncSendTransactions(txs []*types.Transaction) {
	select {
	case p.queuedTxs <- txs:
		// Mark all the transactions as known, but ensure we don't overflow our limits
		for _, tx := range txs {
			p.knownTxs.Add(tx.Hash())
		}
		for p.knownTxs.Cardinality() >= maxKnownTxs {
			p.knownTxs.Pop()
		}
	default:
		p.Log().Debug("Dropping transaction propagation", "count", len(txs))
	}
}

// SendNewEventHashes announces the availability of a number of events through
// a hash notification.
func (p *peer) SendNewEventHashes(hashes []hash.Event) error {
	// Mark all the event hashes as known, but ensure we don't overflow our limits
	for _, hash := range hashes {
		p.knownEvents.Add(hash)
	}
	for p.knownEvents.Cardinality() >= maxKnownEvents {
		p.knownEvents.Pop()
	}
	return p2p.Send(p.rw, NewEventHashesMsg, hashes)
}

// AsyncSendNewEventHash queues the availability of a event for propagation to a
// remote peer. If the peer's broadcast queue is full, the event is silently
// dropped.
func (p *peer) AsyncSendNewEventHash(event *inter.Event) {
	select {
	case p.queuedAnns <- event:
		// Mark all the event hash as known, but ensure we don't overflow our limits
		p.knownEvents.Add(event.Hash())
		for p.knownEvents.Cardinality() >= maxKnownEvents {
			p.knownEvents.Pop()
		}
	default:
		p.Log().Debug("Dropping event announcement", "seq", event.Seq, "hash", event.Hash())
	}
}

// SendNewEvent propagates an entire event to a remote peer.
func (p *peer) SendNewEvents(events []*inter.Event) error {
	// Mark all the event hash as known, but ensure we don't overflow our limits
	for _, event := range events {
		p.knownEvents.Add(event.Hash())
		for p.knownEvents.Cardinality() >= maxKnownEvents {
			p.knownEvents.Pop()
		}
	}
	return p2p.Send(p.rw, NewEventsMsg, events)
}

// AsyncSendNewEvent queues an entire event for propagation to a remote peer. If
// the peer's broadcast queue is full, the event is silently dropped.
func (p *peer) AsyncSendNewEvent(event *inter.Event) {
	select {
	case p.queuedProps <- event:
		// Mark all the event hash as known, but ensure we don't overflow our limits
		p.knownEvents.Add(event.Hash())
		for p.knownEvents.Cardinality() >= maxKnownEvents {
			p.knownEvents.Pop()
		}
	default:
		p.Log().Debug("Dropping event propagation", "seq", event.Seq, "hash", event.Hash())
	}
}

// SendEventHeaders sends a batch of event headers to the remote peer.
/*func (p *peer) SendEventHeaders(headers []*types.Header) error {
	return p2p.Send(p.rw, EventHeadersMsg, headers)
}*/

// SendEventBodies sends a batch of event contents to the remote peer.
func (p *peer) SendEvents(events []*inter.Event) error {
	return p2p.Send(p.rw, EventsMsg, events)
}

// SendEventBodiesRLP sends a batch of event contents to the remote peer from
// an already RLP encoded format.
func (p *peer) SendEventsRLP(events []rlp.RawValue) error {
	return p2p.Send(p.rw, EventsMsg, events)
}

/*// RequestOneHeader is a wrapper around the header query functions to fetch a
// single header. It is used solely by the fetcher.
func (p *peer) RequestOneHeader(hash common.Hash) error {
	p.Log().Debug("Fetching single header", "hash", hash)
	return p2p.Send(p.rw, GetEventHeadersMsg, &getEventHeadersData{Origin: hashOrNumber{Hash: hash}, Amount: uint64(1), Skip: uint64(0), Reverse: false})
}

// RequestHeadersByHash fetches a batch of events' headers corresponding to the
// specified header query, based on the hash of an origin event.
func (p *peer) RequestHeadersByHash(origin common.Hash, amount int, skip int, reverse bool) error {
	p.Log().Debug("Fetching batch of headers", "count", amount, "fromhash", origin, "skip", skip, "reverse", reverse)
	return p2p.Send(p.rw, GetEventHeadersMsg, &getEventHeadersData{Origin: hashOrNumber{Hash: origin}, Amount: uint64(amount), Skip: uint64(skip), Reverse: reverse})
}

// RequestHeadersByNumber fetches a batch of events' headers corresponding to the
// specified header query, based on the number of an origin event.
func (p *peer) RequestHeadersByNumber(origin uint64, amount int, skip int, reverse bool) error {
	p.Log().Debug("Fetching batch of headers", "count", amount, "fromnum", origin, "skip", skip, "reverse", reverse)
	return p2p.Send(p.rw, GetEventHeadersMsg, &getEventHeadersData{Origin: hashOrNumber{Number: origin}, Amount: uint64(amount), Skip: uint64(skip), Reverse: reverse})
}*/

func (p *peer) RequestEvents(ids []hash.Event) error {
	p.Log().Debug("Fetching batch of events", "count", len(ids))
	return p2p.Send(p.rw, GetEventsMsg, ids)
}

// Handshake executes the protocol handshake, negotiating version number,
// network IDs, difficulties, head and genesis object.
func (p *peer) Handshake(network uint64, progress PeerProgress, genesis hash.Hash) error {
	// Send out own handshake in a new thread
	errc := make(chan error, 2)
	var status statusData // safe to read after two values have been received from errc

	go func() {
		errc <- p2p.Send(p.rw, StatusMsg, &statusData{
			ProtocolVersion: uint32(p.version),
			Progress:        progress,
			NetworkId:       network,
			Genesis:         genesis,
		})
	}()
	go func() {
		errc <- p.readStatus(network, &status, genesis)
	}()
	timeout := time.NewTimer(handshakeTimeout)
	defer timeout.Stop()
	for i := 0; i < 2; i++ {
		select {
		case err := <-errc:
			if err != nil {
				return err
			}
		case <-timeout.C:
			return p2p.DiscReadTimeout
		}
	}
	p.progress = progress
	return nil
}

func (p *peer) readStatus(network uint64, status *statusData, genesis hash.Hash) (err error) {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Code != StatusMsg {
		return errResp(ErrNoStatusMsg, "first msg has code %x (!= %x)", msg.Code, StatusMsg)
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	// Decode the handshake and make sure everything matches
	if err := msg.Decode(&status); err != nil {
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}
	if status.Genesis != genesis {
		return errResp(ErrGenesisMismatch, "%x (!= %x)", status.Genesis[:8], genesis[:8])
	}
	if status.NetworkId != network {
		return errResp(ErrNetworkIdMismatch, "%d (!= %d)", status.NetworkId, network)
	}
	if int(status.ProtocolVersion) != p.version {
		return errResp(ErrProtocolVersionMismatch, "%d (!= %d)", status.ProtocolVersion, p.version)
	}
	return nil
}

// String implements fmt.Stringer.
func (p *peer) String() string {
	return fmt.Sprintf("Peer %s [%s]", p.id,
		fmt.Sprintf("fantom/%2d", p.version),
	)
}

// peerSet represents the collection of active peers currently participating in
// the sub-protocol.
type peerSet struct {
	peers  map[string]*peer
	lock   sync.RWMutex
	closed bool
}

// newPeerSet creates a new peer set to track the active participants.
func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[string]*peer),
	}
}

// Register injects a new peer into the working set, or returns an error if the
// peer is already known. If a new peer it registered, its broadcast loop is also
// started.
func (ps *peerSet) Register(p *peer) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.closed {
		return errClosed
	}
	if _, ok := ps.peers[p.id]; ok {
		return errAlreadyRegistered
	}
	ps.peers[p.id] = p
	go p.broadcast()

	return nil
}

// Unregister removes a remote peer from the active set, disabling any further
// actions to/from that particular entity.
func (ps *peerSet) Unregister(id string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	p, ok := ps.peers[id]
	if !ok {
		return errNotRegistered
	}
	delete(ps.peers, id)
	p.close()

	return nil
}

// Peer retrieves the registered peer with the given id.
func (ps *peerSet) Peer(id string) *peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return ps.peers[id]
}

// Len returns if the current number of peers in the set.
func (ps *peerSet) Len() int {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return len(ps.peers)
}

// PeersWithoutEvent retrieves a list of peers that do not have a given event in
// their set of known hashes.
func (ps *peerSet) PeersWithoutEvent(hash hash.Event) []*peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.knownEvents.Contains(hash) {
			list = append(list, p)
		}
	}
	return list
}

// PeersWithoutTx retrieves a list of peers that do not have a given transaction
// in their set of known hashes.
func (ps *peerSet) PeersWithoutTx(hash common.Hash) []*peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.knownTxs.Contains(hash) {
			list = append(list, p)
		}
	}
	return list
}

// BestPeer retrieves the known peer with the currently highest total difficulty.
func (ps *peerSet) BestPeer() *peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	var (
		bestPeer     *peer
		bestProgress PeerProgress
	)
	for _, p := range ps.peers {
		if bestProgress.Less(p.progress) {
			bestPeer, bestProgress = p, p.progress
		}
	}
	return bestPeer
}

// Close disconnects all peers.
// No new peers can be registered after Close has returned.
func (ps *peerSet) Close() {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	for _, p := range ps.peers {
		p.Disconnect(p2p.DiscQuitting)
	}
	ps.closed = true
}