package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Prod       bool   `envconfig:"PROD"`
	EdgeDBDSN  string `envconfig:"EDGEDB_DSN" required:"true"`
	ServerHost string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	ServerPort int    `envconfig:"SERVER_PORT" default:"13337"`
	AuthToken  string `envconfig:"AUTH_TOKEN" required:"true"`
}

func Parse() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
