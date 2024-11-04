package main

import (
	"context"
	"log"

	"github.com/https-whoyan/MafiaBotHelper/internal/app"
	"github.com/https-whoyan/MafiaBotHelper/internal/config"
)

var (
	ctx = context.Background()
)

func main() {
	cfg := config.LoadConfig()
	apl := app.NewApp(cfg)
	err := apl.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = apl.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
