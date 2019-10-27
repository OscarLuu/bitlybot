/*
Copyright © 2019

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/OscarLuu/bitlybot/pkg/bitly"
	"os"
	"os/signal"
	"syscall"

	"github.com/OscarLuu/bitlybot/pkg/handler"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	discordToken string
	bitlyToken   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bitlybot",
	Short: "BitlyBot for Discord Chat.",
	Long: `BitlyBot converts long ugly links to short and friendly ones.
It does this by leveraging the public Bitly API.`,
	Run: func(cmd *cobra.Command, args []string) {

		// set bitly api token
		bitly.SetToken(bitlyToken)

		// create new discord client
		discord, err := discordgo.New(fmt.Sprintf("Bot %s", discordToken))
		if err != nil {
			log.Fatalf("getting discord session %v\n", err)
		}

		discord.AddHandler(handler.OnMessageCreate)

		err = discord.Open()
		if err != nil {
			log.Fatalf("opening websocket connection %v\n", err)
		}

		log.Infoln("ctrl-c to terminate")

		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		<-done

		log.Infoln("exiting")
		discord.Close()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&discordToken, "token", "", "Discord OAuth Token")
	rootCmd.Flags().StringVar(&bitlyToken, "bitly-token", "", "Bitly API Token")
}
