package sql

import (
	"C"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"reflect"
	"strings"
)

var DbTagsMap = make(map[string]string)

type baseSQL struct {
	*DbPool
	table string
	model interface{}
}

func newBaseSQl(pool *DbPool, table string, model interface{}) *baseSQL {
	return &baseSQL{DbPool: pool, table: table, model: model}
}

func (s *baseSQL) exec(query string, args ...any) (pgconn.CommandTag, error) {
	return s.pool.Exec(s.ctx, query, args...)
}

func (s *baseSQL) query(query string, args ...any) (pgx.Rows, error) {
	return s.pool.Query(s.ctx, query, args...)
}

func (s *baseSQL) queryRow(query string, args ...any) pgx.Row {
	return s.pool.QueryRow(s.ctx, query, args...)
}

func (s *baseSQL) selectById(pk int) (pgx.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", getDbTags(s.model), s.table)
	return s.query(query, pk)
}

func (s *baseSQL) selectWhere(whereStatement string, args ...any) (pgx.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", getDbTags(s.model), s.table, whereStatement)
	return s.query(query, args...)
}

func (s *baseSQL) selectAll() (pgx.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", getDbTags(s.model), s.table)
	return s.query(query)
}

func (s *baseSQL) deleteById(pk int) (pgx.Rows, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING %s`, s.table, getDbTags(s.model))
	return s.query(query, pk)
}

func (s *baseSQL) deleteWhere(whereStatement string, args ...any) (pgx.Rows, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s RETURNING %s`, s.table, whereStatement, getDbTags(s.model))
	return s.query(query, args...)
}

func (s *baseSQL) insert(createStruct interface{}) (pgx.Rows, error) {
	query, args, err := getInsertQuery(s.table, createStruct, s.model)
	if err != nil {
		return nil, err
	}
	return s.query(query, args...)
}

func (s *baseSQL) update(pk int, editStruct interface{}) (pgx.Rows, error) {
	query, args, err := getUpdateQuery(s.table, editStruct, s.model, "id=$1", pk)
	if err != nil {
		return nil, err
	}
	return s.query(query, args...)
}

func (s *baseSQL) updateWhere(editStruct interface{}, where string, args ...any) (pgx.Rows, error) {
	query, args, err := getUpdateQuery(s.table, editStruct, s.model, where, args...)
	if err != nil {
		return nil, err
	}
	return s.query(query, args...)
}

func (s *baseSQL) total() (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s ", s.table)
	var count int
	return count, s.queryRow(query).Scan(&count)
}

func collectOneRow[T any](rows pgx.Rows) (*T, error) {
	structObj, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
}

func collectRows[T any](rows pgx.Rows) ([]T, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func getInsertQuery(
	table string,
	createInterface interface{},
	returningInterface interface{},
) (string, []interface{}, error) {
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

func getUpdateQuery(
	table string,
	setInterface interface{},
	returningInterface interface{},
	where string,
	args ...any,
) (string, []interface{}, error) {
	queryArray := make([]string, 0, 3)

	// Получаем тип структуры
	userType := reflect.TypeOf(setInterface)
	// Получаем значение структуры
	userValue := reflect.ValueOf(setInterface)
	var fields []string
	var placeholder = 1 + len(args)
	for i := 0; i < userType.NumField(); i++ {
		value := userValue.Field(i)
		if value.IsNil() {
			continue
		}
		args = append(args, value.Interface())
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

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", table, set)
	queryArray = append(queryArray, updateQuery)

	if where != "" {
		queryArray = append(queryArray, fmt.Sprintf("WHERE %s", where))
	}

	if returningInterface != nil {
		returning := getDbTags(returningInterface)
		if returning != "" {
			queryArray = append(queryArray, fmt.Sprintf("RETURNING %s", returning))
		}
	}
	query := strings.Join(queryArray, " ")
	return query, args, nil
}

func getDbTags(structObj interface{}) string {
	structType := reflect.TypeOf(structObj)
	structName := structType.Name()
	if dbTags, ok := DbTagsMap[structName]; ok {
		return dbTags
	}
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
			if dbTag != "" {
				dbTagArray = append(dbTagArray, dbTag)
			}

		}
	}

	traverseFields(structType)

	dbTags := strings.Join(dbTagArray, ", ")
	DbTagsMap[structName] = dbTags
	return dbTags
}
