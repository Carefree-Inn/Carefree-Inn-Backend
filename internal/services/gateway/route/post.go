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
	pRoute.POST("/search", middlewares.SetAccount(), pHandler.SearchPost)
	
	pRoute.DELETE("/", middlewares.Auth(), pHandler.DeletePost)
	
	pRoute.GET("/category/all", pHandler.GetCategory)
	pRoute.GET("/category", middlewares.SetAccount(), pHandler.GetPostOfCategory)
	pRoute.GET("/tag", middlewares.SetAccount(), pHandler.GetPostOfTag)
	pRoute.GET("/user", middlewares.Auth(), pHandler.GetPostOfUser)
	pRoute.GET("/liked", middlewares.Auth(), pHandler.GetPostOfUserLiked)
	pRoute.GET("/info", pHandler.GetPost)
	pRoute.GET("/square", pHandler.PostSquare)
	
}
