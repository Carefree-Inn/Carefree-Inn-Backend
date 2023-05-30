package middlewares

import (
	"gateway/pkg"
	"github.com/gin-gonic/gin"
	"time"
)

func SetUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("uuid", pkg.GenerateUUid())
	}
}

func SetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			return
		}
		claim, err := pkg.ParseToken(auth)
		if err != nil {
			return
		}
		if claim.ExpiresAt < time.Now().Unix() {
			return
		}
		
		c.Set("account", claim.Account)
	}
}
