package sql

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"metroid_bookmarks/pkg/misc"
	"reflect"
	"strings"
)

type iBaseSQL[T any] interface {
	//Query methods
	collectOneRow(rows pgx.Rows) (*T, error)
	collectRows(rows pgx.Rows) ([]T, error)
	exec(query string, args ...any) (pgconn.CommandTag, error)
	query(query string, args ...any) (pgx.Rows, error)
	queryRow(query string, args ...any) pgx.Row
	//CRUD methods
	delete(pk int) (*T, error)
	deleteWhere(whereStatement string, args ...any) (*T, error)
	insert(createStruct interface{}) (*T, error)
	selectMany() ([]T, error)
	selectManyWhere(whereStatement string, args ...any) ([]T, error)
	selectOne(pk int) (*T, error)
	selectWhere(whereStatement string, args ...any) (*T, error)
	total() (int, error)
	update(pk int, editStruct interface{}) (*T, error)
	updateWhere(editStruct interface{}, where string, args ...any) (*T, error)
}

func newIBaseSQL[T any](dbPool *DbPool, table string) iBaseSQL[T] {

	return &baseSQL[T]{
		DbPool:  dbPool,
		table:   table,
		columns: getDbTags[T](),
	}
}

type baseSQL[T any] struct {
	*DbPool
	table   string
	columns string
}

func (s *baseSQL[T]) collectOneRow(rows pgx.Rows) (*T, error) {
	structObj, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
}

func (s *baseSQL[T]) collectRows(rows pgx.Rows) ([]T, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func (s *baseSQL[T]) exec(query string, args ...any) (pgconn.CommandTag, error) {
	return s.pool.Exec(s.ctx, query, args...)
}

func (s *baseSQL[T]) query(query string, args ...any) (pgx.Rows, error) {
	return s.pool.Query(s.ctx, query, args...)
}

func (s *baseSQL[T]) queryRow(query string, args ...any) pgx.Row {
	return s.pool.QueryRow(s.ctx, query, args...)
}

func (s *baseSQL[T]) delete(pk int) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING %s`, s.table, s.columns)
	rows, err := s.query(query, pk)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) deleteWhere(whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s RETURNING %s`, s.table, whereStatement, s.columns)
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) insert(createStruct interface{}) (*T, error) {
	query, args, err := s.getInsertQuery(createStruct)
	if err != nil {
		return nil, err
	}
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) selectOne(pk int) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", s.columns, s.table)
	rows, err := s.query(query, pk)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)

}

func (s *baseSQL[T]) selectManyWhere(whereStatement string, args ...any) ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *baseSQL[T]) selectWhere(whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) selectMany() ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", s.columns, s.table)
	rows, err := s.query(query)
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *baseSQL[T]) total() (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s ", s.table)
	var count int
	return count, s.queryRow(query).Scan(&count)
}

func (s *baseSQL[T]) update(pk int, editStruct interface{}) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, "id=$1", pk)
	if err != nil {
		return nil, err
	}
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) updateWhere(editStruct interface{}, where string, args ...any) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, where, args...)
	if err != nil {
		return nil, err
	}
	rows, err := s.query(query, args...)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *baseSQL[T]) getInsertQuery(createInterface interface{}) (string, []interface{}, error) {

	t := reflect.TypeOf(createInterface)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return "", nil, errors.New("object must be a pointer of structure")
	}

	elem := reflect.ValueOf(createInterface).Elem()
	var valuesArray []interface{}
	var fieldsArray []string
	var indexRowArray []string
	var placeholder = 1

	for i := 0; i < elem.NumField(); i++ {
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
		return "", nil, errors.New("empty createInterface")
	}

	fields := strings.Join(fieldsArray, ", ")
	placeholders := strings.Join(indexRowArray, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", s.table, fields, placeholders, s.columns)

	return query, valuesArray, nil
}

func (s *baseSQL[T]) getUpdateQuery(
	setInterface interface{},
	where string,
	args ...any,
) (string, []interface{}, error) {
	queryArray := make([]string, 0, 3)

	t := reflect.TypeOf(setInterface)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return "", nil, errors.New("object must be a pointer of structure")
	}

	elem := reflect.ValueOf(setInterface).Elem()

	var fields []string
	var placeholder = 1 + len(args)
	for i := 0; i < elem.NumField(); i++ {
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
		return "", nil, errors.New("empty setInterface")
	}
	set := strings.Join(fields, ", ")

	updateQuery := fmt.Sprintf("UPDATE %s SET %s", s.table, set)
	queryArray = append(queryArray, updateQuery)

	if where != "" {
		queryArray = append(queryArray, fmt.Sprintf("WHERE %s", where))
	}

	queryArray = append(queryArray, fmt.Sprintf("RETURNING %s", s.columns))

	query := strings.Join(queryArray, " ")
	return query, args, nil
}

func getDbTags[T any]() string {
	tags := misc.GetTags[T]("db")
	return strings.Join(tags, ", ")
}
