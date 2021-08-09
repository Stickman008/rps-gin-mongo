package main

import (
	// "rock-paper-scissor/database"
	"rock-paper-scissor/routers"
)

func main() {
	// database.Init()

	server := routers.GenerateRouter()

	server.Run(":8080")
}
