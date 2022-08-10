package main

import (
	"poker-on-dices/bot"
	"github.com/and3rson/telemux/v2"
	"poker-on-dices/config"
	data_service "poker-on-dices/data-service"
)

func main() {
	applicationConfig := config.NewConfig()

	_, err := data_service.NewDataService(applicationConfig)
	if err != nil {
		panic(err)
	}

	tgBot := bot.NewBot(applicationConfig)

	multiplexer := telemux.NewMux()

	tgBot.Dispatch(multiplexer)
}
