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
	Save(users.User)
}
