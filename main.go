package main

import (
	"Gobot-vio/server"
	"Gobot-vio/tgbot"
	"log"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	server := server.New()
	err := server.Run(port)
	log.Println("HTTP server is running on port", port)
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
	}
	//创建一个tgbot
	tgbot.CreateTgbot()
}
