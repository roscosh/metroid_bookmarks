package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func ParseEnv(envStruct interface{}) error {
	var errMessage string

	t := reflect.TypeOf(envStruct)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return &Error{message: "object must be a pointer of structure"}
	}

	elem := reflect.ValueOf(envStruct).Elem()
	for i := range elem.NumField() {
		structField := elem.Type().Field(i)
		envTag, exist := structField.Tag.Lookup("env")

		if !exist {
			errMessage = fmt.Sprintf("struct field '%s' have no env tag", structField.Name)
			return &Error{message: errMessage}
		}

		envValue, exist := os.LookupEnv(envTag)
		if !exist {
			errMessage = fmt.Sprintf("env key '%s' doesn't exist", envTag)
			return &Error{message: errMessage}
		}

		fieldValue := elem.Field(i)
		if !fieldValue.CanSet() {
			errMessage = fmt.Sprintf("'%s' can not set", structField.Name)
			return &Error{message: errMessage}
		}

		switch fieldValue.Type().Kind() {
		case reflect.Bool:
			val, err := strconv.ParseBool(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not boolean type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetBool(val)
		case reflect.Int:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not integer type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetInt(int64(val))
		case reflect.Int8:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not int8 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetInt(int64(val))
		case reflect.Int16:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not int16 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetInt(int64(val))
		case reflect.Int32:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not int32 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetInt(int64(val))
		case reflect.Int64:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not int64 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetInt(int64(val))
		case reflect.Uint:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not unsigned integer type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.Uint8:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not uint8 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.Uint16:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not uint16 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.Uint32:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not uint32 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.Uint64:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not uint64 type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.Uintptr:
			val, err := strconv.Atoi(envValue)
			if err != nil {
				errMessage = fmt.Sprintf("key '%s' is not uintptr type", envValue)
				return &Error{message: errMessage}
			}

			fieldValue.SetUint(uint64(val))
		case reflect.String:
			fieldValue.SetString(envValue)
		default:
			errMessage = fmt.Sprintf("key '%s' have unhandled type", envValue)
			return &Error{message: errMessage}
		}
	}

	return nil
}
