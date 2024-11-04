package server

import (
	"context"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	addr        string
	server      *fasthttp.Server
	router      *fasthttprouter.Router
	middlewares []Middleware
}

func NewHttpServer(cfg *Config) *HttpServer {
	return &HttpServer{
		addr: cfg.Addr,
		server: &fasthttp.Server{
			WriteTimeout: cfg.WriteIdle,
			ReadTimeout:  cfg.ReadIdle,
		},
		router: fasthttprouter.New(),
	}
}

func (s *HttpServer) Start(_ context.Context) error {
	s.registerMiddlewares()
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
