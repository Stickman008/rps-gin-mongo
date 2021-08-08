package signup

import (
	"rock-paper-scissor/app/users"

	"github.com/gin-gonic/gin"
)

func SignUpWithUsernameAndPassword(ctx *gin.Context) {
	var user users.User
	ctx.BindJSON(&user)
	AddUserToDatabase(&user)
	//use signup handlers
}
