package pgpool

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type sqlQuery[T any] interface {
	CollectOneRow(rows pgx.Rows) (*T, error)
	CollectRows(rows pgx.Rows) ([]T, error)
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

type baseQuery[T any] struct {
	*PgPool
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
