package shipwreck

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

type msg any

type Message[T msg] struct {
	SourceID string
	TargetID string
	Msg      T
}

type (
	VoteReply struct {
		Granted bool
	}
	VoteRequest struct {
		Term         int64
		CommitOffset int64
	}
	LogRequest[T nodeMessage] struct {
		CommitOffset int64
		StartOffset  int64
		Entries      []T
	}
	LogReply struct {
		CommitOffset int64
		Success      bool
	}
	ProxyPush[T nodeMessage] struct {
		Value T
	}
	ProxyPushReply struct {
		Ok bool
	}
)

type conn[T nodeMessage] interface {
	ID() string
	RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error)
	AppendEntries(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error)
	ProxyPush(ctx context.Context, log Message[ProxyPush[T]]) (Message[ProxyPushReply], error)
}

type NodeMode int64

const (
	NodeModeFollower NodeMode = iota
	NodeModeCandidate
	NodeModeLeader
)

type peer[T nodeMessage] struct {
	commitOffset int64
	conn         conn[T]
}

type commitCallbackFunc[T nodeMessage] func(ctx context.Context, logs []T) error

type syncStatus struct {
	offset int64
	err    error
}

type node[T nodeMessage] struct {
	// node[T encoding.BinaryMarshaler] data
	id           string // TODO maybe change
	mode         NodeMode
	peers        map[string]peer[T]
	stopped      bool
	stopping     bool
	commitOffset int64
	syncOffset   int64
	storage      Storage[T]

	commitCallback commitCallbackFunc[T]
	syncChLock     sync.Mutex
	syncCh         []chan syncStatus

	// Voting
	electionTimeout *time.Ticker
	votedFor        string
	term            int64
	voteLock        sync.Mutex
	currentLeader   string

	// Leader
	syncTicker *time.Ticker
}

type RaftNode[T nodeMessage] interface {
	ID() string
	String() string
	Mode() NodeMode
	Push(v T) error
	AddPeer(conn conn[T])
	Start(ctx context.Context) error
	Stop(ctx context.Context) error

	conn[T]
}

type nodeMessage any // should be something more concreate

// TODO handle commitCallback
func NewNode[T nodeMessage](storage Storage[T], commitCallback commitCallbackFunc[T]) RaftNode[T] {
	d := time.Duration(rand.Int63n(150)+150) * time.Millisecond
	return &node[T]{
		id:           uuid.New().String(),
		mode:         NodeModeFollower,
		stopped:      true,
		stopping:     true,
		peers:        map[string]peer[T]{},
		commitOffset: 0,
		syncOffset:   0,

		storage:        storage,
		commitCallback: commitCallback,
		syncChLock:     sync.Mutex{},
		syncCh:         []chan syncStatus{},

		electionTimeout: time.NewTicker(d),
		votedFor:        "",
		currentLeader:   "",
		term:            0,

		syncTicker: time.NewTicker(50 * time.Millisecond),
	}
}

// Main function for pushing new values to the log
func (n *node[T]) Push(v T) error {
	if n.stopped {
		return fmt.Errorf("Node is not running")
	}

	if n.stopping {
		return fmt.Errorf("Node is being stopped")
	}

	if n.currentLeader == "" {
		return fmt.Errorf("Cluster is not ready")
	}

	if n.mode == NodeModeFollower {
		leader, ok := n.peers[n.currentLeader]
		if !ok {
			return fmt.Errorf("No leader %v available", n.currentLeader)
		}

		_, err := leader.conn.ProxyPush(context.Background(), Message[ProxyPush[T]]{
			SourceID: n.id,
			TargetID: leader.conn.ID(),
			Msg: ProxyPush[T]{
				Value: v,
			},
		})

		return err
	}

	offset, err := n.storage.Append(v)
	if err != nil {
		return err
	}

	// This will be problematic as channels can block for infinite time
	c := make(chan syncStatus)
	defer func() {
		n.syncChLock.Lock()
		n.syncCh = without(n.syncCh, c)
		n.syncChLock.Unlock()
		close(c)
	}()

	n.syncChLock.Lock()
	n.syncCh = append(n.syncCh, c)
	n.syncChLock.Unlock()

	for status := range c {
		if status.err != nil {
			return status.err
		}
		if status.offset >= offset {
			return nil
		}
	}

	return fmt.Errorf("Failed to commit")
}

func without[T comparable](slice []T, s T) []T {
	result := []T{}
	for _, v := range slice {
		if s != v {
			result = append(result, v)
		}
	}
	return result
}

func (n *node[T]) ID() string {
	return n.id
}

func (n *node[T]) String() string {
	return fmt.Sprintf("Debug %v %v %v %v %v", n.id, n.commitOffset, n.storage.Length(), n.storage.Commited(), len(n.syncCh))
}

func (n *node[T]) Mode() NodeMode {
	return n.mode
}

func (n *node[T]) AddPeer(conn conn[T]) {
	n.peers[conn.ID()] = peer[T]{
		conn:         conn,
		commitOffset: n.commitOffset,
	}
}

func (n *node[T]) Stop(ctx context.Context) error {
	// Signal stopping, this will stop accepting push
	n.stopping = true

	// If leader, we want to wait for next sync tick
	if n.mode == NodeModeLeader {
		<-n.syncTicker.C
	}

	n.stopped = true
	n.electionTimeout.Stop()
	n.syncTicker.Stop()

	return nil
}

const syncTickerDuration = 50 * time.Millisecond

func (n *node[T]) resetTimer() {
	d := time.Duration(rand.Int63n(150)+150) * time.Millisecond
	n.electionTimeout.Reset(d)
	n.syncTicker.Reset(syncTickerDuration)
}

func (n *node[T]) restart() {
	n.mode = NodeModeFollower
	n.stopping = false
	n.stopped = false
	n.resetTimer()
}

func (n *node[T]) Start(ctx context.Context) error {
	n.restart()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-n.electionTimeout.C:
			// No heartbeat recieved, switch to candidate
			if n.mode == NodeModeFollower {
				n.becomeCandidate()
			}
		case <-n.syncTicker.C:
			if n.mode == NodeModeLeader {
				n.syncPeers()
			}
			if n.mode == NodeModeCandidate {
				n.startNewTerm()
			}
		}
	}
}

func (n *node[T]) becomeLeader() {
	if n.mode == NodeModeLeader {
		return
	}
	slog.Debug("node become leader", "ID", n.id)
	n.mode = NodeModeLeader
	n.currentLeader = n.id

	n.syncPeers()
}

func (n *node[T]) becomeFollower() {
	if n.mode == NodeModeFollower {
		return
	}
	slog.Debug("node become follower", "ID", n.id)
	n.mode = NodeModeFollower
}

func (n *node[T]) becomeCandidate() {
	if n.mode == NodeModeCandidate {
		return
	}
	slog.Debug("node become candidate", "ID", n.id)
	n.mode = NodeModeCandidate
	n.resetTimer()
}

func (n *node[T]) commitLogs(offset int64) error {
	commited, err := n.storage.Commit(offset)
	if err != nil {
		return err
	}
	if len(commited) <= 0 {
		return nil
	}

	n.commitOffset = offset

	ctx, cancel := context.WithTimeout(context.Background(), syncTickerDuration/2) // really tight timing
	defer cancel()

	err = n.commitCallback(ctx, commited) //  Can block the goroutine infinitly, should be guarded from such behaviour, or panic when that happen
	if err != nil {
		return err
	}

	if n.mode == NodeModeLeader {
		n.syncChLock.Lock()
		for _, c := range n.syncCh {
			c <- syncStatus{
				offset: offset,
				err:    nil,
			}
		}
		n.syncChLock.Unlock()
	}

	return nil
}

func (n *node[T]) syncPeers() {
	if n.stopped {
		return
	}
	n.syncOffset = n.storage.Length()

	ctx := context.Background()
	wg := sync.WaitGroup{}
	results := make(chan Message[LogReply], len(n.peers))
	for _, peer := range n.peers {
		peer := peer
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := peer.conn.AppendEntries(ctx, Message[LogRequest[T]]{
				SourceID: n.id,
				TargetID: peer.conn.ID(),
				Msg: LogRequest[T]{
					CommitOffset: n.commitOffset,
					StartOffset:  peer.commitOffset,
					Entries:      n.getUncommitedLogs(peer),
				},
			})
			// TODO better error handling, add backoff
			if err != nil {
				slog.ErrorContext(ctx, "Ping failed", "Err", err)
			}
			// TOOD handle catchup
			results <- resp
		}()
	}
	wg.Wait()
	close(results)

	// leader already written the value
	successes := int64(1)
	for result := range results {
		if result.Msg.Success {
			successes += 1
		} else {
			p := n.peers[result.SourceID]
			p.commitOffset = result.Msg.CommitOffset
			n.peers[result.SourceID] = p
		}
	}

	canCommit := successes >= int64(1+len(n.peers)/2)
	if canCommit {
		err := n.commitLogs(n.syncOffset)
		if err != nil {
			// TODO handle
			// return nil, err
		}

		for id, p := range n.peers {
			p.commitOffset = n.syncOffset
			n.peers[id] = p
		}
	} else {
		err := n.storage.Discard(n.commitOffset, n.syncOffset)
		if err != nil {
			// TODO handle
			// return nil, err
		}

		n.syncChLock.Lock()
		for _, c := range n.syncCh {
			c <- syncStatus{
				offset: 0,
				err:    fmt.Errorf("Operation failed to sync"),
			}
		}
		n.syncChLock.Unlock()
	}
}

func (n *node[T]) getUncommitedLogs(p peer[T]) []T {
	values, err := n.storage.Get(min(n.storage.Length(), p.commitOffset), n.storage.Length())
	if err != nil {
		// TODO handle
		// return nil, err
	}
	return values
}

func (n *node[T]) startNewTerm() {
	if n.stopped {
		return
	}

	n.term += 1
	n.votedFor = n.id
	ctx := context.Background()

	granted := atomic.Int64{}
	granted.Add(1) // node voted for itself so we can just add one here

	wg := sync.WaitGroup{}
	for _, peer := range n.peers {
		peer := peer
		wg.Add(1)
		go func() {
			defer wg.Done()
			reply, err := peer.conn.RequestVote(ctx, Message[VoteRequest]{
				SourceID: n.id,
				TargetID: peer.conn.ID(),
				Msg: VoteRequest{
					CommitOffset: n.commitOffset,
					Term:         n.term,
				},
			})
			if err != nil {
				slog.ErrorContext(ctx, "Vote requested failed", "Err", err)
			}
			if reply.Msg.Granted {
				granted.Add(1)
			}
		}()
	}
	wg.Wait()

	hasMajorityVote := granted.Load() >= int64(1+len(n.peers)/2)
	if hasMajorityVote {
		n.becomeLeader()
	}
}

// proxyPush implements conn.
func (n *node[T]) ProxyPush(ctx context.Context, value Message[ProxyPush[T]]) (Message[ProxyPushReply], error) {
	err := n.Push(value.Msg.Value)
	if err != nil {
		return Message[ProxyPushReply]{
			SourceID: n.id,
			TargetID: value.SourceID,
			Msg: ProxyPushReply{
				Ok: false,
			},
		}, err
	}

	return Message[ProxyPushReply]{
		SourceID: n.id,
		TargetID: value.SourceID,
		Msg: ProxyPushReply{
			Ok: true,
		},
	}, nil
}

// requestVote implements conn.
func (n *node[T]) RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error) {
	n.voteLock.Lock()
	defer n.voteLock.Unlock()

	if n.stopped {
		return Message[VoteReply]{
			SourceID: n.id,
			TargetID: vote.SourceID,
		}, fmt.Errorf("node unreachable")
	}

	n.resetTimer()
	if vote.Msg.Term < n.term {
		return Message[VoteReply]{
			SourceID: n.id,
			TargetID: vote.SourceID,
			Msg: VoteReply{
				Granted: false,
			},
		}, nil
	}

	if vote.Msg.CommitOffset < n.commitOffset {
		return Message[VoteReply]{
			SourceID: n.id,
			TargetID: vote.SourceID,
			Msg: VoteReply{
				Granted: false,
			},
		}, nil
	}

	if vote.Msg.Term == n.term {
		return Message[VoteReply]{
			SourceID: n.id,
			TargetID: vote.SourceID,
			Msg: VoteReply{
				Granted: vote.SourceID == n.votedFor,
			},
		}, nil
	}

	if n.mode == NodeModeFollower {
		n.becomeFollower()
	}
	n.term = vote.Msg.Term
	n.votedFor = vote.SourceID
	return Message[VoteReply]{
		SourceID: n.id,
		TargetID: vote.SourceID,
		Msg: VoteReply{
			Granted: true,
		},
	}, nil
}

// ping implements conn.
func (n *node[T]) AppendEntries(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error) {
	if n.stopped {
		return Message[LogReply]{
			SourceID: n.id,
			TargetID: log.SourceID,
		}, fmt.Errorf("node unreachable")
	}

	// Cleanup
	if n.mode == NodeModeCandidate {
		n.becomeFollower()
	}
	n.votedFor = ""
	n.currentLeader = log.SourceID
	n.resetTimer()

	// Messages were not committed
	if n.syncOffset > log.Msg.StartOffset {
		err := n.storage.Discard(log.Msg.StartOffset, n.storage.Length())
		if err != nil {
			// TODO handle
			// return nil, err
		}
		n.syncOffset = log.Msg.StartOffset
	}

	if n.syncOffset != log.Msg.StartOffset {
		return Message[LogReply]{
			SourceID: n.id,
			TargetID: log.SourceID,
			Msg: LogReply{
				CommitOffset: n.commitOffset,
				Success:      false,
			},
		}, nil
	}

	// Write new logs
	_, err := n.storage.Append(log.Msg.Entries...)
	if err != nil {
		// TODO handle
		// return nil, err
	}
	n.syncOffset = log.Msg.StartOffset + int64(len(log.Msg.Entries))

	// Commit
	err = n.commitLogs(log.Msg.CommitOffset)
	if err != nil {
		// TODO handle
		// return nil, err
	}

	return Message[LogReply]{
		SourceID: n.id,
		TargetID: log.SourceID,
		Msg: LogReply{
			CommitOffset: n.commitOffset,
			Success:      true,
		},
	}, nil
}

var _ conn[nodeMessage] = (*node[nodeMessage])(nil)
