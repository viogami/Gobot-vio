package main

import (
	"Gobot-vio/server"
	"Gobot-vio/tgbot"
	"os"
)

func main() {
	var port = os.Getenv("PORT")

	//创建一个tgbot
	tgbot.CreateTgbot()

	//创建一个http server
	server := server.New()
	server.Run(port)
}
