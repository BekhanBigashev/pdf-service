package storage

import (
	"fmt"
	"ilovepdf/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	err = godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	fmt.Println("✅ Подключение к PostgreSQL через GORM успешно!")

	// миграция (создаст таблицы по структурам)
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
