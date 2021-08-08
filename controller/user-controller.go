package controller

import (
	"rock-paper-scissor/app/users"
	"rock-paper-scissor/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []users.User
	Save(ctx *gin.Context) users.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []users.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) users.User {
	var user users.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}
