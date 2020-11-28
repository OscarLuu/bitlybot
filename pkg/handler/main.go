package handler

import (
	"regexp"
	"strings"

	"github.com/OscarLuu/bitlybot/pkg/bitly"
	"github.com/OscarLuu/bitlybot/pkg/scraper"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func regexCompile(s *discordgo.Session, content string, chanID string, messageID string, user string) {
	re := regexp.MustCompile(`http([^\s]+)`)
	if re.MatchString(content) {
		link := re.FindString(content)
		short, err := bitly.Shorten(link)
		log.Infof("creating short link: %v", short)
		if err != nil {
			log.Errorf("Error in creating short linK: %v", err)
			s.ChannelMessageSend(chanID, "Request resulted in an error, please try again.")
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
			message := strings.Replace(content, link, short, 1)
			shortAuthor := " - " + (string(user) + "\n" + scrape + "\n" + message)
			s.ChannelMessageSend(chanID, shortAuthor)

			// deleting the message
			log.Infof("deleting former message: %v", content)
			err = MessageDelete(s, chanID, messageID)
			if err != nil {
				log.Errorf("deleting former message: %v", err)
			}
			log.Infof("deleted former message: %v", content)
		}
	}
}

// OnMessageCreate checks for new messages that are created
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Checks if the message creator is the same as the bot
	// if it is the same then ignore it
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "~shorten") {
		regexCompile(s, m.Content, m.ChannelID, m.Message.ID, m.Author.Username)
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
