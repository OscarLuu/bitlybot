package handler

import (
	"regexp"

	"github.com/OscarLuu/bitlybot/pkg/bitly"
	"github.com/OscarLuu/bitlybot/pkg/scraper"

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
	// blah bit.ly blah
	re := regexp.MustCompile(`http([^\s]+)`)
	if re.MatchString(m.Content) {
		link := re.FindString(m.Content)
		short, err := bitly.Shorten(link)
		log.Infof("creating short link %v", short)
		if err != nil {
			log.Errorf("creating short link %v", err)
			s.ChannelMessageSend(m.ChannelID, "Request resulted in an error, please try again.")
		} else {
			// scraping the title
			log.Infoln("scraping the webpage for the title")
			scrape, err := scraper.ScrapeWebPage(short)
			if err != nil {
				log.Infof("scraping webpage failed: %v", err)
			}
			log.Infof("scraped webpage: %v", scrape)

			// creating and sending the message
			log.Infof("created short link %v", short)
			shortAuthor := " - " + (string(m.Author.Username) + "\n\n" + scrape + "\n" + short)
			s.ChannelMessageSend(m.ChannelID, shortAuthor)

			// deleting the message
			log.Infof("deleting former message: %v", m.Content)
			err = MessageDelete(s, m.ChannelID, m.Message.ID)
			if err != nil {
				log.Errorf("deleting former message: %v", err)
			}
			log.Infof("deleted former message: %v", m.Content)
		}
	}
}

// MessageDelete deletes message
func MessageDelete(s *discordgo.Session, chanID string, mID string) error {
	// this function is designed to delete the link message
	err := s.ChannelMessageDelete(chanID, mID)
	if err != nil {
		return err
	}
	return nil
}
