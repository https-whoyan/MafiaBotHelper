package config

import (
	"log"
	"os"
	"sync"

	"MafiaBotHelper/internal"
	myLog "MafiaBotHelper/log"

	"github.com/joho/godotenv"
)

type Config struct {
	token string
}

var (
	configOnce sync.Once
)

func LoadConfig() *Config {
	cfg := &Config{}
	configOnce.Do(func() {
		loadEnd()
		token := os.Getenv("BOT_TOKEN")
		cfg.token = token
	})
	return cfg
}

func (c *Config) Run() {
	logger := myLog.NewLogger()
	bot := internal.NewBot(c.token, logger)
	err := bot.Init()

	if err != nil {
		log.Fatal(err)
	}

	bot.Run()
}

func loadEnd() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
