package handler

import (
	"regexp"

	"github.com/OscarLuu/bitlybot/pkg/api"
	"github.com/bwmarrin/discordgo"
)

// OnMessageCreate checks for new messages that are created
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Checks if the message creator is the same as the bot
	// if it is the same then ignore it
	if m.Author.ID == s.State.User.ID {
		return
	}

	// regex match for anything that starts with http
	// and has at least one . in it
	re := regexp.MustCompile(`^http*\.*`)
	if re.MatchString(m.Content) {
		short := api.Bitly(m.Content)
		s.ChannelMessageSend(m.ChannelID, short)
	}
}

// MessageDelete deletes message
func MessageDelete(s *discordgo.Session, chanID string, mID string) {
	// this function is designed to delete the link message
	s.ChannelMessageDelete(chanID, mID)
	s.ChannelMessageSend(chanID, "Hit message delete function")
}
