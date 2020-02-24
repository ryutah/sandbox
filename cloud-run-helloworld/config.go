package helloworld

import (
	"github.com/caarlos0/env/v6"
)

var cfg Config

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
}

type Config struct {
	ProjectID string `env:"GOOGLE_CLOUD_PROJECT"`
}

func Configuration() Config {
	return cfg
}
