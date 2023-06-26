package route

import (
	"gateway/internal/handler/comment"
	"gateway/internal/handler/like"
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func userPostRoute(engine *gin.RouterGroup) {
	likeRoute := engine.Group("/like")
	likeHandler := like.NewLikeHandler()
	
	likeRoute.POST("/", middlewares.Auth(), likeHandler.MakeLike)
	likeRoute.DELETE("/", middlewares.Auth(), likeHandler.CancelLike)
	
	commentRoute := engine.Group("/comment")
	commentHandler := comment.NewCommentHandler()
	
	commentRoute.POST("/", middlewares.Auth(), commentHandler.MakeComment)
	commentRoute.DELETE("/", middlewares.Auth(), commentHandler.DeleteComment)
	commentRoute.GET("/post", middlewares.Auth(), commentHandler.GetCommentOfPost)
	commentRoute.GET("/user", middlewares.Auth(), commentHandler.GetCommentOfUser)
	
}
