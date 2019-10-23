package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Vars used for command line params
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "token", "", "Auth Token")
	flag.Parse()
}

func main() {
	// initialize discordgo bot session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session ", err)
		return
	}

	dg.AddHandler(messageCreate)

	// open connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("error with connection ", err)
		return
	}

	// allow control c termination
	fmt.Println("Running, awaiting CTRL-C to turn shutdown.")
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Exiting")

	// Close connection
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is ping reply with pong
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong")
	}

	// If the message is pong reply with ping"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping")
	}
}
