package main

import (
	_ "gateway/docs" // necessary
	"gateway/pkg/log"
	"gateway/route"
	"github.com/gin-gonic/gin"
)

//	@Title			Inn
//	@Version		1.0
//	@Description	Inn
//	@Host			139.196.30.123
//	@BasePath		/inn/api/v1
func main() {
	log.NewLogger()
	engine := gin.New()
	route.Route(engine)
	engine.Run(":8080")
}
