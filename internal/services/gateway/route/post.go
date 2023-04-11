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
	pRoute.DELETE("/", middlewares.Auth(), pHandler.DeletePost)
	pRoute.GET("/category/all", pHandler.GetCategory)
	pRoute.GET("/category", pHandler.GetPostOfCategory)
	pRoute.GET("/tag", pHandler.GetPostOfTag)
}
