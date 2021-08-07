package routers

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/users"
	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	router.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
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

func (s *userService) Save(user users.User) users.User {
	s.users = append(s.users, user)
}
func (s *userService) FindAll() []users.User {
	return s.users
}
