package main

import (
	"flag"
	"fmt"

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

}
