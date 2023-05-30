package route

import (
	"gateway/internal/handler/notification"
	"github.com/gin-gonic/gin"
)

func notificationRoute(engine *gin.RouterGroup) {
	notificationHandler := notification.NewNotificationHandler()
	
	engine.GET("/notification", notificationHandler.SendNotification)
}
