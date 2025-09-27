// cmd/bot/main.go
package main

import (
	"ilovepdf/internal/storage"
	"ilovepdf/internal/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	storage.InitDB()
	telegram.StartBot(os.Getenv("TELEGRAM_BOT_TOKEN"))
}
