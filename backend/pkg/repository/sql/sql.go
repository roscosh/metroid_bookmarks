package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"strings"
)

type baseSQL struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func newPostgresPool(dsn string) (*baseSQL, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &baseSQL{ctx: ctx, pool: pool}, nil
}

func (s *baseSQL) query(query string, args ...any) (pgx.Rows, error) {
	return s.pool.Query(s.ctx, query, args...)
}

func (s *baseSQL) queryRow(query string, args ...any) pgx.Row {
	return s.pool.QueryRow(s.ctx, query, args...)
}

func selectById[T any](baseSQL *baseSQL, table string, pk int) (*T, error) {
	var structObj T
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", getDbTags(structObj), table)
	rows, err := baseSQL.query(query, pk)
	structObj, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
}

func selectAll[T any](baseSQL *baseSQL, table string) ([]T, error) {
	var structObj T
	query := fmt.Sprintf("SELECT %s FROM %s", getDbTags(structObj), table)
	rows, _ := baseSQL.query(query)
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func deleteById[T any](baseSQL *baseSQL, table string, pk int) (*T, error) {
	var structObj T
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING %s`, table, getDbTags(structObj))
	rows, err := baseSQL.query(query, pk)
	structObj, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
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

func insert[T any](baseSQL *baseSQL, table string, createStruct interface{}) (*T, error) {
	var returningStruct T
	query, args, err := getInsertQuery(table, createStruct, returningStruct)
	if err != nil {
		return nil, err
	}
	rows, err := baseSQL.query(query, args...)
	returningStruct, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &returningStruct, err
}

func update[T any](baseSQL *baseSQL, table string, pk int, editStruct interface{}) (*T, error) {
	var returningStruct T
	query, args, err := getUpdateQuery(table, pk, editStruct, returningStruct)
	if err != nil {
		return nil, err
	}
	rows, err := baseSQL.query(query, args...)
	returningStruct, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &returningStruct, err
}

func getInsertQuery(table string, createInterface interface{}, returningInterface interface{}) (string, []interface{}, error) {
	// Получаем тип структуры
	userType := reflect.TypeOf(createInterface)
	// Получаем значение структуры
	userValue := reflect.ValueOf(createInterface)
	var valuesArray []interface{}
	var fieldsArray []string
	var indexRowArray []string
	var placeholder = 1
	for i := 0; i < userType.NumField(); i++ {
		value := userValue.Field(i)
		valuesArray = append(valuesArray, value.Interface())
		// Получаем название поля
		fieldName := userType.Field(i).Tag.Get("db")
		// Добавляем позиционный индекс
		placeholderStr := fmt.Sprintf("$%d", placeholder)
		indexRowArray = append(indexRowArray, placeholderStr)
		fieldsArray = append(fieldsArray, fieldName)
		placeholder++

	}
	if len(fieldsArray) == 0 {
		return "", nil, errors.New("empty createInterface")
	}

	fields := strings.Join(fieldsArray, ", ")
	placeholders := strings.Join(indexRowArray, ", ")
	returning := getDbTags(returningInterface)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", table, fields, placeholders, returning)

	return query, valuesArray, nil
}

func getUpdateQuery(table string, pk int, setInterface interface{}, returningInterface interface{}) (string, []interface{}, error) {
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
		fieldName := userType.Field(i).Tag.Get("db")
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

	returning := getDbTags(returningInterface)

	query := fmt.Sprintf("UPDATE %s SET %s  WHERE id = $%v RETURNING %s", table, set, placeholder, returning)
	return query, values, nil
}

func total(baseSQL *baseSQL, table string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s ", table)
	var count int
	return count, baseSQL.queryRow(query).Scan(&count)
}
