package config

import (
	"sync"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	PgDsn string `env:"PG_DSN"`
	Port  int    `env:"PORT"`
}

var (
	config Config
	once   sync.Once
	err    error
)

func GetDefault() (*Config, error) {
	once.Do(func() {
		err = env.Parse(&config)
	})

	return &config, err
}
