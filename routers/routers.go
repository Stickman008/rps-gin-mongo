package routers

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/users"
	"github.com/gin-gonic/gin"
)
var {
	userService service.UserService = service.New()
	userController controller.UserController = conreoller.New(userService)

}
func GenerateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	return router
}

type UserService interface {
	Save(users.User) users.User
	FindAll() []users.User
}

type userService struct {
	users []users.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user users.User) users.User {
	s.users = append(s.users, user)
}
func (service *userService) FindAll() []users.User {
	return s.users
}

type UserController interface {
	Save(users.User) users.User
	FindAll() []users.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []users.User{
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) users.User{
	var user users.User
	ctx.BindJSON(&user)
	c.service.Save()
	return user
}
