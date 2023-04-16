package service

import (
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	pb "github.com/jackj-ohn1/package/proto/user"
	"github.com/opentracing/opentracing-go"
	micro "go-micro.dev/v4"
)

func NewUserService() pb.UserService {
	service := micro.NewService(
		micro.Name("userService"),
		micro.WrapClient(opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	
	return pb.NewUserService("userService", service.Client())
}
