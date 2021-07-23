package bot

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mostwantedrbx/discord-go/config"
	"github.com/mostwantedrbx/discord-go/net"
	"github.com/mostwantedrbx/discord-go/pyscripts"
	"github.com/mostwantedrbx/discord-go/storage"

	_ "embed"
)

//	init some variables
var BotID string

//var goBot *discordgo.Session

//	this function gets called from the main.go file
func Start() {
	//pyscripts.RunScript()
	//	create a new discord session
	goBot, err := discordgo.New("Bot " + config.Token)

	//	error checking
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
	switch command := strings.SplitAfter(cont, " "); command[0] {
	case "!ping ":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")

	case "!echo ":
		if len(command) > 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, command[1])
		}
	case "!pokemon ":
		_, _ = s.ChannelMessageSend(m.ChannelID, storage.ReturnRandomPokemon())
	case "!roll ":
		if b, err := strconv.Atoi(command[1]); err == nil {
			var a = rand.Intn(6 * b)
			for ok := true; ok; ok = (a < b) {
				a = rand.Intn(6 * b)
				fmt.Println("rerolling die...")
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, "You rolled "+strconv.Itoa(b)+" dice. \nThe result was: "+strconv.Itoa(a))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You need to supply the number of dice to roll.\nFor example, for three dice: !roll 3")
		}

	case "!convert ":
		fmt.Println("tacos")
		net.DownloadFile(command[1], "tacos.png")
		pyscripts.RunScript("convert")

		fmt.Println("Opening a file ")
		var file, err = os.ReadFile("./ascii-image.txt")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p := pyscripts.Pastebin{}
		link, err := p.Put(string(file), "Ascii Image")

		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "The image failed to convert! Let my owner know!")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, "Your image has been pasted at: "+link)

	}
}
