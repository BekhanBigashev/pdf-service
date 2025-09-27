package models

// Модель User
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}

type AuthForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Структура состояния
type BotUserState struct {
	Step string            `json:"step"`
	Data map[string]string `json:"data"`
}
