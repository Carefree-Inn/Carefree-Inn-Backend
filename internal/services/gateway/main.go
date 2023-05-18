package main

import (
	_ "gateway/docs" // necessary
	"gateway/internal/handler/notification"
	"github.com/gin-gonic/gin"
)

//	@Title			Inn
//	@Version		1.0
//	@Description	Inn
//	@Host			139.196.30.123
//	@BasePath		/inn/api/v1
func main() {
	nh := notification.NewNotificationHandler()
	
	engine := gin.New()
	
	engine.GET("/ws", nh.SendNotification)
	
	engine.Run(":8080")
}
