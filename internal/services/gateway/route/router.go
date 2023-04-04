package route

import (
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Route(engine *gin.Engine) {
	engine.Use(middlewares.SetUuid(),
		middlewares.Logger(), gin.Recovery(),
	)
	
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "page is not exist")
	})
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	base := engine.Group("/inn/api/v1")
	userRoute(base)
	postRoute(base)
}
