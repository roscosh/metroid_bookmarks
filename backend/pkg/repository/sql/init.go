package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"metroid_bookmarks/misc"
	"reflect"
	"strings"
)

var logger = misc.GetLogger()

type postgresPool struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func newPostgresPool(dsn string) (*postgresPool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &postgresPool{ctx: ctx, pool: pool}, nil
}

func getDbTags(structObj interface{}) string {
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
			dbTag := field.Tag.Get("db")
			dbTagArray = append(dbTagArray, dbTag)
		}
	}

	traverseFields(structType)

	return strings.Join(dbTagArray, ", ")
}

func update(table string, pk int, setInterface interface{}, returningInterface interface{}) (string, []interface{}, error) {
	// Получаем тип структуры
	userType := reflect.TypeOf(setInterface)
	// Получаем значение структуры
	userValue := reflect.ValueOf(setInterface)
	var values []interface{}
	var fields []string
	var placeholder = 1
	for i := 0; i < userType.NumField(); i++ {
		value := userValue.Field(i)
		if value.IsNil() {
			continue
		}
		values = append(values, value.Interface())
		// Получаем название поля
		fieldName := userType.Field(i).Tag.Get("json")
		// Добавляем позиционный индекс
		fieldStr := fmt.Sprintf("%s = $%v", fieldName, placeholder)
		fields = append(fields, fieldStr)
		placeholder++

	}
	if len(fields) == 0 {
		return "", nil, errors.New("empty setInterface")
	}
	set := strings.Join(fields, ", ")
	values = append(values, pk)

	var attributes []string
	typ := reflect.TypeOf(returningInterface)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		attributes = append(attributes, field.Tag.Get("json"))
	}
	returning := strings.Join(attributes, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s  WHERE id = $%v RETURNING %s", table, set, placeholder, returning)
	return query, values, nil
}

func total(table string) string {
	return fmt.Sprintf("SELECT COUNT(*) FROM %s ", table)
}
