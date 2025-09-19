package handlers

import (
	"fmt"
	"ilovepdf/internal/pdf"
	"ilovepdf/internal/storage"
	"ilovepdf/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register Регистрация
func Register(c *gin.Context) {
	var input models.AuthForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хэшируем пароль
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	// Сохраняем в БД (псевдо)
	user := models.User{Email: input.Email, Password: string(hashedPassword)}
	storage.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

// JWT секрет
var JwtKey = []byte("my_secret_key")

func GenerateTokens(userID string) (string, string, error) {
	accessExp := time.Now().Add(15 * time.Minute)
	refreshExp := time.Now().Add(7 * 24 * time.Hour)

	accessClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(accessExp),
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	refreshClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshExp),
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	at, err := accessToken.SignedString(JwtKey)
	if err != nil {
		return "", "", err
	}
	rt, err := refreshToken.SignedString(JwtKey)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

func Login(c *gin.Context) {
	var input models.AuthForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем пользователя в БД
	var user models.User
	if err := storage.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email"})
		return
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	access, refresh, err := GenerateTokens(fmt.Sprint(user.ID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// refresh лучше положить в HttpOnly cookie, а не отдавать в JSON
	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

// HandleMergePDF Соединение pdf файлов
func HandleMergePDF(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	merged, err := pdf.MergeFiles(files, c)
	if err != nil {
		return
	}

	c.File(merged)
	err = os.Remove(merged)
	if err != nil {
		return
	}

	return
}

// HandleWatermark Добавление водяной марки
func HandleWatermark(c *gin.Context) {
	form, _ := c.MultipartForm()
	file := form.File["file"][0]

	watermarked, err := pdf.AddWaterMark(file, c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.File(watermarked)

	return
}

// GET /profile
func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, user " + userID.(string),
	})
}
