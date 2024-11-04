package bots

import (
	"github.com/https-whoyan/MafiaBotHelper/internal/http/util"
	"github.com/valyala/fasthttp"
)

func (c *Controller) handleEvent(ctx *fasthttp.RequestCtx) {
	var handleBody handleEventBody
	err := handleBody.UnmarshalJSON(ctx.Request.Body())
	if err != nil {
		util.NewErrorResponse(ctx, err)
		return
	}
	err = c.botsService.ProcessEvent(handleBody.BotID, handleBody.toEntity())
	if err != nil {
		util.NewErrorResponse(ctx, err)
		return
	}
	util.NewSuccessResponse(ctx)
}
