package main

import (
	"fmt"
	"time"

	"github.com/mostwantedrbx/discord-go/bot"
	"github.com/mostwantedrbx/discord-go/config"
)

func main() {

	//	read the config
	err := config.ReadConfig()

	//	catch error if needed
	if err != nil {
		fmt.Println("Couldn't connect to discord... maybe try again later? (ツ) \n", err.Error())
		time.Sleep(time.Second)
		return
	}

	//	start bot
	bot.Start()

	//	wait for commands
	<-make(chan struct{})

}
