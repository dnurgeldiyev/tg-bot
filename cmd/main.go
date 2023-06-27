package main

import (
	"dovran/tg-bot/config"
	"dovran/tg-bot/internal/app"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Start(cfg)
}
