package models

type UserConnection struct {
	UserID1 uint `gorm:"index"`
	UserID2 uint `gorm:"index"`
	User1   User `gorm:"foreignKey:UserID1;references:ID;constraint:OnDelete:CASCADE;"`
	User2   User `gorm:"foreignKey:UserID2;references:ID;constraint:OnDelete:CASCADE;"`
}
