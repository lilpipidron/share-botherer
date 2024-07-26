package models

type User struct {
	ID         int
	ChatID     int64  `gorm:"unique;not null;index"`
	TelegramID int64  `gorm:"unique;not null;index"`
	Username   string `gorm:"not null;index"`
}
