package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	AppPort     int    `env:"APP_PORT" envDefault:"8004"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	DataBaseDNS string `env:"DATABASE_DSN" envDefault:"postgresql://myuser:mypassword@localhost:5444/main_db?sslmode=disable"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
