package models

type Message struct {
	FromUserID int64 `gorm:"not null;index"`
	ToUserID   int64 `gorm:"not null;index"`
	Text       string
	FromUser   User `gorm:"foreignKey:FromUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToUser     User `gorm:"foreignKey:ToUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
