package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

var (
	Token         string
	PastebinToken string
	BotPrefix     string

	config *configStruct
)

type configStruct struct {
	Token         string `json:"DiscordBotToken"`
	PastebinToken string `json:"PastebinToken"`
	BotPrefix     string `json:"BotPrefix"`
}

//	function is used to get token and other config settings at the start of the bot
func ReadConfig() error {
	log.Logger.Info().Msg("Reading ./config.json")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Logger.Warn().Msg("Could not read config file. Trying to make default. Please add your tokens/bot prefix to config.json")
		err = createConfig()
		if err != nil {
			log.Logger.Fatal().Msg("Creation of config.json failed. Maybe the program doesn't have permission to write?")
			return err
		}
		return nil
	}

	//	unpack json
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Logger.Fatal().Msg("Could not unpack json file.")
		return err
	}

	//	set global vars to be used
	Token = config.Token
	PastebinToken = config.PastebinToken
	BotPrefix = config.BotPrefix

	return nil
}

func createConfig() error {
	conf := configStruct{
		Token:         "",
		PastebinToken: "",
		BotPrefix:     "!",
	}
	defaultConfig, err := json.MarshalIndent(conf, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./config.json", defaultConfig, 0644)
	if err != nil {
		return err
	}
	return nil
}
