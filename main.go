/*
Copyright © 2019

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
	"github.com/OscarLuu/bitlybot/lib"
	"github.com/bwmarrin/discordgo"
)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/OscarLuu/bitlybot/cmd"
// Vars used for command line params
var (
	Token string
	Bitly string
)

func init() {
	flag.StringVar(&Token, "token", "", "Auth Token")
	flag.StringVar(&Bitly, "bitly", "", "Auth Token")
	flag.Parse()
}
>>>>>>> tweaks

func main() {
	cmd.Execute()
	// initialize discordgo bot session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session ", err)
		return
	}

	dg.AddHandler(lib.OnMessageCreate)

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
