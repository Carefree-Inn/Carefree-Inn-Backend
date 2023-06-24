package route

import (
	"gateway/config"
	"gateway/internal/handler/qiniu"
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Route(engine *gin.Engine, config *config.Config) {
	engine.Use(middlewares.SetUuid(),
		middlewares.Logger(), gin.Recovery(),
	)
	
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "page is not exist")
	})
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	base := engine.Group("/inn/api/v1")
	base.GET("/upload/token", qiniu.NewQiNiuHandler(config.QiNiu.AccessKey, config.QiNiu.SecretKey, config.QiNiu.Bucket).GetToken)
	
	userRoute(base)
	postRoute(base)
	userPostRoute(base)
	notificationRoute(base)
}
