package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"tg-worthlisteningto/pkg/client"
	"tg-worthlisteningto/pkg/helper"
)

func BotInit(token string) (tgbotapi.BotAPI, tgbotapi.UpdatesChannel, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var updateChannel = tgbotapi.NewUpdate(0)
	updateChannel.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateChannel)
	if err != nil {
		log.Panic(err)
	}

	return *bot, updates, nil
}

func Run(updates tgbotapi.UpdatesChannel, cl client.Azure) error {
	for update := range updates {
		pk, err := helper.GetGenre(update.ChannelPost.Text)
		if err != nil {
			return err
		}
		d := map[string]interface{}{"message": update.ChannelPost.Text}
		cl.AddData(pk, "0", d)
	}

	return nil
}
