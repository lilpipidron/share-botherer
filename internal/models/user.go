package models

type User struct {
	ID         uint
	ChatID     uint `gorm:"unique;not null;index"`
	TelegramID uint `gorm:"unique;not null;index"`
}
