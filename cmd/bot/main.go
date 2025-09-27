// cmd/bot/main.go
package main

import (
	"ilovepdf/internal/storage"
	"ilovepdf/internal/telegram"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	//}

	storage.InitDB()
	telegram.StartBot("7145546881:AAGY-aHsoL6NDgSi1vYntEHFmrabCTkueY8")
}
