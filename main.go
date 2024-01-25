package main

import (
	"Gobot-vio/tgbot"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	//创建一个tgbot
	tgbot.CreateTgbot()
}
