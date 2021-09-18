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
	"github.com/rs/zerolog/log"

	_ "embed"
)

//	init some variables
var botID string

//var goBot *discordgo.Session

//	this function gets called from the main.go file
func Start() {
	//	create a new discord session
	goBot, err := discordgo.New("Bot " + config.Token)

	//	error checking
	if err != nil {
		log.Logger.Fatal().Msg("Bot could not be started")
	}

	u, err := goBot.User("@me")

	if err != nil {
		log.Logger.Fatal().Msg("Bot could not find its user")
	}

	botID = u.ID

	//	tells discordgo what function will process messages

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(channelUpdateHandler)
	goBot.AddHandler(channelLeave)

	err = goBot.Open()
	if err != nil {
		log.Logger.Fatal().Msg("Bot could not add a message handler")
		return
	}

	log.Logger.Info().Msg("Bot is now running")
}

func channelLeave(s *discordgo.Session, m *discordgo.Event) {
	//Dunno how to track when the person leaves the channel.
}

func channelUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		log.Logger.Warn().Caller().Msg("Could not get channels")
		return
	}

	for i := 0; i < len(channels)-1; i++ {
		if channels[i].ID == m.ChannelID && channels[i].Name == "Dynamic Channel" {
			c, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
				Name:     channels[i].Name + " 1", //TODO: Gonna make this more dynamic
				Type:     2,
				ParentID: channels[i].ParentID,
			})
			if err != nil {
				log.Logger.Warn().Msg("Couldn't create channel\n" + err.Error())
			}
			s.GuildMemberMove(m.GuildID, m.UserID, &c.ID)
		}
	}

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.GuildChannels(m.GuildID)
	//	make sure the bot isn't going to trigger itself and check to make sure the bots prefix was used
	if m.Author.ID == botID || !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}

	//	save some time re-writing this
	var cont = strings.ToLower(m.Content)

	//	long list of if statements to check what we need to do
	//	command splits the message recieved into the command, on command[0] and the arguments on command[1]
	switch command := strings.SplitAfter(cont, " "); command[0] {

	case "!ping":
		//	used to test if the bot is on
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			log.Logger.Warn().Caller().Msg("Message failed to send")
		}

	case "!echo ":
		if len(command) > 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, command[1])
			if err != nil {
				log.Logger.Warn().Caller().Msg("Message failed to send")
			}
		}

	case "!pokemon ":
		//	silly command for funsies
		_, err := s.ChannelMessageSend(m.ChannelID, storage.ReturnRandomPokemon())
		if err != nil {
			log.Logger.Warn().Caller().Msg("Message failed to send")
		}

	case "!roll ":
		//	rolls command[1] amount of dice
		if b, err := strconv.Atoi(command[1]); err == nil {
			var a = rand.Intn(6 * b)
			for ok := true; ok; ok = (a < b) {
				a = rand.Intn(6 * b)
				fmt.Println("rerolling die...")
			}
			_, err := s.ChannelMessageSend(m.ChannelID, "You rolled "+strconv.Itoa(b)+" dice. \nThe result was: "+strconv.Itoa(a))
			if err != nil {
				log.Logger.Warn().Caller().Msg("Message failed to send")
			}
		} else {
			_, err := s.ChannelMessageSend(m.ChannelID, "You need to supply the number of dice to roll.\nFor example, for three dice: !roll 3")
			if err != nil {
				log.Logger.Warn().Caller().Msg("Message failed to send")
			}
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
			log.Logger.Warn().Caller().Msg("Failed to read image file")
			return
		}

		//	send the contents to the pastebin function to be pasted
		p := storage.Pastebin{}
		link, err := p.Put(string(file), "Ascii Image")
		if err != nil {
			_, err2 := s.ChannelMessageSend(m.ChannelID, "The image failed to convert! Let my owner know!")
			if err2 != nil {
				log.Logger.Warn().Caller().Msg("Message failed to send")
			}
			return
		} else {
			_, err = s.ChannelMessageSend(m.ChannelID, "Your image has been pasted at: "+link)
			if err != nil {
				log.Logger.Warn().Caller().Msg("Message failed to send")
			}
		}

	}
}
