package postgresql

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/config"
	"github.com/lilpipidron/share-botherer/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StorageGorm struct {
	DB *gorm.DB
}

func InitDB(cfg *config.Config) *StorageGorm {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successfully opened postgresql connection")

	if err := db.AutoMigrate(&models.User{}, &models.Message{}, &models.UserConnection{}); err != nil {
		log.Fatal(err)
	}
	return &StorageGorm{DB: db}
}
