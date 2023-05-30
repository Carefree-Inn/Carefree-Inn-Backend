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
	pRoute.GET("/category", middlewares.SetAccount(), pHandler.GetPostOfCategory)
	pRoute.GET("/tag", middlewares.SetAccount(), pHandler.GetPostOfTag)
	pRoute.POST("/search", middlewares.SetAccount(), pHandler.SearchPost)
	pRoute.GET("/user", middlewares.Auth(), pHandler.GetPostOfUser)
	pRoute.GET("/liked", middlewares.Auth(), pHandler.GetPostOfUserLiked)
}
