package bots

import (
	"github.com/https-whoyan/MafiaBotHelper/internal/bot"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/entity"
)

type Service interface {
	ProcessEvent(botNum int, event *entity.BotEventWithType) error
}

type service struct {
	bots bot.Bots
}

func NewService(bots bot.Bots) Service {
	return &service{
		bots: bots,
	}
}

func (s service) ProcessEvent(botNum int, event *entity.BotEventWithType) error {
	return s.bots.GetByNum(botNum).Handle(event)
}
