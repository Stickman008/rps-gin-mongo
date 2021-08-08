package challenge

import (
	"log"
	"rock-paper-scissor/app/users"

	"github.com/gin-gonic/gin"
)

func Challenge(ctx *gin.Context) {
	var challenger users.Challenger
	err := ctx.BindJSON(&challenger)
	if err != nil {
		log.Fatal(err)
	}
	// ctx.JSON(http.StatusOK, gin.H{

	// })
}
