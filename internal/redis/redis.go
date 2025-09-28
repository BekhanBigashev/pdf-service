package redis

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})
}
