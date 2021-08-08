package routers

import (
	"rock-paper-scissor/app/tests"
	"rock-paper-scissor/config"

	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(config.GinMode)
	router := gin.New()

	// router.Use(gin.Recovery(), middlewares.BasicAuth())

	router.GET("/", tests.GetTest)
	router.POST("/challenge", func(ctx *gin.Context) {

	})
	router.GET("/ranking", func(ctx *gin.Context) {

	})

	return router
}
