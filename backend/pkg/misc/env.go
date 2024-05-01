package misc

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func ParseEnv(envStruct interface{}) error {
	var errMessage string

	t := reflect.TypeOf(envStruct)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return errors.New("object must be a pointer of structure")
	}

	elem := reflect.ValueOf(envStruct).Elem()
	for i := 0; i < elem.NumField(); i++ {
		structField := elem.Type().Field(i)
		envTag, exist := structField.Tag.Lookup("env")
		if !exist {
			errMessage = fmt.Sprintf("%s have no env tag", structField.Name)
			return errors.New(errMessage)
		}
		envValue, exist := os.LookupEnv(envTag)
		if !exist {
			errMessage = fmt.Sprintf("%s don't exist", envTag)
			return errors.New(errMessage)
		}
		fieldValue := elem.Field(i)
		if !fieldValue.CanSet() {
			errMessage = fmt.Sprintf("%s can not set", structField.Name)
			return errors.New(errMessage)
		}
		switch fieldValue.Type().Kind() {
		case reflect.Bool:
			val, err := strconv.ParseBool(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not boolean type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetBool(val)
		case reflect.Int:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.Int8:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.Int16:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.Int32:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.Int64:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", envValue)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.String:
			fieldValue.SetString(envValue)
		default:
			errMessage = fmt.Sprintf("%s - unhandled type", envValue)
			return errors.New(errMessage)
		}
	}
	return nil
}
