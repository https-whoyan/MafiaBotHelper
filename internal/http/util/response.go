package util

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

func NewSuccessResponse(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func NewErrorResponse(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(http.StatusInternalServerError)
	ctx.SetBodyString(err.Error())
}
