package service

import (
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	micro "go-micro.dev/v4"
	pb "user/proto"
)

func NewUserService() pb.UserService {
	service := micro.NewService(
		micro.Name("userService"),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	
	return pb.NewUserService("userService", service.Client())
}
