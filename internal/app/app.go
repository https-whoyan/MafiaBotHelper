package app

import (
	"context"

	"github.com/https-whoyan/MafiaBotHelper/internal/bot"
	"github.com/https-whoyan/MafiaBotHelper/internal/config"
	"github.com/https-whoyan/MafiaBotHelper/internal/http/server"
)

type App struct {
	Server *server.HttpServer
	Bots   []*bot.Bot
}

func NewApp(cfg *config.Config) *App {
	bots := make([]*bot.Bot, len(cfg.BotsConfigs))
	for botNum, botCfg := range cfg.BotsConfigs {
		bots[botNum] = bot.NewBot(botCfg, cfg.Logger)
	}
	srv := server.NewHttpServer(cfg.SrvConfig)
	return &App{
		Server: srv,
		Bots:   bots,
	}
}

func (app *App) Start(ctx context.Context) error {
	var botFns []func(ctx context.Context) error
	for _, b := range app.Bots {
		botStartFunc := func(ctx context.Context) error {
			err := b.Init()
			if err != nil {
				return err
			}
			return b.Run(ctx)
		}
		botFns = append(botFns, botStartFunc)
	}
	fns := []func(ctx context.Context) error{
		app.Server.Start,
	}
	fns = append(fns, botFns...)
	return paralleledRun(ctx, fns)
}

func (app *App) Shutdown(ctx context.Context) error {
	fns := []func(ctx context.Context) error{
		app.Server.Shutdown,
	}
	for _, b := range app.Bots {
		fns = append(fns, func(ctx context.Context) error {
			return b.Close(ctx)
		})
	}
	return paralleledRun(ctx, fns)
}
