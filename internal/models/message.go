package models

type Message struct {
	ID         int64
	FromUserID int64  `gorm:"not null;index"`
	ToUserID   int64  `gorm:"not null;index"`
	DeleteKey  string `gorm:"not null"`
	Text       string
	FromUser   User `gorm:"foreignKey:FromUserID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToUser     User `gorm:"foreignKey:ToUserID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
