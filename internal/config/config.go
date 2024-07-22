package config

import "github.com/caarlos0/env/v11"

type Config struct {
	BotConfig
}

type BotConfig struct {
	Toket string `env:"BOT_TOKEN,required"`
}

func Load() *Config {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
