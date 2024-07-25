package config

import "github.com/caarlos0/env/v11"

type Config struct {
	BotConfig
  PostgresConfig
}

type PostgresConfig struct {
	PostgresUser    string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	PostgresDB      string `env:"POSTGRES_DB,required"`
	PostgresHost    string `env:"POSTGRES_HOST,required"`
	PostgresPort    int    `env:"POSTGRES_PORT,required"`
}

type BotConfig struct {
	Toket             string `env:"BOT_TOKEN,required"`
	WorkersForUpdates int    `env:"WORKERS_FOR_UPDATES" envDefault:"1"`
}

func Load() *Config {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
