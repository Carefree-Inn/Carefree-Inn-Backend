package handler

import (
	"context"
	"io"
	"time"

	"go-micro.dev/v4/logger"

	pb "user-post/proto"
)

type UserPost struct{}

func (e *UserPost) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received UserPost.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *UserPost) ClientStream(ctx context.Context, stream pb.UserPost_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *UserPost) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.UserPost_ServerStreamStream) error {
	logger.Infof("Received UserPost.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *UserPost) BidiStream(ctx context.Context, stream pb.UserPost_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}