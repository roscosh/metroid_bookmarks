package pgpool

import (
	"errors"
	"fmt"
	"metroid_bookmarks/pkg/misc"
	"reflect"
	"strings"
)

var (
	ErrPointerStruct = errors.New("object must be a pointer of structure")
	ErrEmptyStruct   = errors.New("empty struct")
)

type SQL[T any] interface {
	sqlQuery[T]
	Delete(pk int) (*T, error)
	DeleteWhere(whereStatement string, args ...any) (*T, error)
	Insert(createStruct interface{}) (*T, error)
	SelectMany() ([]T, error)
	SelectManyWhere(whereStatement string, args ...any) ([]T, error)
	SelectOne(pk int) (*T, error)
	SelectWhere(whereStatement string, args ...any) (*T, error)
	Total() (int, error)
	Update(pk int, editStruct interface{}) (*T, error)
	UpdateWhere(editStruct interface{}, where string, args ...any) (*T, error)
}

func NewSQL[T any](dbPool *PgPool, table string) SQL[T] {
	return &sql[T]{
		sqlQuery: &baseQuery[T]{dbPool},
		table:    table,
		columns:  getDBTags[T](),
	}
}

type sql[T any] struct {
	sqlQuery[T]
	table   string
	columns string
}

func (s *sql[T]) Delete(pk int) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING %s`, s.table, s.columns)

	rows, err := s.Query(query, pk)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) DeleteWhere(whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s RETURNING %s`, s.table, whereStatement, s.columns)

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) Insert(createStruct interface{}) (*T, error) {
	query, args, err := s.getInsertQuery(createStruct)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) SelectOne(pk int) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", s.columns, s.table)

	rows, err := s.Query(query, pk)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) SelectManyWhere(whereStatement string, args ...any) ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectRows(rows)
}

func (s *sql[T]) SelectWhere(whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) SelectMany() ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", s.columns, s.table)

	rows, err := s.Query(query)
	if err != nil {
		return nil, err
	}

	return s.CollectRows(rows)
}

func (s *sql[T]) Total() (int, error) {
	var count int

	query := "SELECT COUNT(*) FROM " + s.table

	return count, s.QueryRow(query).Scan(&count)
}

func (s *sql[T]) Update(pk int, editStruct interface{}) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, "id=$1", pk)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) UpdateWhere(editStruct interface{}, where string, args ...any) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, where, args...)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) getInsertQuery(createInterface interface{}) (string, []interface{}, error) {
	t := reflect.TypeOf(createInterface)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return "", nil, ErrPointerStruct
	}

	elem := reflect.ValueOf(createInterface).Elem()
	valuesArray := make([]interface{}, 0, elem.NumField())
	fieldsArray := make([]string, 0, elem.NumField())
	indexRowArray := make([]string, 0, elem.NumField())
	placeholder := 1

	for i := range elem.NumField() {
		value := elem.Field(i)
		valuesArray = append(valuesArray, value.Interface())
		// Получаем название поля
		fieldName := elem.Type().Field(i).Tag.Get("db")
		// Добавляем позиционный индекс
		placeholderStr := fmt.Sprintf("$%d", placeholder)
		indexRowArray = append(indexRowArray, placeholderStr)
		fieldsArray = append(fieldsArray, fieldName)
		placeholder++
	}

	if len(fieldsArray) == 0 {
		return "", nil, ErrEmptyStruct
	}

	fields := strings.Join(fieldsArray, ", ")
	placeholders := strings.Join(indexRowArray, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", s.table, fields, placeholders, s.columns)

	return query, valuesArray, nil
}

func (s *sql[T]) getUpdateQuery(
	setInterface interface{},
	where string,
	args ...any,
) (string, []interface{}, error) {
	queryArray := make([]string, 0, 3) //nolint:mnd

	t := reflect.TypeOf(setInterface)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return "", nil, ErrPointerStruct
	}

	elem := reflect.ValueOf(setInterface).Elem()

	fields := make([]string, 0, elem.NumField())
	placeholder := 1 + len(args)

	for i := range elem.NumField() {
		value := elem.Field(i)
		if value.IsNil() {
			continue
		}

		args = append(args, value.Interface())
		// Получаем название поля
		fieldName := elem.Type().Field(i).Tag.Get("db")
		// Добавляем позиционный индекс
		fieldStr := fmt.Sprintf("%s = $%v", fieldName, placeholder)
		fields = append(fields, fieldStr)
		placeholder++
	}

	if len(fields) == 0 {
		return "", nil, ErrEmptyStruct
	}

	set := strings.Join(fields, ", ")

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", s.table, set)
	queryArray = append(queryArray, updateQuery)

	if where != "" {
		queryArray = append(queryArray, "WHERE "+where)
	}

	queryArray = append(queryArray, "RETURNING "+s.columns)

	query := strings.Join(queryArray, " ")

	return query, args, nil
}

func getDBTags[T any]() string {
	tags := misc.GetTags[T]("db")
	return strings.Join(tags, ", ")
}
