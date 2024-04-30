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
		return errors.New("object must be a structure")
	}

	v := reflect.ValueOf(envStruct).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		envTag, exist := field.Tag.Lookup("env")
		if !exist {
			errMessage = fmt.Sprintf("%s have no env tag", field.Name)
			return errors.New(errMessage)
		}
		value, exist := os.LookupEnv(envTag)
		if !exist {
			errMessage = fmt.Sprintf("%s don't exist", envTag)
			return errors.New(errMessage)
		}
		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			errMessage = fmt.Sprintf("%s can not set", field.Name)
			return errors.New(errMessage)
		}
		switch fieldValue.Type().Kind() {
		case reflect.Bool:
			val, err := strconv.ParseBool(value)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not boolean type", value)
				return errors.New(errMessage)
			}
			fieldValue.SetBool(val)
		case reflect.Int:
			val, err := strconv.Atoi(value)
			if err != nil {
				errMessage = fmt.Sprintf("%s is not integer type", value)
				return errors.New(errMessage)
			}
			fieldValue.SetInt(int64(val))
		case reflect.String:
			fieldValue.SetString(value)
		default:
			errMessage = fmt.Sprintf("%s - unhandled type of field", value)
			return errors.New(errMessage)
		}
	}
	return nil
}
