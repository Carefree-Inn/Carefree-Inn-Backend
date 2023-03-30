package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateUUid() string {
	uid := uuid.New()
	return uid.String()
}

func GetUUid(c *gin.Context) string {
	return c.MustGet("uuid").(string)
}
