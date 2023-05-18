package service

import (
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	pb "user-post/proto"
)

func NewUserPostService() pb.UserPostService {
	service := micro.NewService(
		micro.Name("userPostService"),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	
	return pb.NewUserPostService("userPostService", service.Client())
}
