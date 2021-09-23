package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/mostwantedrbx/discord-go/net"
	"github.com/mostwantedrbx/discord-go/pyscripts"
	"github.com/mostwantedrbx/discord-go/storage"
)

func Roll(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	//	rolls {command} 6 sided die and sends it back
	if b, err := strconv.Atoi(command); err == nil {
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
}

func Convert(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	// download the file from the url
	net.DownloadFile(command, "tacos.png")

	//	run the python script to convert the image, and it saves it in a txt file.
	pyscripts.RunScript("convert")

	//	gets the results
	fmt.Println("Opening a file ")
	file, err := ioutil.ReadFile("./ascii-image.txt")
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

func Pokemon(s *discordgo.Session, m *discordgo.MessageCreate) {
	//	simply sends a message containing a pokemon quote
	_, err := s.ChannelMessageSend(m.ChannelID, storage.ReturnRandomPokemon())
	if err != nil {
		log.Logger.Warn().Caller().Msg("Message failed to send")
	}
}

func Echo(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	//	repeats {command} back to the person that said it
	_, err := s.ChannelMessageSend(m.ChannelID, command)
	if err != nil {
		log.Logger.Warn().Caller().Msg("Message failed to send")
	}
}

func Bored(s *discordgo.Session, m *discordgo.MessageCreate) {
	//	fetch an activity from the boredapi.com api, unmarshal json, then send the activity to the user that requested it
	res, err := http.Get("https://www.boredapi.com/api/activity?participants=1")
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "Could not reach the API endpoint")
		if err != nil {
			panic(err)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger.Err(err).Msg("Could not read data from res.Body")
	}

	byteBody := []byte(body)
	var inbound storage.Bored

	err = json.Unmarshal(byteBody, &inbound)
	if err != nil {
		log.Logger.Err(err).Msg("Could not unmarshal json")
	}

	_, err = s.ChannelMessageSend(m.ChannelID, inbound.Activity+", "+m.Author.Username+".")
	if err != nil {
		log.Logger.Err(err).Msg("Could not send response message")
	}
}
