package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"pdf-service/internal/storage"
	"pdf-service/router"
)

func main() {
	err := godotenv.Load("/app/.env")
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
