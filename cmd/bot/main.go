// cmd/bot/main.go
package main

import (
	"context"
	"log"
	"os"
	"pdf-service/internal/redis"
	"pdf-service/internal/storage"
	"pdf-service/internal/telegram"

	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	storage.InitDB()
	redis.InitRedis()

	telegram.StartBot(ctx, os.Getenv("TELEGRAM_BOT_TOKEN"))
}
