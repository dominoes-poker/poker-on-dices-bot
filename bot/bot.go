package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"github.com/and3rson/telemux/v2"
	"fmt"
	"poker-on-dices/config"
)

type Bot struct {
	*tgbotapi.BotAPI
	tgbotapi.UpdateConfig
}

func NewBot(config config.Config) *Bot {
	token := config.GetToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	return &Bot{
		bot,
		newUpdatesConfig(),
	}
}

func (bot *Bot) Dispatch(multiplexer *telemux.Mux) {
	updatesChannel := bot.GetUpdatesChan(bot.UpdateConfig)
	for update := range updatesChannel {
		fmt.Println(update, multiplexer, bot)
	}
}
