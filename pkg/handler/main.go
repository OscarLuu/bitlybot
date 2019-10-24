package handler

import (
	"regexp"
	
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
		// s.ChannelMessageSend(m.ChannelID, m.Content)
		// this is where we want to call bitly api and pass the link as string
		// should return here and have channel message send the short link
		// calls ChannelMessageDelete to delete the users link
		// append the username of the poster to the message send
	}
}

// MessageDelete deletes message
func MessageDelete(s *discordgo.Session, chanID string, mID string) {
	// this function is designed to delete the link message
	s.ChannelMessageDelete(chanID, mID)
}
