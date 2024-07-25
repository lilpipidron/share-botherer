package models

type Message struct {
	FromUserID uint `gorm:"not null;index"`
	ToUserID   uint `gorm:"not null;index"`
	Text       string
	FromUser   User `gorm:"foreignKey:FromUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToUser     User `gorm:"foreignKey:ToUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
