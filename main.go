package main

import (
	"os"

	"github.com/viogami/Gobot-vio/server"
)

func main() {
	var port = os.Getenv("PORT")

	server.Run(port)
}
