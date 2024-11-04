package server

import (
	fasthttprouter "github.com/fasthttp/router"
)

type Option interface {
	accept(srv *HttpServer)
}

type Controller interface {
	Init(r *fasthttprouter.Router)
}

type controller struct{ Controller }

func (c *controller) accept(srv *HttpServer) {
	c.Init(srv.router)
}

func WithController(c Controller) Option {
	return &controller{c}
}
