package notification

import (
	"gateway/internal/service"
	"github.com/go-redis/redis/v8"
	pb "user-post/proto"
)

func NewNotificationHandler() *notificationHandler {
	return &notificationHandler{
		service.NewUserPostService(),
		redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}

type notificationHandler struct {
	pb.UserPostService
	client *redis.Client
}
