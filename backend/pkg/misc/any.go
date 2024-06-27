package misc

import (
	"encoding/json"
	"errors"
	"reflect"
)

var ErrJSONToStruct = errors.New("failing convert json to struct")

func Contains[T comparable](item T, arr []T) bool {
	for _, value := range arr {
		if value == item {
			return true
		}
	}

	return false
}

func JSONToStruct[T any](bytes []byte) (*T, error) {
	var config *T

	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, ErrJSONToStruct
	}

	return config, nil
}

func GetTags[T any](tagName string) []string {
	var (
		structObj      T
		dbTagArray     []string
		traverseFields func(reflect.Type)
	)

	structType := reflect.TypeOf(structObj)
	traverseFields = func(t reflect.Type) {
		for i := range t.NumField() {
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
