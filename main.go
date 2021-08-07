package main

import (
	"rock-paper-scissor/routers"
)

func main() {
	server := routers.GenerateRouter()
	server.Run(":8080")
}
