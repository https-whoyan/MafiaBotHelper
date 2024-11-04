package config

import (
	"github.com/https-whoyan/MafiaBotHelper/internal/http/server"
	"log"
	"os"
	"sync"

	"github.com/https-whoyan/MafiaBotHelper/internal/bot"
	"github.com/joho/godotenv"
)

type Config struct {
	SrvConfig   *server.Config
	Logger      *log.Logger
	BotsConfigs []*bot.Config
}

var (
	configOnce sync.Once
)

const (
	botsCount = 2
)

func LoadConfig() *Config {
	cfg := &Config{}
	configOnce.Do(func() {
		loadEnd()
	})
	cfg.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	for numOfBot := range botsCount {
		botCfg := bot.NewConfig(numOfBot)
		cfg.BotsConfigs = append(cfg.BotsConfigs, botCfg)
	}
	cfg.SrvConfig = server.NewConfig()
	return cfg
}

func loadEnd() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
