package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	PgDsn        string `env:"PG_DSN"`
	Port         int    `env:"PORT"`
	DoMigrations bool   `env:"DO_MIGRATIONS"`
	LogLevel     string `env:"LOG_LEVEL"`
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

func (c *Config) String() string {
	return fmt.Sprintf(`PG_DSN=%s, PORT=%d, DO_MIGRATIONS=%v, LOG_LEVEL=%s`, c.PgDsn, c.Port, c.DoMigrations, c.LogLevel)
}
