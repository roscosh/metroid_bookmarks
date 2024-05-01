package models

import (
	"metroid_bookmarks/pkg/misc"
)

type PostgreSQL struct {
	Dsn string `json:"dsn"`
}
type Redis struct {
	Dsn string `json:"dsn"`
}
type AppConfig struct {
	PostgreSQL          PostgreSQL `json:"postgresql"`
	Redis               Redis      `json:"redis"`
	DbmateMigrationsDir string     `json:"dbmate_migrations_dir"`
	PhotosPath          string     `json:"photos_path"`
}

func NewAppConfig(dbConfig string) (*AppConfig, error) {
	return misc.JsonToStruct[AppConfig](dbConfig)
}
