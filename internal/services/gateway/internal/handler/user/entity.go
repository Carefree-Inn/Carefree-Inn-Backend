package user

import (
	"gateway/internal/service"
	pb "github.com/jackj-ohn1/package/proto/user"
)

func NewUserHandler() *userHandler {
	return &userHandler{
		service.NewUserService(),
	}
}

type userHandler struct {
	pb.UserService
}
