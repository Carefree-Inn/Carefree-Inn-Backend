package middlewares

import (
	"fmt"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		fmt.Println("trace")
		log.Trace(log.WithFields(logrus.Fields{
			"status_code":  c.Writer.Status(),
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"used_time":    strconv.Itoa(int(end.Sub(start).Milliseconds())) + "ms",
			"ip":           c.ClientIP(),
			"x-request-id": c.MustGet("uuid").(string),
		}))
	}
}
