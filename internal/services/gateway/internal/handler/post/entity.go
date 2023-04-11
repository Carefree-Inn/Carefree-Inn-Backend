package post

import (
	"gateway/internal/handler/user"
	"gateway/internal/service"
	pbPost "post/proto"
	pbUser "user/proto"
)

func NewPostHandler() *postHandler {
	return &postHandler{
		service.NewPostService(),
		service.NewUserService(),
	}
}

type postHandler struct {
	pbPost.PostService
	pbUser.UserService
}

type tagInfo struct {
	TagId uint32 `json:"tag_id"`
	Title string `json:"title" binding:"required"`
}

type categoryInfo struct {
	CategoryId uint32 `json:"category_id" binding:"required"`
	Title      string `json:"title"`
}

type PostInfo struct {
	UserInfo   *user.UserInfo `json:"user_info"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Category   categoryInfo   `json:"category"`
	Tags       []*tagInfo     `json:"tags"`
	Star       uint32         `json:"star"`
	CreateTime string         `json:"create_time"`
	PostId     uint32         `json:"post_id"`
}
