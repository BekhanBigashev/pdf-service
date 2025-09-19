package main

import (
	"fmt"
	"ilovepdf/internal/storage"
	"ilovepdf/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	storage.InitDB()

	r := router.GetRouter()
	err = r.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		fmt.Println("Ошибка роутинга: ", err)
		return
	}
}
