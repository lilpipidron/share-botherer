package storage

import "github.com/lilpipidron/share-botherer/internal/models"

type IStorage interface {
	SaveUser(user *models.User) error
	GetRandomMessages() ([]models.Message, error)
	DeleteMessage(text string, toUserID int64) error
	FindUserByUsername(username string) (*models.User, error)
	FindUserConnection(userID1, userID2 int64) (*models.UserConnection, error)
	SaveMessage(message *models.Message) error
	SaveUserConnection(connection *models.UserConnection) error
}
