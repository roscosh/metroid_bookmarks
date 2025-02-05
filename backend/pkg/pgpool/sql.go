package pgpool

import (
	"context"
	"errors"
	"fmt"
	"metroid_bookmarks/pkg/misc"
	"reflect"
	"strings"
)

var (
	ErrNotPointerStruct = errors.New("object must be a pointer of structure")
	ErrAnonymousField   = errors.New("anonymous fields are not allowed")
	ErrUnexportedField  = errors.New("unexported fields are not allowed")
	ErrMissingDBTag     = errors.New("field must have a 'db' tag")
	ErrEmptyStruct      = errors.New("empty struct")
	ErrEmptyWhereClause = errors.New("WHERE clause is required for UPDATE query")
	ErrNotPointerField  = errors.New("field must be a pointer")
)

// SQL provides a generic interface for performing basic CRUD operations (Create, Read, Update, Delete)
// with a database, supporting various methods of selection and filtering.
// T — the type of the entity the interface works with (e.g., a database model).
type SQL[T any] interface {
	// sqlQuery[T] provides a generic interface for executing SQL queries and collecting results.
	// It includes methods for executing queries, querying data, and collecting rows into Go structures.
	// T — the type of the entity (e.g., a database model) that the interface works with.
	sqlQuery[T]

	// Delete removes a record by primary key (pk) and returns the deleted object.
	// ctx — context for managing the request.
	// pk — the primary key of the record to delete.
	// Returns a pointer to the deleted object or an error.
	Delete(ctx context.Context, pk int) (*T, error)

	// DeleteWhere removes records that match the provided condition (whereStatement).
	// ctx — context for managing the request.
	// whereStatement — the condition used to filter the records to delete.
	// args — arguments for the filtering condition.
	// Returns a pointer to the deleted object or an error.
	DeleteWhere(ctx context.Context, whereStatement string, args ...any) (*T, error)

	// Insert adds a new record to the database.
	// ctx — context for managing the request.
	// createStruct — a pointer to the structure containing the data for the new record.
	// If a non-pointer is provided, it returns an ErrNotPointerStruct error.
	// If a pointer to an empty structure is provided, it returns an ErrEmptyStruct error.
	// Returns a pointer to the created object or an error.
	Insert(ctx context.Context, createStruct any) (*T, error)

	// Select retrieves a single record by primary key (pk).
	// ctx — context for managing the request.
	// pk — the primary key of the record to retrieve.
	// Returns a pointer to the object or an error.
	Select(ctx context.Context, pk int) (*T, error)

	// SelectWhere retrieves a single record that matches the provided condition (whereStatement).
	// ctx — context for managing the request.
	// whereStatement — the condition used to filter the records.
	// args — arguments for the filtering condition.
	// Returns a pointer to the object or an error.
	SelectWhere(ctx context.Context, whereStatement string, args ...any) (*T, error)

	// SelectAll retrieves all records from the database.
	// ctx — context for managing the request.
	// Returns a slice of objects or an error.
	SelectAll(ctx context.Context) ([]T, error)

	// SelectAllWhere retrieves records that match the provided condition (whereStatement).
	// ctx — context for managing the request.
	// whereStatement — the condition used to filter the records.
	// args — arguments for the filtering condition.
	// Returns a slice of objects or an error.
	SelectAllWhere(ctx context.Context, whereStatement string, args ...any) ([]T, error)

	// Total returns the total number of records in the database.
	// ctx — context for managing the request.
	// Returns the number of records or an error.
	Total(ctx context.Context) (int, error)

	// Update modifies a record by primary key (pk) with the data provided in editStruct.
	// ctx — context for managing the request.
	// pk — the primary key of the record to update.
	// editStruct — a pointer to the structure containing the updated data, where every field must be a pointer too.
	// If a non-pointer is provided, it returns an ErrNotPointerStruct error.
	// If a non-pointer field of structure is provided, it returns an ErrNotPointerField error.
	// If a pointer to an empty structure or zero value of structure is provided, it returns an ErrEmptyStruct error.
	// Returns a pointer to the updated object or an error.
	Update(ctx context.Context, pk int, editStruct any) (*T, error)

	// UpdateWhere modifies records that match the provided condition (whereStatement)
	// using the data provided in editStruct.
	// ctx — context for managing the request.
	// editStruct — a pointer to the structure containing the updated data, where every field must be a pointer too.
	// If a non-pointer is provided, it returns an ErrNotPointerStruct error.
	// If a non-pointer field of structure is provided, it returns an ErrNotPointerField error.
	// If a pointer to an empty structure or zero value of structure is provided, it returns an ErrEmptyStruct error.
	// where — the condition used to filter the records to update.
	// args — arguments for the filtering condition.
	// Returns a pointer to the updated object or an error.
	UpdateWhere(ctx context.Context, editStruct any, where string, args ...any) (*T, error)
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

func (s *sql[T]) Delete(ctx context.Context, pk int) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING %s`, s.table, s.columns)

	rows, err := s.Query(ctx, query, pk)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) DeleteWhere(ctx context.Context, whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s RETURNING %s`, s.table, whereStatement, s.columns)

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) Insert(ctx context.Context, createStruct any) (*T, error) {
	query, args, err := s.getInsertQuery(createStruct)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) Select(ctx context.Context, pk int) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", s.columns, s.table)

	rows, err := s.Query(ctx, query, pk)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) SelectAllWhere(ctx context.Context, whereStatement string, args ...any) ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectRows(rows)
}

func (s *sql[T]) SelectWhere(ctx context.Context, whereStatement string, args ...any) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.columns, s.table, whereStatement)

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) SelectAll(ctx context.Context) ([]T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", s.columns, s.table)

	rows, err := s.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return s.CollectRows(rows)
}

func (s *sql[T]) Total(ctx context.Context) (int, error) {
	var count int

	query := "SELECT COUNT(*) FROM " + s.table

	return count, s.QueryRow(ctx, query).Scan(&count)
}

func (s *sql[T]) Update(ctx context.Context, pk int, editStruct any) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, "id=$1", pk)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) UpdateWhere(ctx context.Context, editStruct any, where string, args ...any) (*T, error) {
	query, args, err := s.getUpdateQuery(editStruct, where, args...)
	if err != nil {
		return nil, err
	}

	rows, err := s.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.CollectOneRow(rows)
}

func (s *sql[T]) getInsertQuery(createInterface any) (string, []any, error) {
	t := reflect.TypeOf(createInterface)
	if t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Struct {
		return "", nil, ErrNotPointerStruct
	}

	elem := reflect.ValueOf(createInterface).Elem()
	fieldsAmount := elem.NumField()
	valuesArray := make([]any, 0, fieldsAmount)
	fieldNamesArray := make([]string, 0, fieldsAmount)
	indexRowArray := make([]string, 0, fieldsAmount)
	placeholder := 1

	if fieldsAmount == 0 {
		return "", nil, ErrEmptyStruct
	}

	for i := range fieldsAmount {
		// Получаем название поля
		structField := elem.Type().Field(i)
		fieldName := structField.Tag.Get("db")

		// Проверяем на анонимные поля
		if structField.Anonymous {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrAnonymousField)
		}

		// Проверяем на неэкспортируемые поля
		if !structField.IsExported() {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrUnexportedField)
		}

		// Проверяем на наличие тега `db`
		if fieldName == "" {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrMissingDBTag)
		}

		fieldNamesArray = append(fieldNamesArray, fieldName)

		// Получаем значение поля
		value := elem.Field(i)
		valuesArray = append(valuesArray, value.Interface())

		// Добавляем позиционный индекс
		placeholderStr := fmt.Sprintf("$%d", placeholder)
		indexRowArray = append(indexRowArray, placeholderStr)
		placeholder++
	}

	fields := strings.Join(fieldNamesArray, ", ")
	placeholders := strings.Join(indexRowArray, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", s.table, fields, placeholders, s.columns)

	return query, valuesArray, nil
}

func (s *sql[T]) getUpdateQuery(
	setInterface any,
	where string,
	args ...any,
) (string, []any, error) {
	queryArray := make([]string, 0, 3) //nolint:mnd

	t := reflect.TypeOf(setInterface)
	if t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Struct {
		return "", nil, ErrNotPointerStruct
	}

	if where == "" {
		return "", nil, ErrEmptyWhereClause
	}

	elem := reflect.ValueOf(setInterface).Elem()
	fieldsAmount := elem.NumField()
	fields := make([]string, 0, fieldsAmount)
	placeholder := 1 + len(args)

	for i := range fieldsAmount {
		// Получаем название поля
		structField := elem.Type().Field(i)
		fieldName := structField.Tag.Get("db")

		// Проверяем на анонимные поля
		if structField.Anonymous {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrAnonymousField)
		}

		// Проверяем на неэкспортируемые поля
		if !structField.IsExported() {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrUnexportedField)
		}

		// Проверяем на наличие тега `db`
		if fieldName == "" {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrMissingDBTag)
		}

		// Проверяем поле является ли оно указателем
		if structField.Type.Kind() != reflect.Pointer {
			return "", nil, fmt.Errorf("field '%s' is invalid: %w", structField.Name, ErrNotPointerField)
		}

		// Получаем значение поля
		value := elem.Field(i)
		if value.IsNil() {
			continue
		}

		args = append(args, value.Interface())

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

	queryArray = append(queryArray, "WHERE "+where)

	queryArray = append(queryArray, "RETURNING "+s.columns)

	query := strings.Join(queryArray, " ")

	return query, args, nil
}

func getDBTags[T any]() string {
	tags := misc.GetTags[T]("db")
	return strings.Join(tags, ", ")
}
