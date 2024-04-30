package models

import "metroid_bookmarks/pkg/misc"

type EnvConfig struct {
	Production          bool   `env:"PRODUCTION"`
	DbConfig            string `env:"DB_CONFIG"`
	DbmateMigrationsDir string `env:"DBMATE_MIGRATIONS_DIR"`
	LogLevel            string `env:"LOG_LEVEL"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var conf EnvConfig
	err := misc.ParseEnv(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
