package bot

import (
	"os"
	"strconv"
)

type Config struct {
	Num   int
	Token string
}

func NewConfig(num int) *Config {
	key := "BOT_TOKEN_" + strconv.Itoa(num)
	return &Config{
		Num:   num,
		Token: os.Getenv(key),
	}
}
