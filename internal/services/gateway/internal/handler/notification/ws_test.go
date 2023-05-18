package notification

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestSendNotification(t *testing.T) {
	nh := NewNotificationHandler()
	
	engine := gin.New()
	
	engine.GET("/ws", nh.SendNotification)
	
	engine.Run(":8080")
}
