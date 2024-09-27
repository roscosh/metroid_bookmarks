package models

import (
	"metroid_bookmarks/pkg/misc/env"
)

type EnvConfig struct {
	AppConfigPath     string `env:"APP_CONFIG_PATH"`
	LogLevel          string `env:"LOG_LEVEL"`
	MaxConnLifetime   int64  `env:"MAX_CONN_LIFE_TIME"`
	MaxConnIdleTime   int64  `env:"MAX_CONNIDLE_TIME"`
	HealthCheckPeriod int64  `env:"HEALTH_CHECK_PERIOD"`
	MaxConns          int32  `env:"MAX_CONNS"`
	MinConns          int32  `env:"MIN_CONNS"`
	Production        bool   `env:"PRODUCTION"`
	DatabaseURL       string `env:"DATABASE_URL"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var conf EnvConfig

	err := env.ParseEnv(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
