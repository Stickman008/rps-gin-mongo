package routers

import (
	"rock-paper-scissor/app/signup"
	"rock-paper-scissor/app/tests"
	"rock-paper-scissor/config"
	"rock-paper-scissor/controller"
	"rock-paper-scissor/service"

	"github.com/gin-gonic/gin"
)

var (
	usersService    service.UserService       = service.New()
	usersController controller.UserController = controller.New(usersService)
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(config.GinMode)
	// router := gin.New()
	router := gin.Default()

	// router.Use(gin.Recovery(), middlewares.BasicAuth())

	router.GET("/", tests.GetTest)

	router.POST("/signup", signup.SignUpWithUsernameAndPassword)

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, usersController.FindAll())
	})

	router.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, usersController.Save(ctx))
	})

	// router.POST("/challenge", challenge.Challenge)
	// router.POST("/challenge", challenge.Challenge)
	// router.GET("/ranking", func(ctx *gin.Context) {

	// })

	return router
}
