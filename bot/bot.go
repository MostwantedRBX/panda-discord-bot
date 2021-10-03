package bot

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/mostwantedrbx/discord-go/commands"
	"github.com/mostwantedrbx/discord-go/config"
	"github.com/mostwantedrbx/discord-go/crypt"
)

//	init some variables
var botID string

// type VoiceUsers struct {
// 	GuildID   string
// 	ChannelID string
// 	UserID    string
// }

// var ChannelUsers []VoiceUsers

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

	goBot.AddHandler(messageHandler)       // function to fire when a message is posted
	goBot.AddHandler(channelUpdateHandler) // function to fire when a channel is joined
	//	goBot.AddHandler(channelUpdateJoinHandler)
	err = goBot.Open()
	if err != nil {
		log.Logger.Fatal().Msg("Bot could not add a message handler")
		return
	}
	log.Logger.Info().Msg("Bot is now running")
}

func channelUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	// get array of channel structs

	// vu := VoiceUsers{
	// 	GuildID:   m.GuildID,
	// 	ChannelID: m.ChannelID,
	// 	UserID:    m.UserID,
	// }
	// ChannelUsers = append(ChannelUsers, vu)
	// fmt.Println("added" + ChannelUsers[0].ChannelID)

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
	//	command splits the message received into the command, on command[0] and the arguments on command[1]
	switch command := strings.SplitAfter(cont, " "); command[0] {

	case "!help":
		//	sends an embed message with commands
		err := commands.Help(s, m)
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Help command could not go through.")
		}

	case "!ping":
		//	used to test if the bot is on
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Ping command could not go through.")
		}

	case "!echo ":
		//	responds to the user with the same thing they told the bot
		if len(command) > 1 {
			err := commands.Echo(s, m, command[1])
			if err != nil {
				log.Logger.Warn().Err(err).Msg("Echo command could not go through.")
			}
		}

	case "!pokemon ":
		//	silly command for funsies
		err := commands.Pokemon(s, m)
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Pokemon command could not go through.")
		}

	case "!roll ":
		//	rolls command[1] amount of dice
		if len(command) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Please give an integer as an argument after !roll, I.E.: !roll 4")
		}

		err := commands.Roll(s, m, command[1])
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Roll command could not go through.")
		}

	case "!convert ":
		//	converts and image link to an ascii char image then posts to pastebin
		if len(command) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Please give a direct URL to an image as an argument after !convert, I.E.: !convert URL_HERE")
		}

		err := commands.Convert(s, m, command[1])
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Convert command could not go through.")
		}

	case "!imbored":
		//	gives an activity idea to the user
		err := commands.Bored(s, m)
		if err != nil {
			log.Logger.Warn().Err(err).Msg("ImBored command could not go through.")
		}
	case "!coins ":
		if len(command) > 2 && command[1] == "remindme " {

			realCoin, err := crypt.ConfirmCoin(strings.TrimSpace(command[2]))
			if err != nil {
				log.Logger.Warn().Err(err).Msg("Could not get coin due to an error.")
			}

			if !realCoin {

				_, err := s.ChannelMessageSend(m.ChannelID, "Please give a real crypto currency name after !coins remindme I.E.: !coins remindme bitcoin 5 minutes")
				if err != nil {
					log.Logger.Warn().Err(err).Caller().Msg("Could not send chat message")
				}

				return
			}

			go func() {

				amount, err := strconv.Atoi(strings.TrimSpace(command[3]))
				if err != nil {
					log.Logger.Warn().Err(err).Msg("Could convert string")
				}

				switch command[4] {
				case "minutes", "minute":

					_, err := s.ChannelMessageSend(m.ChannelID, "I will remind you in "+command[3]+" "+command[4]+", about the price of "+command[2])
					if err != nil {
						log.Logger.Warn().Err(err).Msg("Could not send message!")
					}

					time.Sleep(time.Duration(amount) * time.Minute)
				case "hours", "hour":

					_, err := s.ChannelMessageSend(m.ChannelID, "I will remind you in "+command[3]+" "+command[4]+", about the price of "+command[2])
					if err != nil {
						log.Logger.Warn().Err(err).Msg("Could not send message!")
					}

					time.Sleep(time.Duration(amount) * time.Hour)
				case "days", "day":

					_, err := s.ChannelMessageSend(m.ChannelID, "I will remind you in "+command[3]+" "+command[4]+", about the price of "+command[2])
					if err != nil {
						log.Logger.Warn().Err(err).Msg("Could not send message!")
					}

					time.Sleep(time.Duration(amount) * (time.Hour * 24))

				default:
					_, err := s.ChannelMessageSend(m.ChannelID, "That was in invalid length of time. Try !coins remindme bitcoin 2 hours")

					log.Logger.Info().Msg(command[0] + ":" + command[1] + ":" + command[2] + ":" + command[3] + ":" + command[4]) //debug

					if err != nil {
						log.Logger.Warn().Err(err).Msg("Could not send message!")
					}
				}

				err = commands.Coins(s, m, strings.ReplaceAll(command[2], " ", ""))
				if err != nil {
					log.Logger.Warn().Err(err).Msg("Coins command failed.")
				}

			}()
			return
		}

		if len(command) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Please give a cryptocurrency name as an argument after !coins, I.E.: !coins bitcoin")
			return
		}

		err := commands.Coins(s, m, strings.ToLower(command[1]))
		if err != nil {
			log.Logger.Warn().Err(err).Msg("Coins command failed.")
		}
	}

}
