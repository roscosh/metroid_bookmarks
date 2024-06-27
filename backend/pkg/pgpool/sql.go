package pgpool

import (
	"context"
	"errors"
	"fmt"
	"metroid_bookmarks/pkg/misc"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrPointerStruct = errors.New("object must be a pointer of structure")
	ErrEmptyStruct   = errors.New("empty struct")
)

type SQL[T any] interface {
	// Query methods
	CollectOneRow(rows pgx.Rows) (*T, error)
	CollectRows(rows pgx.Rows) ([]T, error)
	Exec(query string, args ...any) (pgconn.CommandTag, error)
	Query(query string, args ...any) (pgx.Rows, error)
	QueryRow(query string, args ...any) pgx.Row
	// CRUD methods
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
		PgPool:  dbPool,
		table:   table,
		columns: getDBTags[T](),
	}
}

type sql[T any] struct {
	*PgPool
	table   string
	columns string
}

func (s *sql[T]) CollectOneRow(rows pgx.Rows) (*T, error) {
	structObj, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
}

func (s *sql[T]) CollectRows(rows pgx.Rows) ([]T, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func (s *sql[T]) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	return s.pool.Exec(context.Background(), query, args...)
}

func (s *sql[T]) Query(query string, args ...any) (pgx.Rows, error) {
	return s.pool.Query(context.Background(), query, args...)
}

func (s *sql[T]) QueryRow(query string, args ...any) pgx.Row {
	return s.pool.QueryRow(context.Background(), query, args...)
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
	query, args, err := s.GetInsertQuery(createStruct)
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
	query, args, err := s.GetUpdateQuery(editStruct, "id=$1", pk)
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
	query, args, err := s.GetUpdateQuery(editStruct, where, args...)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) GetInsertQuery(createInterface interface{}) (string, []interface{}, error) {
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

func (s *sql[T]) GetUpdateQuery(
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
