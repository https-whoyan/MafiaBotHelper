package config

import (
	goLog "log"
	"sync"

	"github.com/https-whoyan/MafiaBotHelper/internal/bot"
	"github.com/https-whoyan/MafiaBotHelper/internal/http/server"
	"github.com/https-whoyan/MafiaBotHelper/internal/log"
	"github.com/joho/godotenv"
)

type Config struct {
	SrvConfig   *server.Config
	Logger      log.Logger
	BotsConfigs []*bot.Config
}

var (
	configOnce sync.Once
)

const (
	botsCount = 1
)

func LoadConfig() *Config {
	cfg := &Config{}
	configOnce.Do(func() {
		loadEnd()
	})
	log.SetLogger(log.NewDefaultLogger())
	cfg.Logger = log.GetLogger()
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
		goLog.Fatal("Error loading .env file")
	}
}
