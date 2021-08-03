package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
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
	fmt.Println("reading config...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second)
		return err
	}

	//	unpack json
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second)
		return err
	}

	//	set global vars to be used
	Token = config.Token
	PastebinToken = config.PastebinToken
	BotPrefix = config.BotPrefix

	return nil
}
