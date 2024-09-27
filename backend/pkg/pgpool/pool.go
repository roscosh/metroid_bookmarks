package pgpool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Error struct {
	message string
}

func NewErr(message string) error {
	return &Error{message: message}
}

func (e *Error) Error() string {
	return e.message
}

type PgPool struct {
	pool *pgxpool.Pool
}

func NewPgPool(dsn string, minConns, maxConns int32, maxConnLifetime, maxConnIdleTime, healthCheckPeriod int64) (*PgPool, error) {
	ctx := context.Background()

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conf.MinConns = minConns
	conf.MaxConns = maxConns
	conf.MaxConnLifetime = time.Duration(maxConnLifetime)
	conf.MaxConnIdleTime = time.Duration(maxConnIdleTime)
	conf.HealthCheckPeriod = time.Duration(healthCheckPeriod)

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &PgPool{pool: pool}, nil
}

func (p *PgPool) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = NewErr(fmt.Sprintf("postgreSQL couldn't close %s", r))
		}
	}()

	p.pool.Close()

	return nil
}
