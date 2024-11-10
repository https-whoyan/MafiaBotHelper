package controller

import (
	botsC "github.com/https-whoyan/MafiaBotHelper/internal/http/controller/bots"
	"github.com/https-whoyan/MafiaBotHelper/internal/http/server"
	botsS "github.com/https-whoyan/MafiaBotHelper/internal/service/bots"
)

func Accept(
	s *server.HttpServer,
	botsS botsS.Service,
) {
	s.WithOptions(
		server.WithController(botsC.New(botsS)),
	)
}
