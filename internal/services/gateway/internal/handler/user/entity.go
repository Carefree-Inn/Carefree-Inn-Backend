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

type userInfo struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Sex      int8   `json:"sex"`
	Avatar   string `json:"avatar"`
	Days     int32  `json:"days"`
}
