package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// General
	Prod bool `envconfig:"PROD"`

	// HTTP server
	ServerHost      string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	ServerPort      int    `envconfig:"SERVER_PORT" default:"13337"`
	ServerTLSDomain string `envconfig:"SERVER_TLS_DOMAIN"`

	// Authentication
	AuthToken string `envconfig:"AUTH_TOKEN" required:"true"`

	// Static
	Static bool `envconfig:"STATIC"`

	// EdgeDB (deprecated)
	EdgeDBDSN string `envconfig:"EDGEDB_DSN"`

	// Directus
	DirectusURL   string `envconfig:"DIRECTUS_URL"`
	DirectusToken string `envconfig:"DIRECTUS_TOKEN"`

}

func Parse() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
