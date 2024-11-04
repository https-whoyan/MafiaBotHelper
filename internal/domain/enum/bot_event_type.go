package enum

//go:generate enumer -type=BotEventType -json -trimprefix BotEventType -transform=snake -output bot_event_status_enumer.go

type BotEventType uint8

const (
	BotEventTypeStayReaction BotEventType = iota
	BotEventTypeSendVote
	BotEventTypeSendTwoVote
	BotEventTypeSendDayVote
)
