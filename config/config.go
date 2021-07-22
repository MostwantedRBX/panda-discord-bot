package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// function is used to get token and other config settings at the start of the bot
func ReadConfig() error {
	fmt.Println("reading config...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println(string(file))
	//unpack json
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println(config.Token)
	//set global vars to be used
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
