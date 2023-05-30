package route

import (
	"gateway/internal/handler/user"
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func userRoute(engine *gin.RouterGroup) {
	uRoute := engine.Group("/user")
	uHandler := user.NewUserHandler()
	
	uRoute.POST("/login", uHandler.Login)
	uRoute.GET("/profile", middlewares.Auth(), uHandler.GetProfile)
	uRoute.PUT("/profile", middlewares.Auth(), uHandler.UpdateProfile)
}
