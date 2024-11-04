package bots

import (
	fasthttprouter "github.com/fasthttp/router"
	"github.com/https-whoyan/MafiaBotHelper/internal/service/bots"
)

type Controller struct {
	botsService bots.Service
}

func New(botsService bots.Service) *Controller {
	return &Controller{
		botsService: botsService,
	}
}

func (c *Controller) Init(r *fasthttprouter.Router) {
	group := r.Group("/bots")
	group.POST("/post", c.handleEvent)
}
