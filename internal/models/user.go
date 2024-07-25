package models

type User struct {
	ID         int
	ChatID     int64 `gorm:"unique"`
	TelegramID int64 `gorm:"unique"`
}
