package bots

import (
	"encoding/json"
	"github.com/disgoorg/snowflake/v2"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/entity"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/enum"
)

type handleEventBody struct {
	BotID     int               `json:"bot_id"`
	EventType enum.BotEventType `json:"event_type"`
	MessageID int               `json:"message_id"`
	ChannelID int               `json:"channel_id"`
	GuildID   int               `json:"guild_id"`
	Message   string            `json:"message"`
}

func (h *handleEventBody) toEntity() *entity.BotEventWithType {
	return &entity.BotEventWithType{
		BotEvent: entity.BotEvent{
			BotID:     h.BotID,
			MessageID: snowflake.ID(h.MessageID),
			ChannelID: snowflake.ID(h.ChannelID),
			GuildID:   snowflake.ID(h.GuildID),
			Message:   h.Message,
		},
		Type: h.EventType,
	}
}

func (h *handleEventBody) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, h)
}

func (h handleEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(h)
}
