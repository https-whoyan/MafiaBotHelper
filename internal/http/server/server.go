package server

import (
	"context"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/https-whoyan/MafiaBotHelper/internal/log"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	addr        string
	server      *fasthttp.Server
	router      *fasthttprouter.Router
	logger      fasthttp.Logger
	middlewares []Middleware
}

func NewHttpServer(cfg *Config) *HttpServer {
	httpLogger := fasthttp.Logger(log.GetLogger())
	return &HttpServer{
		addr: cfg.Addr,
		server: &fasthttp.Server{
			WriteTimeout: cfg.WriteIdle,
			ReadTimeout:  cfg.ReadIdle,
		},
		logger: httpLogger,
		router: fasthttprouter.New(),
	}
}

func (s *HttpServer) Start(_ context.Context) error {
	s.registerMiddlewares()
	s.logger.Printf("http server start with port %v\n", s.addr)
	return s.server.ListenAndServe(s.addr)
}

func (s *HttpServer) registerMiddlewares() {
	handler := s.server.Handler
	for _, m := range s.middlewares {
		handler = m.Next(handler)
	}
	s.server.Handler = handler
}

func (s *HttpServer) Shutdown(_ context.Context) error {
	return s.server.Shutdown()
}
