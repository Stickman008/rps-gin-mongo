package routers

import "github.com/gin-gonic/gin"

func GenerateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return router
}
