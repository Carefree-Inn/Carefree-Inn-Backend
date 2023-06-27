package route

import (
	"gateway/internal/handler/notification"
	"gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func notificationRoute(engine *gin.RouterGroup) {
	notificationHandler := notification.NewNotificationHandler()
	
	engine.GET("/notification", middlewares.Auth(), notificationHandler.SendNotification)
	engine.GET("/notification/history", middlewares.Auth(), notificationHandler.GetNotificationHistory)
}
