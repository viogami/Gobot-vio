package main

import (
	"github.com/viogami/Gobot-vio/config"
	"github.com/viogami/Gobot-vio/server"
)

func main() {

	env := config.GetEnv()

	server.Run(env.PORT)
}
