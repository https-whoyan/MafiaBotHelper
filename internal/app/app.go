package app

import (
	"context"
	"github.com/https-whoyan/MafiaBotHelper/internal/bot"
	"github.com/https-whoyan/MafiaBotHelper/internal/config"
	"github.com/https-whoyan/MafiaBotHelper/internal/http/controller"
	"github.com/https-whoyan/MafiaBotHelper/internal/http/server"
	"github.com/https-whoyan/MafiaBotHelper/internal/log"
	"github.com/https-whoyan/MafiaBotHelper/internal/service/bots"
)

type App struct {
	logger log.Logger

	botsS bots.Service
	bots  bot.Bots

	server *server.HttpServer
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	var (
		app = &App{}
		err error
	)

	for _, f := range []func(ctx context.Context, cfg *config.Config) error{
		app.initLogger,
		app.initServices,
		app.initHttp,
	} {
		err = f(ctx, cfg)
		if err != nil {
			return nil, err
		}
	}
	return app, nil
}

func (app *App) initLogger(_ context.Context, cfg *config.Config) error {
	app.logger = cfg.Logger
	return nil
}

func (app *App) initServices(ctx context.Context, cfg *config.Config) error {
	appBots := make([]*bot.Bot, len(cfg.BotsConfigs))
	for botNum, botCfg := range cfg.BotsConfigs {
		appBots[botNum] = bot.NewBot(botCfg)
	}
	app.bots = appBots
	app.botsS = bots.NewService(appBots)
	return nil
}

func (app *App) initHttp(ctx context.Context, cfg *config.Config) error {
	srv := server.NewHttpServer(cfg.SrvConfig)
	app.server = srv
	return nil
}

func (app *App) Start(ctx context.Context) error {
	app.logger.Println("Starting")
	var botFns []func(ctx context.Context) error
	for _, b := range app.bots {
		botStartFunc := func(ctx context.Context) error {
			err := b.Init()
			if err != nil {
				return err
			}
			return b.Run(ctx)
		}
		botFns = append(botFns, botStartFunc)
	}
	registerHttpServer := func(_ context.Context) error {
		controller.Accept(app.server, app.botsS)
		return nil
	}
	fns := []func(ctx context.Context) error{
		app.server.Start,
		registerHttpServer,
	}
	fns = append(fns, botFns...)
	return paralleledRun(ctx, fns)
}

func (app *App) Shutdown(ctx context.Context) error {
	app.logger.Println("Shutdown")
	fns := []func(ctx context.Context) error{
		app.server.Shutdown,
	}
	for _, b := range app.bots {
		fns = append(fns, func(ctx context.Context) error {
			return b.Close(ctx)
		})
	}
	return paralleledRun(ctx, fns)
}
