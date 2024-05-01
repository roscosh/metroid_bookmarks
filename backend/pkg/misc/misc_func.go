package misc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
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

func GetTags[T any](tagName string) []string {
	var structObj T
	structType := reflect.TypeOf(structObj)
	var dbTagArray []string
	var traverseFields func(reflect.Type)

	traverseFields = func(t reflect.Type) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			// Если поле встраивается из другой структуры
			if field.Anonymous {
				traverseFields(field.Type)
				continue
			}

			// Иначе получаем тэги и добавляем их к списку
			dbTag := field.Tag.Get(tagName)
			if dbTag != "" {
				dbTagArray = append(dbTagArray, dbTag)
			}

		}
	}

	traverseFields(structType)

	return dbTagArray
}
