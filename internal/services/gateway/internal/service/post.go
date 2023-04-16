package service

import (
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	pb "github.com/jackj-ohn1/package/proto/post"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
)

func NewPostService() pb.PostService {
	service := micro.NewService(
		micro.Name("postService"),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	
	return pb.NewPostService("postService", service.Client())
}
