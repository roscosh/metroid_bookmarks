package models

import (
	"errors"
	"io"
	"metroid_bookmarks/pkg/misc"
	"os"
)

var (
	ErrFileDoesNotExist   = errors.New("no file in path")
	ErrFailingConvertFile = errors.New("failing convert file to bytes")
)

type Redis struct {
	Dsn string `json:"dsn"`
}
type AppConfig struct {
	Redis      Redis  `json:"redis"`
	PhotosPath string `json:"photos_path"`
}

func NewAppConfig(appConfigPath string) (*AppConfig, error) {
	jsonFile, err := os.Open(appConfigPath)
	if err != nil {
		return nil, ErrFileDoesNotExist
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, ErrFailingConvertFile
	}

	return misc.JSONToStruct[AppConfig](bytes)
}
