package bot

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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
	//	create a new discord session
	goBot, err := discordgo.New("Bot " + config.Token)
	//	error checking
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second)
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second)
	}

	BotID = u.ID

	//	tells discordgo what function will process messages
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second)
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
	//	command splits the message recieved into the command, on command[0] and the arguments on command[1]
	switch command := strings.SplitAfter(cont, " "); command[0] {

	case "!ping ":
		//	used to test if the bot is on
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")

	case "!echo ":
		if len(command) > 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, command[1])
		}

	case "!pokemon ":
		//	silly command for funsies
		_, _ = s.ChannelMessageSend(m.ChannelID, storage.ReturnRandomPokemon())

	case "!roll ":
		//	rolls command[1] amount of dice
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
		// download the file from the url
		net.DownloadFile(command[1], "tacos.png")

		//	run the python script to convert the image, and it saves it in a txt file.
		pyscripts.RunScript("convert")

		//	gets the results
		fmt.Println("Opening a file ")
		var file, err = os.ReadFile("./ascii-image.txt")
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second)
			return
		}

		//	send the contents to the pastebin function to be pasted
		p := storage.Pastebin{}
		link, err := p.Put(string(file), "Ascii Image")
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "The image failed to convert! Let my owner know!")
			time.Sleep(time.Second)
			return
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Your image has been pasted at: "+link)
		}

	}
}
