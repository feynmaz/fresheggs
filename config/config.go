package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port         int    `envconfig:"PORT" default:"8080"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	LogJson      bool   `envconfig:"LOG_JSON" default:"false"`
	PgDsn        string `envconfig:"PG_DSN" required:"true"`
	DoMigrations bool   `envconfig:"DO_MIGRATIONS" default:"false"`
}

var (
	cfg  config
	once sync.Once
)

func GetDefault() *config {
	once.Do(func() {
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatal(err)
		}
	})
	return &cfg
}

func (c config) String() string {
	str, _ := json.Marshal(c)
	return string(str)
}
