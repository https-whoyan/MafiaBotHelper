package server

import (
	"os"
	"time"
)

type Config struct {
	Addr      string
	WriteIdle time.Duration
	ReadIdle  time.Duration
}

func NewConfig() *Config {
	return &Config{
		Addr:      os.Getenv("HTTP_ADDR"),
		WriteIdle: 5 * time.Second,
		ReadIdle:  5 * time.Second,
	}
}
