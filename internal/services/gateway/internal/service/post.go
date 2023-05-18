package service

import (
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	pb "post/proto"
)

func NewPostService() pb.PostService {
	service := micro.NewService(
		micro.Name("postService"),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	
	return pb.NewPostService("postService", service.Client())
}
