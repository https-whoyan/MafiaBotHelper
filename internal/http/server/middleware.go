package server

import "github.com/valyala/fasthttp"

type Middleware interface {
	Next(ctx fasthttp.RequestHandler) fasthttp.RequestHandler
}
