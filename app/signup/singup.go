package signup

import (
	"log"
	"net/http"
	"rock-paper-scissor/app/users"

	"github.com/gin-gonic/gin"
)

func SignUpWithUsernameAndPassword(ctx *gin.Context) {
	var user users.User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	id, err := AddUserToDatabase(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
	//use signup handlers
}
