package storage

import (
	"fmt"
	"ilovepdf/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// создаст файл app.db в текущей директории
	//dsn := "host=go_db user=pdfuser password=pdfpassword dbname=pdfservice port=5432 sslmode=disable"
	dsn := "postgresql://root:root@db:5432/postgres"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
