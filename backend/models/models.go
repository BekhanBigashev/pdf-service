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
