package entity

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/enum"
)

const (
	BotEventTypeKey = "bot_event_type"
	AttrVoteKey     = "votes"
)

type BotEvent struct {
	BotID     int
	MessageID snowflake.ID
	GuildID   snowflake.ID
	ChannelID snowflake.ID
	Message   string
	Attr      map[string]interface{}
}

type BotEventWithType struct {
	BotEvent
	Type enum.BotEventType
}
