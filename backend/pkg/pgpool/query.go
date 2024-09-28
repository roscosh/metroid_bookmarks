package pgpool

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// sqlQuery[T] provides a generic interface for executing SQL queries and collecting results.
// It includes methods for executing queries, querying data, and collecting rows into Go structures.
//
// T — the type of the entity (e.g., a database model) that the interface works with.
type sqlQuery[T any] interface {
	// CollectOneRow collects a single row from the result set and maps it to a Go structure of type T.
	// rows — the pgx.Rows result set to collect the data from.
	// Returns a pointer to the structure or an error if something goes wrong.
	CollectOneRow(rows pgx.Rows) (*T, error)

	// CollectRows collects multiple rows from the result set and maps them to a slice of Go structures of type T.
	// rows — the pgx.Rows result set to collect the data from.
	// Returns a slice of structures or an error if something goes wrong.
	CollectRows(rows pgx.Rows) ([]T, error)

	// Exec executes a query (such as INSERT, UPDATE, or DELETE) without returning rows, but provides the command result.
	// ctx — context for managing the request.
	// query — the SQL query string to execute.
	// args — arguments for the query.
	// Returns a pgconn.CommandTag, which contains the result of the command (e.g., number of rows affected) or an error.
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)

	// Query executes a query that returns multiple rows (e.g., a SELECT statement).
	// ctx — context for managing the request.
	// query — the SQL query string to execute.
	// args — arguments for the query.
	// Returns a pgx.Rows result set or an error.
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)

	// QueryRow executes a query that returns a single row (e.g., a SELECT statement with a LIMIT 1).
	// ctx — context for managing the request.
	// query — the SQL query string to execute.
	// args — arguments for the query.
	// Returns a pgx.Row, which can be used to scan a single result.
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

type baseQuery[T any] struct {
	*PgPool // PgPool is a connection pool for PostgreSQL.
}

func (q *baseQuery[T]) CollectOneRow(rows pgx.Rows) (*T, error) {
	structObj, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &structObj, err
}

func (q *baseQuery[T]) CollectRows(rows pgx.Rows) ([]T, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func (q *baseQuery[T]) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return q.pool.Exec(ctx, query, args...)
}

func (q *baseQuery[T]) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return q.pool.Query(ctx, query, args...)
}

func (q *baseQuery[T]) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return q.pool.QueryRow(ctx, query, args...)
}
