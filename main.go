package main

import (
	"os"

	"github.com/viogami/Gobot-vio/server"
)

func main() {
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")
	s := server.NewServer(port, redisURL)
	
	s.Run()
}
