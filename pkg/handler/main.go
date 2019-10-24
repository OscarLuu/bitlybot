package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate checks for new messages that are created
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "hello") {
		MessageDelete(s, m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, "Received")
	}
}

// MessageDelete deletes message
func MessageDelete(s *discordgo.Session, chanID string, mID string) {
	s.ChannelMessageDelete(chanID, mID)
	s.ChannelMessageSend(chanID, "Hit message delete function")
}
