package sqlite

import (
	"os"

	"github.com/lilpipidron/share-botherer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

const (
	dataDir    = "./_data"
	connection = "./_data/sqlite.db"
)

func NewStorage() (*Storage, error) {
	var err error

	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Message{},
		&models.UserConnection{},
	); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) SaveUser(user *models.User) error {
	return s.db.Save(user).Error
}

func (s *Storage) GetRandomMessages() ([]models.Message, error) {
	var messages []models.Message
	sql := `
SELECT *
FROM (
    SELECT *, ROW_NUMBER() OVER(PARTITION BY from_user_id, to_user_id ORDER BY RANDOM() DESC) AS rn
    FROM messages
) t
WHERE rn = 1
ORDER BY RANDOM() DESC;
`
	if err := s.db.Raw(sql).Scan(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *Storage) DeleteMessage(text string, toUserID int64) error {
	return s.db.Delete(
		models.Message{},
		"text = ? and to_user_id = ?",
		text,
		toUserID,
	).Error
}

func (s *Storage) FindUserByUsername(
	username string,
) (*models.User, error) {
	user := &models.User{}
	if err := s.db.First(
		user,
		"username = ?",
		username,
	).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Storage) FindUserConnection(
	userID1, userID2 int64,
) (*models.UserConnection, error) {
	pair := &models.UserConnection{}
	if err := s.db.First(
		pair,
		"user_id1 = ? and user_id2 = ? or user_id1 = ? and user_id2 = ?",
		userID1,
		userID2,
		userID2,
		userID1,
	).Error; err != nil {
		return nil, err
	}
	return pair, nil
}

func (s *Storage) SaveMessage(message *models.Message) error {
	return s.db.Save(message).Error
}

func (s *Storage) SaveUserConnection(connection *models.UserConnection) error {
	return s.db.Save(connection).Error
}
