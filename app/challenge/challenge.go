package challenge

import (
	"log"
	"net/http"
	"rock-paper-scissor/app/users"

	"github.com/gin-gonic/gin"
)

func ChallengeToOppenent(ctx *gin.Context) {
	var challenger users.Challenger
	if isChallenged(challenger) {

		return
	}

	err := ctx.BindJSON(&challenger)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func isChallenged(user users.Challenger) bool {
	// check on server is opponent user parameter have challenge this user
	return false
}

func isWinningAgainst(challenger users.Challenger) bool {
	// return true if this user is winning against challenger
	return false
}
