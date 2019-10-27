package handler

import (
	"github.com/OscarLuu/bitlybot/pkg/bitly"
	"regexp"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// OnMessageCreate checks for new messages that are created
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Checks if the message creator is the same as the bot
	// if it is the same then ignore it
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Instead of searching for anything that starts with ^http
	// we should parse the link out of m.Content
	re := regexp.MustCompile(`^http*\.*`)
	if re.MatchString(m.Content) {
		short, err := bitly.Shorten(m.Content)
		if err != nil {
			log.Errorf("creating short link %v\n", err)
			s.ChannelMessageSend(m.ChannelID, "Request resulted in an error, please try again.")
		} else {
			shortAuthor := short + " - Linked By: " + (string(m.Author.Username))
			s.ChannelMessageSend(m.ChannelID, shortAuthor)
			MessageDelete(s, m.ChannelID, m.Message.ID)
		}
	}
}

// MessageDelete deletes message
func MessageDelete(s *discordgo.Session, chanID string, mID string) {
	// this function is designed to delete the link message
	s.ChannelMessageDelete(chanID, mID)
}
