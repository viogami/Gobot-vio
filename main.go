package main

import (
	"os"

	_ "github.com/viogami/Gobot-vio/conf"
	"github.com/viogami/Gobot-vio/server"
)

func main() {
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDISCLOUD_URL")
	s := server.NewServer(port, redisURL)

	s.Run()
}
