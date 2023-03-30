package middlewares

import (
	"gateway/pkg"
	"github.com/gin-gonic/gin"
)

func SetUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("uuid", pkg.GenerateUUid())
	}
}
