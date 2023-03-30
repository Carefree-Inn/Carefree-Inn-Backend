package user

import (
	"gateway/internal/service"
	pb "user/proto"
)

func NewUserHandler() *userHandler {
	return &userHandler{
		service.NewUserService(),
	}
}

type userHandler struct {
	pb.UserService
}
