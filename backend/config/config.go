package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/urlerrors"
)

// Config — is an initial configuration params for deploying backend.
type Config struct {
	APPPort string `env:"APP_PORT" envDefault:"1323"`
	// Redis params
	REDISHost string `env:"REDIS_HOST" envDefault:"localhost"`
	REDISPort string `env:"REDIS_PORT" envDefault:"6379"`
	// ClickHouse params
	CHUser     string `env:"CLICKHOUSE_USER" envDefault:"oleg"`
	CHPassword string `env:"CLICKHOUSE_PASSWORD" envDefault:"tinkoff"`
	CHDBName   string `env:"CLICKHOUSE_DB" envDefault:"tbank_academy"`
	CHHost     string `env:"CLICKHOUSE_HOST" envDefault:"localhost"`
	CHPort     string `env:"CLICKHOUSE_PORT" envDefault:"19000"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, urlerrors.Error{
			Err:  urlerrors.ErrCannotInitConfig,
			Desc: err.Error(),
		}
	}
	return cfg, nil
}

// GetAppAddress  — builds addres that is used when starting server.
func (cfg *Config) GetAppAddress() string {
	return fmt.Sprintf(":%s", cfg.APPPort)
}

// GetRedisDSN — builds connection string for Redis that used in repository.New().
func (cfg *Config) GetRedisDSN() string {
	return fmt.Sprintf("%s:%s", cfg.REDISHost, cfg.REDISPort)
}

// GetClickHouseDSN — builds connection string for ClickHouse that used in repository.New().
func (cfg *Config) GetClickHouseDSN() string {
	return fmt.Sprintf("%s:%s", cfg.CHHost, cfg.CHPort)
}
