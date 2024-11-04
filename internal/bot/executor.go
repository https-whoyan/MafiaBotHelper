package bot

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/https-whoyan/MafiaBot/pkg/commands"
	"github.com/https-whoyan/MafiaBot/pkg/stickers"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/entity"
	"github.com/https-whoyan/MafiaBotHelper/internal/domain/enum"
)

const (
	registrationSticker = stickers.RegistrationPlayerSticker

	oneVoteCommand = commands.VoteGameCommandName
	twoVoteCommand = commands.TwoVoteGameCommandName
	dayVoteCommand = commands.DayVoteGameCommandName
)

func (b *Bot) Handle(event *entity.BotEventWithType) error {
	b.logger.Printf("Got event: %v", event.Type)
	switch event.Type {
	case enum.BotEventTypeStayReaction:
		return b.staySticker(event, registrationSticker)
	case enum.BotEventTypeSendVote:
		return b.sendOneVote(event)
	case enum.BotEventTypeSendTwoVote:
		return b.setTwoVote(event)
	case enum.BotEventTypeSendDayVote:
		return b.sendDayVote(event)
	}
	return nil
}

func (b *Bot) staySticker(event *entity.BotEventWithType, sticker string) error {
	return b.sess.Rest().AddReaction(event.ChannelID, event.MessageID, sticker)
}

func (b *Bot) sendOneVote(event *entity.BotEventWithType) error {
	vote := event.Attr[entity.AttrVoteKey].(string)
	message := "/" + oneVoteCommand + " " + vote
	return b.sendMessage(event, message)
}

func (b *Bot) setTwoVote(event *entity.BotEventWithType) error {
	votes := event.Attr[entity.AttrVoteKey].([]string)
	vote1, vote2 := votes[0], votes[1]
	message := "/" + twoVoteCommand + " " + vote1 + " " + vote2
	return b.sendMessage(event, message)
}

func (b *Bot) sendDayVote(event *entity.BotEventWithType) error {
	vote := event.Attr[entity.AttrVoteKey].(string)
	message := "/" + dayVoteCommand + " " + vote
	return b.sendMessage(event, message)
}

func (b *Bot) sendMessage(event *entity.BotEventWithType, message string) error {
	replyMessage := discord.NewMessageCreateBuilder().SetContent(message).Build()
	_, err := b.sess.Rest().CreateMessage(event.ChannelID, replyMessage)
	return err
}
