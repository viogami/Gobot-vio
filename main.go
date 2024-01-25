package main

import (
	"Gobot-vio/server"
	"os"
)

func main() {
	var port = os.Getenv("PORT")

	server := server.New()
	server.Run(port)
}
