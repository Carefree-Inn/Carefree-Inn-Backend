package route

import (
	"gateway/internal/handler/user"
	"github.com/gin-gonic/gin"
)

func userRoute(engine *gin.RouterGroup) {
	uRoute := engine.Group("/user")
	uHandler := user.NewUserHandler()
	
	uRoute.POST("/login", uHandler.Login)
	uRoute.GET("/profile", uHandler.GetProfile)
	uRoute.PUT("/profile", uHandler.UpdateProfile)
}
