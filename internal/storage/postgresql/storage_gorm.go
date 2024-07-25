package postgresql

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StorageGorm struct {
	DB *gorm.DB
}

func InitDB(cfg *config.Config) (*StorageGorm, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	log.Info("Successfully opened postgresql connection")

	return &StorageGorm{DB: db}, nil
}
