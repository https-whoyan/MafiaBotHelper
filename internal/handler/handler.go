package handler

import (
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/entity"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/enum"
)

type Handler interface {
	Handle(event *entity.BotEvent) error
}

type TyppedHandler interface {
	Handle(event *entity.BotEventWithType) error
}

type EventTypeDefiner interface {
	Define(event *entity.BotEvent) (Type *enum.BotEventType)
}
