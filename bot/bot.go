package bot

import (
	"fmt"
	"strings"

	"github.com/mostwantedrbx/discord-go/config"

	"github.com/bwmarrin/discordgo"
)

//	init some variables
var BotID string
var goBot *discordgo.Session

//	this function gets called from the main.go file
func Start() {

	//	create a new discord session
	goBot, err := discordgo.New("Bot " + config.Token)

	// error checking
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	//	tells discordgo what function will process messages
	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running.")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//	make sure the bot isn't going to trigger itself and check to make sure the bots prefix was used
	if m.Author.ID == BotID || !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}
	//	sanity check
	fmt.Println("caught message")

	//	save some time re-writing this
	var cont = strings.ToLower(m.Content)

	//	long list of if statements to check what we need to do
	if strings.Contains(cont, "ping") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

}
