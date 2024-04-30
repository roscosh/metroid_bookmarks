package misc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func Contains[T comparable](item T, arr []T) bool {
	for _, value := range arr {
		if value == item {
			return true
		}
	}
	return false
}

func JsonToStruct[T any](path string) (*T, error) {
	var config *T
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, errors.New("no file in path")
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			return
		}
	}(jsonFile)
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.New("failing convert file to bytes")
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, errors.New("failing convert file to struct")
	}
	return config, nil
}
