package middlewares

import (
	"ilovepdf/handlers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ApiErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Внутренняя ошибка сервера",
					"message": "Произошла непредвиденная ошибка",
				})
			}
		}()
		c.Next()
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Забираем токен из заголовка
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Проверяем токен
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "i1nvalid token"})
			return
		}

		// Сохраняем user_id в контексте Gin
		c.Set("user_id", claims.Subject)

		c.Next()
	}
}
