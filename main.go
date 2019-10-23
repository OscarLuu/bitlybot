package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	B "./lib"
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

	dg.AddHandler(B.MessageCreate)

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
