package middlewares

import (
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			internal.Custom(c, http.StatusUnauthorized, nil, errno.UserNotVerifyError.Error())
			c.Abort()
			return
		}
		claim, err := pkg.ParseToken(auth)
		if err != nil {
			internal.Custom(c, http.StatusUnauthorized, nil, errno.TokenNotValidate.Error())
			c.Abort()
			return
		}
		
		if claim.ExpiresAt < time.Now().Unix() {
			internal.Custom(c, http.StatusUnauthorized, nil, errno.TokenNotValidate.Error())
			c.Abort()
			return
		}
		
		c.Set("account", claim.Account)
		c.Next()
	}
}
