package server

import (
	fasthttprouter "github.com/fasthttp/router"
)

type Option interface {
	accept(srv *HttpServer)
	name() string
}

type Controller interface {
	Init(r *fasthttprouter.Router)
	Name() string
}

type controller struct{ Controller }

func (c *controller) accept(srv *HttpServer) {
	c.Init(srv.router)
}

func (c *controller) name() string {
	return c.Controller.Name()
}

func WithController(c Controller) Option {
	return &controller{c}
}

func (s *HttpServer) WithOptions(o ...Option) {
	for _, opt := range o {
		s.logger.Printf("accept %v option", opt.name())
		opt.accept(s)
	}
}
