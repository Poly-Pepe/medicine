package config

import (
	"github.com/caarlos0/env/v9"
)

//todo: это конфиги для теста

//nolint:lll //envDefault.
type Config struct {
	AppPort     int    `env:"APP_PORT" envDefault:"8004"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	DataBaseDNS string `env:"DATABASE_DSN"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
