package shipwreck

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net"

	proto "github.com/wexder/shipwreck/protos/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type GrpcNodeServer[T any] struct {
	n    *Node[T]
	port int64

	proto.UnsafeNodeServer
}

// RequestVote implements proto.NodeServer.
func (gns *GrpcNodeServer[T]) RequestVote(ctx context.Context, req *proto.VoteRequest) (*proto.VoteReply, error) {
	vote, err := gns.n.RequestVote(ctx, Message[VoteRequest]{
		SourceID: req.SourceId,
		TargetID: req.TargeId,
		Msg: VoteRequest{
			Term:         req.Term,
			CommitOffset: req.CommitOffset,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Vote request failed: %v", err.Error())
	}

	return &proto.VoteReply{
		SourceId: vote.SourceID,
		TargeId:  vote.TargetID,
		Granted:  vote.Msg.Granted,
	}, nil
}

// SyncLog implements proto.NodeServer.
func (gns *GrpcNodeServer[T]) SyncLog(ctx context.Context, req *proto.LogRequest) (*proto.LogReply, error) {
	entries := []T{}
	err := gob.NewDecoder(bytes.NewBuffer(req.Entries)).Decode(&entries)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Entries decoding failed: %v", err.Error())
	}

	sync, err := gns.n.SyncLog(ctx, Message[LogRequest[T]]{
		SourceID: req.SourceId,
		TargetID: req.TargeId,
		Msg: LogRequest[T]{
			CommitOffset: req.CommitOffset,
			StartOffset:  req.StartOffset,
			Entries:      entries,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Log sync failed: %v", err.Error())
	}

	return &proto.LogReply{
		SourceId:     sync.SourceID,
		TargeId:      sync.TargetID,
		CommitOffset: sync.Msg.CommitOffset,
		Success:      sync.Msg.Success,
	}, nil
}

func NewGrpcNodeServer[T any](n *Node[T], port int64) *GrpcNodeServer[T] {
	return &GrpcNodeServer[T]{
		n:    n,
		port: port,
	}
}

func (gns *GrpcNodeServer[T]) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", gns.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterNodeServer(grpcServer, gns)

	// TODO FIX
	// TODO add graceful shutdown
	go func() {
		err := gns.n.Start(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	// TODO add graceful shutdown
	return grpcServer.Serve(lis)
}

type GrpcConn[T any] struct {
	client proto.NodeClient
	id     string
}

func NewGrpcConn[T any](addr string, nodeID string) (*GrpcConn[T], error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewNodeClient(conn)
	return &GrpcConn[T]{
		client: client,
		id:     nodeID,
	}, nil
}

var _ conn[any] = (*GrpcConn[any])(nil)

// ID implements conn.
func (gc *GrpcConn[T]) ID() string {
	return gc.id
}

// RequestVote implements conn.
func (gc *GrpcConn[T]) RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error) {
	reply, err := gc.client.RequestVote(ctx, &proto.VoteRequest{
		SourceId:     vote.SourceID,
		TargeId:      vote.TargetID,
		Term:         vote.Msg.Term,
		CommitOffset: vote.Msg.CommitOffset,
	})
	if err != nil {
		return Message[VoteReply]{}, err
	}

	return Message[VoteReply]{
		SourceID: reply.SourceId,
		TargetID: reply.TargeId,
		Msg: VoteReply{
			Granted: reply.Granted,
		},
	}, nil
}

// SyncLog implements conn.
func (gc *GrpcConn[T]) SyncLog(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error) {
	entries := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(entries).Encode(log.Msg.Entries)
	if err != nil {
		return Message[LogReply]{}, err
	}

	reply, err := gc.client.SyncLog(ctx, &proto.LogRequest{
		SourceId:     log.SourceID,
		TargeId:      log.TargetID,
		CommitOffset: log.Msg.CommitOffset,
		StartOffset:  log.Msg.StartOffset,
		Entries:      entries.Bytes(),
	})
	if err != nil {
		return Message[LogReply]{}, err
	}

	return Message[LogReply]{
		SourceID: reply.SourceId,
		TargetID: reply.TargeId,
		Msg: LogReply{
			CommitOffset: reply.CommitOffset,
			Success:      reply.Success,
		},
	}, nil
}
