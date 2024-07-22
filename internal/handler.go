package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/https-whoyan/MafiaCore/game"
	"github.com/https-whoyan/MafiaCore/roles"
)

func (b *Bot) handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := m.Content
	messageID := m.Message.ID
	executedGuildID := m.GuildID
	executedChannelID := m.ChannelID

	if isNeedToStayReaction(message) {
		b.stayRegistrationReaction(executedGuildID, executedChannelID, messageID)
		return
	}
	if !hasAMention(message, b.userID) {
		return
	}

	if isInvasionToVote(message) {
		err := b.setVote(executedGuildID, executedChannelID)
		if err != nil {
			b.logger.Log(context.Background(), "Error sending message: ", err)
			return
		}
		b.logger.Log(context.Background(), "Message of vote sent")
		return
	}

	// Then, it is a new game initialized
	botRole, err := findBotRole(message)
	if err != nil {
		b.logger.Log(context.Background(), "Error find role: ", err)
		return
	}
	b.Lock()
	defer b.Unlock()
	b.roles[executedGuildID] = botRole
}

func (b *Bot) stayRegistrationReaction(guildID, channelID, messageID string) {
	registrationSticker := "😁"
	err := b.sess.MessageReactionAdd(channelID, messageID, registrationSticker)
	if err != nil {
		b.logger.Log(context.Background(), "Error adding sticker: ", err)
		return
	}
	b.logger.Log(context.Background(), "Message of stay reaction sent")
}

func hasAMention(message string, memberID string) bool {
	return strings.Contains(message, "<@")
}

func isInvasionToVote(message string) bool {
	return strings.Contains(message, "It's your turn to Vote.")
}

func (b *Bot) setVote(executedGuildID, executedChannel string) error {
	b.RLock()
	role, isContains := b.roles[executedGuildID]
	b.RUnlock()
	if !isContains {
		return errors.New("in guild id bot isn't set a role")
	}

	var message string
	var emptyVote = game.EmptyVoteStr

	if role.IsTwoVotes {
		message = fmt.Sprintf("/two_vote %v %v", emptyVote, emptyVote)
	} else {
		message = fmt.Sprintf("/vote %v %v", emptyVote, emptyVote)
	}

	_, err := b.sess.ChannelMessageSend(executedChannel, message)
	return err
}

func isNeedToStayReaction(message string) bool {
	return strings.Contains(message, "Registration has begun.")
}

func findBotRole(message string) (*roles.Role, error) {
	allRoles := roles.MappedRoles

	var botRole *roles.Role
	for _, role := range allRoles {
		if strings.Contains(message, role.Name) {
			botRole = role
			break
		}
	}
	if botRole == nil {
		return nil, errors.New("no role found")
	}
	return botRole, nil
}
