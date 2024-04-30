package models

import (
	"metroid_bookmarks/pkg/misc"
)

type Config struct {
	Db         Db     `json:"db"`
	Redis      Redis  `json:"redis"`
	PhotosPath string `json:"photos_path"`
}

type Db struct {
	Dsn string `json:"dsn"`
}
type Redis struct {
	Dsn string `json:"dsn"`
}

func NewConfig(dbConfig string) (*Config, error) {
	return misc.JsonToStruct[Config](dbConfig)
}
