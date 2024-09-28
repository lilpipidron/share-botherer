package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/charmbracelet/log"
)

type Config struct {
	BotConfig
	PostgresConfig
}

type PostgresConfig struct {
	Empty            bool
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
}

type BotConfig struct {
	Token             string `env:"BOT_TOKEN,required"`
	WorkersForUpdates int    `env:"WORKERS_FOR_UPDATES" envDefault:"1"`
}

func Load() *Config {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	if config.PostgresUser == "" ||
		config.PostgresPassword == "" ||
		config.PostgresDB == "" ||
		config.PostgresHost == "" ||
		config.PostgresPort == 0 {
		config.PostgresConfig = PostgresConfig{}
		config.PostgresConfig.Empty = true
	}

	return &config
}
