package main

import (
	"fmt"
	"ilovepdf/internal/storage"
	"ilovepdf/router"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	//}

	storage.InitDB()

	r := router.GetRouter()
	err := r.Run(":" + "8080")
	if err != nil {
		fmt.Println("Ошибка роутинга: ", err)
		return
	}
}
