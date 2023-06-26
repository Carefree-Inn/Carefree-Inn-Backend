package route

import (
	"gateway/internal/handler/post"
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func postRoute(engine *gin.RouterGroup) {
	pRoute := engine.Group("/post")
	pHandler := post.NewPostHandler()
	
	pRoute.POST("/", middlewares.Auth(), pHandler.CreatePost)
	pRoute.POST("/search", middlewares.Auth(), pHandler.SearchPost)
	
	pRoute.DELETE("/", middlewares.Auth(), pHandler.DeletePost)
	
	pRoute.GET("/category/all", pHandler.GetCategory)
	pRoute.GET("/category", middlewares.Auth(), pHandler.GetPostOfCategory)
	pRoute.GET("/tag", middlewares.Auth(), pHandler.GetPostOfTag)
	pRoute.GET("/user", middlewares.Auth(), pHandler.GetPostOfUser)
	pRoute.GET("/liked", middlewares.Auth(), pHandler.GetPostOfUserLiked)
	pRoute.GET("/info", middlewares.Auth(), pHandler.GetPost)
	pRoute.GET("/square", middlewares.Auth(), pHandler.PostSquare)
	
}
