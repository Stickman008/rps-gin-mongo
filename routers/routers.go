package routers

import (
	"rock-paper-scissor/app/tests"
	"rock-paper-scissor/config"

	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(config.GinMode)
	router := gin.New()

	router.GET("/", tests.GetTest)

	return router
}
