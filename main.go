package main

import (
	"os"

	_ "github.com/viogami/viogo/conf"
	"github.com/viogami/viogo/server"
)

func main() {
	port := os.Getenv("PORT")
	s := server.NewServer(port)

	redisURL := os.Getenv("REDISCLOUD_URL")
	s.WithRedis(redisURL)

	s.Run()
}
