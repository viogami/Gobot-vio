package config

import (
	"flag"
	"log"
	"testing"
)

func Test(t *testing.T) {
	appConfig := flag.String("config", "./app.yaml", "application config path")
	k, _ := ConfigParse(*appConfig)
	log.Println("打印：", k.Server.Port)
}
