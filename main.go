package main

import (
	"fmt"

	"github.com/mostwantedrbx/discord-go/bot"
	"github.com/mostwantedrbx/discord-go/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})

}
