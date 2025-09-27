package telegram

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Команды: /merge - слияние файлов, /watermark - ")
			_, err := bot.Send(msg)
			if err != nil {
				return
			}
		}

		if update.Message.Text == "/merge" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте файлы в формате PDF")
			_, err := bot.Send(msg)
			if err != nil {
				return
			}
		}
	}
}
