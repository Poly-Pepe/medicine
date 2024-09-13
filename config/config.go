package config

import (
	"github.com/caarlos0/env/v9"
)

//nolint:lll //envDefault.
type Config struct {
	AppPort     int    `env:"APP_PORT" envDefault:"8004"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	DataBaseDNS string `env:"DATABASE_DSN" envDefault:"postgresql://architect:09970207@localhost:5444/medecine?sslmode=disable"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
