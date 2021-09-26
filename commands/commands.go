package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/mostwantedrbx/discord-go/crypt"
	"github.com/mostwantedrbx/discord-go/net"
	"github.com/mostwantedrbx/discord-go/pyscripts"
	"github.com/mostwantedrbx/discord-go/storage"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) error {
	//	when someone calls !help the bot responds with these embed messages.

	//	they needed to be split since discord wont take one massive pile of data-
	//	for an embed message
	helpField := []*discordgo.MessageEmbedField{
		{
			Name: "-", Value: "Random/Basic commands",
		},
		{
			Name: "!help", Value: "Brings up this message.",
		},
		{
			Name: "!ping", Value: "Pong!",
		},
		{
			Name: "!echo message", Value: "Makes the bot say message.",
		},
		{
			Name: "!roll numOfDice", Value: "Rolls numOfDice 6 sided dice.",
		},
		{
			Name: "!pokemon", Value: "Says a random pokemon related quote.",
		},
		{
			Name: "!imbored", Value: "Returns a random activity idea from boredapi.com",
		},
	}
	helpField2 := []*discordgo.MessageEmbedField{
		{
			Name: "-", Value: "Cool python scripts to try out!",
		},
		{
			Name: "!convert directImageURL", Value: "Converts directImageURL into an ascii image and returns a link to pastebin!",
		},
	}

	eMes := discordgo.MessageEmbed{
		Title:  "Commands!",
		Color:  3066993,
		Fields: helpField,
	}
	eMes2 := discordgo.MessageEmbed{
		Title:  "Python Scripts",
		Color:  3066993,
		Fields: helpField2,
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &eMes)
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &eMes2)
	if err != nil {
		return err
	}
	return nil
}

func Roll(s *discordgo.Session, m *discordgo.MessageCreate, command string) error {
	//	rolls {command} 6 sided die and sends it back
	if b, err := strconv.Atoi(command); err == nil {
		var a = rand.Intn(6 * b)
		for ok := true; ok; ok = (a < b) {
			a = rand.Intn(6 * b)
			fmt.Println("rerolling die...")
		}
		_, err := s.ChannelMessageSend(m.ChannelID, "You rolled "+strconv.Itoa(b)+" dice. \nThe result was: "+strconv.Itoa(a))
		if err != nil {
			return err
		}
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "You need to supply the number of dice to roll.\nFor example, for three dice: !roll 3")
		if err != nil {
			return err
		}
	}
	return nil
}

func Convert(s *discordgo.Session, m *discordgo.MessageCreate, command string) error {
	// download the file from the url
	net.DownloadFile(command, "tacos.png")

	//	run the python script to convert the image, and it saves it in a txt file.
	pyscripts.RunScript("convert")

	//	gets the results
	fmt.Println("Opening a file ")
	file, err := ioutil.ReadFile("./ascii-image.txt")
	if err != nil {
		return err
	}

	//	send the contents to the pastebin function to be pasted
	p := storage.Pastebin{}
	link, err := p.Put(string(file), "Ascii Image")
	if err != nil {
		_, err2 := s.ChannelMessageSend(m.ChannelID, "The image failed to convert! Let my owner know!")
		if err2 != nil {
			return err
		}
		return err
	} else {
		_, err = s.ChannelMessageSend(m.ChannelID, "Your image has been pasted at: "+link)
		if err != nil {
			return err
		}
	}
	return nil
}

func Pokemon(s *discordgo.Session, m *discordgo.MessageCreate) error {
	//	simply sends a message containing a pokemon quote
	_, err := s.ChannelMessageSend(m.ChannelID, storage.ReturnRandomPokemon())
	if err != nil {
		return err
	}
	return nil
}

func Echo(s *discordgo.Session, m *discordgo.MessageCreate, command string) error {
	//	repeats {command} back to the person that said it
	_, err := s.ChannelMessageSend(m.ChannelID, command)
	if err != nil {
		return err
	}
	return nil
}

func Bored(s *discordgo.Session, m *discordgo.MessageCreate) error {
	//	fetch an activity from the boredapi.com api, unmarshal json, then send the activity to the user that requested it
	res, err := http.Get("https://www.boredapi.com/api/activity?participants=1")
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "Could not reach the API endpoint")
		if err != nil {
			return err
		}
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	byteBody := []byte(body)
	var inbound storage.Bored

	err = json.Unmarshal(byteBody, &inbound)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, inbound.Activity+", "+m.Author.Username+".")
	if err != nil {
		return err
	}
	return nil
}

func Coins(s *discordgo.Session, m *discordgo.MessageCreate, coin string) error {
	var coinStats crypt.CoinData
	coinStats, err := crypt.FetchCoinData(coin)
	if err != nil {
		return err
	}
	coinField := []*discordgo.MessageEmbedField{
		{
			Name: "ID:", Value: coinStats.ID,
		},
		{
			Name: "Rank:", Value: strconv.Itoa(coinStats.Rank),
		},
		{
			Name: "Last Update:", Value: coinStats.LastUpdate,
		},
		{
			Name: "Price/(1 Coin):", Value: "$" + fmt.Sprintf("%.2f", coinStats.Quotes.USD.Price),
		},
		{
			Name: "Change last 30m:", Value: fmt.Sprintf("%.2f", coinStats.Quotes.USD.Last30) + "%",
		},
		{
			Name: "Change last 24h:", Value: fmt.Sprintf("%.2f", coinStats.Quotes.USD.LastDay) + "%",
		},
		{
			Name: "Change last 7d:", Value: fmt.Sprintf("%.2f", coinStats.Quotes.USD.LastWeek) + "%",
		},
	}

	eMes := discordgo.MessageEmbed{
		Title:  coinStats.Name,
		Color:  15844367, //gold
		Fields: coinField,
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &eMes)
	if err != nil {
		return err
	}

	return nil
}
