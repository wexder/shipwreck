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

type GrpcNodeServer[T nodeMessage] struct {
	n    RaftNode[T]
	addr string

	proto.UnsafeNodeServer
}

// ProxyPush implements proto.NodeServer.
func (gns *GrpcNodeServer[T]) ProxyPush(ctx context.Context, req *proto.ProxyPushRequest) (*proto.ProxyPushReply, error) {
	var value T
	err := gob.NewDecoder(bytes.NewBuffer(req.Push)).Decode(&value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Entries decoding failed: %v", err.Error())
	}
	vote, err := gns.n.ProxyPush(ctx, Message[ProxyPush[T]]{
		SourceID: req.SourceId,
		TargetID: req.TargeId,
		Msg: ProxyPush[T]{
			Value: value,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Proxy push request failed: %v", err.Error())
	}

	return &proto.ProxyPushReply{
		SourceId: vote.SourceID,
		TargeId:  vote.TargetID,
		Ok:       true,
	}, nil
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

// AppendEntries implements proto.NodeServer.
func (gns *GrpcNodeServer[T]) AppendEntries(ctx context.Context, req *proto.LogRequest) (*proto.LogReply, error) {
	entries := []T{}
	err := gob.NewDecoder(bytes.NewBuffer(req.Entries)).Decode(&entries)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Entries decoding failed: %v", err.Error())
	}

	sync, err := gns.n.AppendEntries(ctx, Message[LogRequest[T]]{
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

func NewGrpcNodeServer[T nodeMessage](n RaftNode[T], addr string) *GrpcNodeServer[T] {
	return &GrpcNodeServer[T]{
		n:    n,
		addr: addr,
	}
}

func (gns *GrpcNodeServer[T]) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", gns.addr)
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

type grpcConn[T nodeMessage] struct {
	client proto.NodeClient
	id     string
}

func NewGrpcConn[T nodeMessage](addr string) (conn[T], error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewNodeClient(conn)
	return &grpcConn[T]{
		client: client,
		id:     addr,
	}, nil
}

var _ conn[nodeMessage] = (*grpcConn[nodeMessage])(nil)

// ID implements conn.
func (gc *grpcConn[T]) ID() string {
	return gc.id
}

// ProxyPush implements conn.
func (gc *grpcConn[T]) ProxyPush(ctx context.Context, value Message[ProxyPush[T]]) (Message[ProxyPushReply], error) {
	valueBytes := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(valueBytes).Encode(value.Msg.Value)
	if err != nil {
		return Message[ProxyPushReply]{
			SourceID: gc.id,
			TargetID: value.SourceID,
			Msg: ProxyPushReply{
				Ok: false,
			},
		}, err
	}
	reply, err := gc.client.ProxyPush(ctx, &proto.ProxyPushRequest{
		Push: valueBytes.Bytes(),
	})
	if err != nil {
		return Message[ProxyPushReply]{
			SourceID: gc.id,
			TargetID: value.SourceID,
			Msg: ProxyPushReply{
				Ok: false,
			},
		}, err
	}

	return Message[ProxyPushReply]{
		SourceID: gc.id,
		TargetID: value.SourceID,
		Msg: ProxyPushReply{
			Ok: reply.Ok,
		},
	}, nil
}

// RequestVote implements conn.
func (gc *grpcConn[T]) RequestVote(ctx context.Context, vote Message[VoteRequest]) (Message[VoteReply], error) {
	reply, err := gc.client.RequestVote(ctx, &proto.VoteRequest{
		SourceId:     gc.id,
		TargeId:      vote.SourceID,
		Term:         vote.Msg.Term,
		CommitOffset: vote.Msg.CommitOffset,
	})
	if err != nil {
		return Message[VoteReply]{
			SourceID: gc.id,
			TargetID: vote.SourceID,
			Msg: VoteReply{
				Granted: false,
			},
		}, err
	}

	return Message[VoteReply]{
		SourceID: gc.id,
		TargetID: vote.SourceID,
		Msg: VoteReply{
			Granted: reply.Granted,
		},
	}, nil
}

// AppendEntries implements conn.
func (gc *grpcConn[T]) AppendEntries(ctx context.Context, log Message[LogRequest[T]]) (Message[LogReply], error) {
	entries := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(entries).Encode(log.Msg.Entries)
	if err != nil {
		return Message[LogReply]{
			SourceID: gc.id,
			TargetID: log.SourceID,
			Msg: LogReply{
				CommitOffset: 0,
				Success:      false,
			},
		}, err
	}

	reply, err := gc.client.AppendEntries(ctx, &proto.LogRequest{
		SourceId:     log.SourceID,
		TargeId:      log.TargetID,
		CommitOffset: log.Msg.CommitOffset,
		StartOffset:  log.Msg.StartOffset,
		Entries:      entries.Bytes(),
	})
	if err != nil {
		return Message[LogReply]{
			SourceID: gc.id,
			TargetID: log.SourceID,
			Msg: LogReply{
				CommitOffset: 0,
				Success:      false,
			},
		}, err
	}

	return Message[LogReply]{
		SourceID: gc.id,
		TargetID: reply.SourceId,
		Msg: LogReply{
			CommitOffset: reply.CommitOffset,
			Success:      reply.Success,
		},
	}, nil
}
