package telegram

import (
	"context"
	"fmt"
	"ilovepdf/internal/redis"
	"log"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MergeWaitingState = "merge:waiting"

func StartBot(ctx context.Context, token string) {
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

		chatIdStr := strconv.FormatInt(update.Message.Chat.ID, 10)

		//state, _ := redis.RedisClient.Get(ctx, chatIdStr).Result()
		//
		//if state == MergeWaitingState {
		//
		//}

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

			stateKey := fmt.Sprintf("user:%s:state", chatIdStr)

			err = redis.RedisClient.Set(ctx, stateKey, MergeWaitingState, 10*time.Minute).Err()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
