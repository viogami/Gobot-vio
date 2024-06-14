package main

import (
	"log"

	"github.com/viogami/Gobot-vio/config"
	"github.com/viogami/Gobot-vio/server"
)

func main() {
	env := config.GetEnv()
	log.Println(env)
	server.Run(env.PORT)
}
