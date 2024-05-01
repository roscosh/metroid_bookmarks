package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"metroid_bookmarks/pkg/misc"
	"time"
)

var logger = misc.GetLogger()

type SQL struct {
	Users     *UsersSQL
	Areas     *AreasSQL
	Rooms     *RoomsSQL
	Skills    *SkillsSQL
	Bookmarks *BookmarksSQL
	Photos    *PhotosSQL
}

func NewSQL(dbPool *DbPool) *SQL {
	return &SQL{
		Users:     NewUsersSQL(dbPool),
		Areas:     NewAreasSQL(dbPool),
		Rooms:     NewRoomsSQL(dbPool),
		Skills:    NewSkillsSQL(dbPool),
		Bookmarks: NewBookmarksSQL(dbPool),
		Photos:    NewPhotosSQL(dbPool),
	}
}

type DbPool struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func (d *DbPool) Close() error {
	var errMessage string
	d.pool.Close()
	defer func() {
		if r := recover(); r != nil {
			errMessage = fmt.Sprintf("postgreSQL couldn't close %s", r)
			return
		}
	}()
	if errMessage != "" {
		return errors.New(errMessage)
	}
	return nil
}

func NewDbPool(dsn string, minConns int32, maxConns int32, maxConnLifetime int64, maxConnIdleTime int64, healthCheckPeriod int64) (*DbPool, error) {
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
	return &DbPool{pool: pool, ctx: ctx}, nil
}
