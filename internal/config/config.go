package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName            string        `envconfig:"app_name"`
	ServerPort         int           `envconfig:"server_port" default:"8080"`
	ServerReadTimeout  time.Duration `envconfig:"server_read_timeout" default:"5s"`
	ServerWriteTimeout time.Duration `envconfig:"server_write_timeout" default:"5s"`
	LogLevel           int           `envconfig:"log_level" default:"1"`
	TelemetryEndpoint  string        `envconfig:"telemetry_endpoint"`
}

func GetDefault() (*Config, error) {
	var cfg Config
	err := envconfig.Process("dotapickme", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to process env: %w", err)
	}
	return &cfg, nil
}
