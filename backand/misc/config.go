package misc

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Db    Db    `json:"db"`
	Redis Redis `json:"redis"`
	Debug bool  `json:"debug"`
}

type Db struct {
	Dsn string `json:"dsn"`
}
type Redis struct {
	Dsn string `json:"dsn"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		//todo брать из переменной окружения
		jsonFile, err := os.Open("etc/config/config.json")
		if err != nil {
			panic("No config file!")
		}
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				panic("Failing close config file!")
			}
		}(jsonFile)
		bytes, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic("Failing convert config file to bytes!")
		}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic("Failing convert config file to struct 'Config'!")
		}
	})
	return config
}
