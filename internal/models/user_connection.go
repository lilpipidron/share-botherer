package models

type UserConnection struct {
	ID int64 
  UserID1 int64 `gorm:"index"`
	UserID2 int64 `gorm:"index"`
  Paired bool
	User1   User `gorm:"foreignKey:UserID1;references:TelegramID;constraint:OnDelete:CASCADE;"`
	User2   User `gorm:"foreignKey:UserID2;references:TelegramID;constraint:OnDelete:CASCADE;"`
}
