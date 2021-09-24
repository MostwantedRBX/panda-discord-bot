package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/mostwantedrbx/discord-go/commands"
	"github.com/mostwantedrbx/discord-go/config"
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

	goBot.AddHandler(messageHandler)       // function to fire when a message is posted
	goBot.AddHandler(channelUpdateHandler) // function to fire when a channel is joined

	err = goBot.Open()
	if err != nil {
		log.Logger.Fatal().Msg("Bot could not add a message handler")
		return
	}

	log.Logger.Info().Msg("Bot is now running")
}

func channelUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	// get array of channel structs
	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		log.Logger.Warn().Caller().Msg("Could not get channels")
	}

	if m.ChannelID == "" {
		log.Logger.Info().Msg("Channel left. " + m.BeforeUpdate.ChannelID)
		for i := 0; i < len(channels); i++ {
			if channels[i].ID == m.BeforeUpdate.ChannelID && strings.Contains(channels[i].Name, "Team Room; 1 Users") {
				log.Logger.Debug().Msg("Deleting channel with ID: " + m.BeforeUpdate.ChannelID)
				_, err := s.ChannelDelete(m.BeforeUpdate.ChannelID)
				if err != nil {
					log.Logger.Warn().Msg("Could not delete channel with ID: " + m.BeforeUpdate.ChannelID)
				}
			}
		}
		return
	}

	log.Logger.Info().Msg("Channel Joined. ChannelID: " + m.ChannelID)

	for i := 0; i < len(channels); i++ {
		if channels[i].ID == m.ChannelID {
			if strings.Contains(channels[i].Name, "Create") && strings.Contains(channels[i].Name, "Channel") {
				fmt.Println("Um check passed?")
				// create channel with the same name
				numOfUsers := 5
				c, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
					Name:      "Team Room; 1 Users",
					Type:      2,
					ParentID:  channels[i].ParentID,
					UserLimit: numOfUsers, //TODO: Gonna make this more dynamic
				})
				if err != nil {
					log.Logger.Warn().Msg("Couldn't create channel\n" + err.Error())
					return
				}
				// move people to the new channel
				s.GuildMemberMove(m.GuildID, m.UserID, &c.ID)
			} else if strings.Contains(channels[i].Name, "Team Room; ") { //TODO: Figure out how to get the number of users in a channel at any given point
				for i := 1; i < 9; i++ {
					if strings.Contains(channels[i].Name, strconv.Itoa(i)) {
						s.ChannelEdit(m.ChannelID, "Team Room; "+strconv.Itoa(i)+" Users")
					}
				}
			}
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
	//	command splits the message received into the command, on command[0] and the arguments on command[1]
	switch command := strings.SplitAfter(cont, " "); command[0] {

	case "!help":
		//	sends an embed message with commands
		commands.Help(s, m)

	case "!ping":
		//	used to test if the bot is on
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			log.Logger.Warn().Caller().Msg("Message failed to send")
		}

	case "!echo ":
		//	responds to the user with the same thing they told the bot
		if len(command) > 1 {
			commands.Echo(s, m, command[1])
		}

	case "!pokemon ":
		//	silly command for funsies
		commands.Pokemon(s, m)

	case "!roll ":
		//	rolls command[1] amount of dice
		commands.Roll(s, m, command[1])

	case "!convert ":
		//	converts and image link to an ascii char image then posts to pastebin
		commands.Convert(s, m, command[1])
	case "!imbored":
		//	gives an activity idea to the user
		commands.Bored(s, m)
	}
}
