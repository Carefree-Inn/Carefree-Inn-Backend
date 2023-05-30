package like

import (
	"gateway/internal/service"
	pb "user-post/proto"
	pbUser "user/proto"
)

type likeHandler struct {
	pb.UserPostService
	pbUser.UserService
}

func NewLikeHandler() *likeHandler {
	return &likeHandler{
		service.NewUserPostService(),
		service.NewUserService(),
	}
}

func (l *likeHandler) Assemble() {

}
