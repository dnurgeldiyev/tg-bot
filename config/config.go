package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Config struct {
	TgBot   TgBot   `json:"tgBot"`
	Storage Storage `json:"storage"`
}
type TgBot struct {
	Host  string `json:"host"`
	Token string
}
type Storage struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func NewConfig() (*Config, error) {

	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	tgBotTokenToken := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *tgBotTokenToken == "" {
		log.Fatal("token is not specified")
	}

	cfg.TgBot.Token = *tgBotTokenToken

	return &cfg, nil
}
