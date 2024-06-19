package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/urlerrors"
)

// Config — is an initial configuration params for deploying backend.
type Config struct {
	APPPort string `env:"APP_PORT" envDefault:"1323"`
	// TODO: Redis
	REDISHost string `env:"REDIS_HOST" envDefault:"localhost"`
	REDISPort string `env:"REDIS_PORT" envDefault:"6379"`
	// TODO: ClickHouse
	CHUser     string `env:"CH_USER" envDefault:"oleg"`
	CHPassword string `env:"CH_PASSWORD" envDefault:"tinkoff"`
	CHDBName   string `env:"CH_DB" envDefault:"tbank_academy"`
	CHHost     string `env:"CH_HOST" envDefault:"localhost"`
	CHPort     string `env:"CH_PORT" envDefault:"19000"`
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

// GetAppAddress  — ...
func (cfg *Config) GetAppAddress() string {
	return fmt.Sprintf(":%s", cfg.APPPort)
}

// GetRedisDSN — ...
func (cfg *Config) GetRedisDSN() string {
	return fmt.Sprintf("%s:%s", cfg.REDISHost, cfg.REDISPort)
}

// GetClickHouseDSN — ...
func (cfg *Config) GetClickHouseDSN() string {
	return fmt.Sprintf("%s:%s", cfg.CHHost, cfg.CHPort)
}
