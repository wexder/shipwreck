package shipwreck

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
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
	LogRequest[T any] struct {
		CommitOffset int64
		StartOffset  int64
		Entries      []T
	}
	LogReply struct {
		CommitOffset int64
		Success      bool
	}
)

type conn[T any] interface {
	// This might not be neccessary
	ID() string
	RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error)
	SyncLog(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error)
}

type NodeMode int64

const (
	NodeModeFollower NodeMode = iota
	NodeModeCandidate
	NodeModeLeader
)

type peer[T any] struct {
	commitOffset int64
	conn         conn[T]
}

type commitCallbackFunc[T any] func(logs []T) error

type Node[T any] struct {
	// Node[T any] data
	id           string // TODO maybe change
	mode         NodeMode
	peers        map[string]peer[T]
	stopped      bool
	commitOffset int64
	syncOffset   int64
	storage      Storage[T]

	commitCallback commitCallbackFunc[T]

	// Voting
	electionTimeout *time.Ticker
	votedFor        string
	term            int64
	voteLock        sync.Mutex
	currentLeader   string

	// Leader
	syncTicker *time.Ticker
}

// TODO handle commitCallback
func NewNode[T any](id string, storage Storage[T], commitCallback commitCallbackFunc[T]) *Node[T] {
	d := time.Duration(rand.Int63n(150)+150) * time.Millisecond
	return &Node[T]{
		id:           id,
		mode:         NodeModeFollower,
		stopped:      true,
		peers:        map[string]peer[T]{},
		commitOffset: 0,
		syncOffset:   0,

		storage:        storage,
		commitCallback: commitCallback,

		electionTimeout: time.NewTicker(d),
		votedFor:        "",
		currentLeader:   "",
		term:            0,

		syncTicker: time.NewTicker(50 * time.Millisecond),
	}
}

// Main function for pushing new values to the log
func (n *Node[T]) Push(v T) error {
	if n.mode == NodeModeFollower {
		// TODO msg proxy
		// leader, ok := n.peers[n.currentLeader]
		// if !ok {
		// 	return fmt.Errorf("Leader not found")
		// }

		// leader.conn.Push()
		return nil
	}

	err := n.storage.Append(v)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node[T]) ID() string {
	return n.id
}

func (n *Node[T]) Mode() NodeMode {
	return n.mode
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v (%v, %v) %+v", n.mode, n.commitOffset, n.syncOffset, n.storage.Commited())
}

func (n *Node[T]) AddPeer(conn conn[T]) {
	n.peers[conn.ID()] = peer[T]{
		conn:         conn,
		commitOffset: n.commitOffset,
	}
}

func (n *Node[T]) Stop() {
	n.stopped = true
	n.electionTimeout.Stop()
	n.syncTicker.Stop()
}

func (n *Node[T]) resetTimer() {
	d := time.Duration(rand.Int63n(150)+150) * time.Millisecond
	n.electionTimeout.Reset(d)
	n.syncTicker.Reset(50 * time.Millisecond)
}

func (n *Node[T]) Restart() {
	n.mode = NodeModeFollower
	n.stopped = false
	n.resetTimer()
}

func (n *Node[T]) Start(ctx context.Context) error {
	n.stopped = false
	n.resetTimer()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-n.electionTimeout.C:
			// resets
			if n.mode == NodeModeFollower {
				n.becomeCandidate()
			}
		case <-n.syncTicker.C:
			if n.mode == NodeModeLeader {
				n.syncPeers()
			}
			if n.mode == NodeModeCandidate {
				n.candidateRequestPeerVotes()
			}
		}
	}
}

func (n *Node[T]) becomeLeader() {
	if n.mode == NodeModeLeader {
		return
	}
	slog.Debug("Node become leader", "ID", n.id)
	n.mode = NodeModeLeader

	n.syncPeers()
}

func (n *Node[T]) becomeFollower() {
	if n.mode == NodeModeFollower {
		return
	}
	slog.Debug("Node become follower", "ID", n.id)
	n.mode = NodeModeFollower
}

func (n *Node[T]) becomeCandidate() {
	if n.mode == NodeModeCandidate {
		return
	}
	slog.Debug("Node become candidate", "ID", n.id)
	n.mode = NodeModeCandidate
	n.resetTimer()
}

func (n *Node[T]) commitLogs(offset int64) error {
	commited, err := n.storage.Commit(offset)
	if err != nil {
		return err
	}
	if len(commited) <= 0 {
		return nil
	}

	err = n.commitCallback(commited)
	if err != nil {
		return err
	}

	n.commitOffset = offset
	return nil
}

func (n *Node[T]) syncPeers() {
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
			resp, err := peer.conn.SyncLog(ctx, Message[LogRequest[T]]{
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

	canCommit := successes > int64(1+len(n.peers)/2)
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
		// TODO handle callback
	} else {
		err := n.storage.Discard(n.commitOffset, n.syncOffset)
		if err != nil {
			// TODO handle
			// return nil, err
		}
	}
}

func (n *Node[T]) getUncommitedLogs(p peer[T]) []T {
	values, err := n.storage.Get(min(n.storage.Length(), p.commitOffset), n.storage.Length())
	if err != nil {
		// TODO handle
		// return nil, err
	}
	return values
}

func (n *Node[T]) candidateRequestPeerVotes() {
	if n.stopped {
		return
	}

	n.term += 1
	n.votedFor = n.id
	ctx := context.Background()

	granted := atomic.Int64{}
	granted.Add(1) // Node voted for itself so we can just add one here

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

	hasMajorityVote := granted.Load() > int64(1+len(n.peers)/2)
	if hasMajorityVote {
		n.becomeLeader()
	}
}

// requestVote implements conn.
func (n *Node[T]) RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error) {
	n.voteLock.Lock()
	defer n.voteLock.Unlock()

	if n.stopped {
		return Message[VoteReply]{
			SourceID: n.id,
			TargetID: vote.SourceID,
		}, fmt.Errorf("Node unreachable")
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
func (n *Node[T]) SyncLog(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error) {
	if n.stopped {
		return Message[LogReply]{
			SourceID: n.id,
			TargetID: log.SourceID,
		}, fmt.Errorf("Node unreachable")
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
	err := n.storage.Append(log.Msg.Entries...)
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

var _ conn[any] = (*Node[any])(nil)
