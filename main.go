package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/mostwantedrbx/discord-go/bot"
	"github.com/mostwantedrbx/discord-go/config"
)

func main() {
	//	log setup
	//	log setup
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Logger = log.Output(io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, file))
	log.Logger.Info().Msg("Logs started")

	//	read the config
	err = config.ReadConfig()

	//	catch error if needed
	if err != nil {
		log.Logger.Fatal().Msg("Could not connect to Discord")
		return
	}

	//	start bot
	bot.Start()

	//	wait for commands
	<-make(chan struct{})

}
