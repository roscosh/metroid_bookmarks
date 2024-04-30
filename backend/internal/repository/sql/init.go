package sql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"metroid_bookmarks/pkg/misc"
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
		Users:     NewUsersSQL(dbPool, usersTable),
		Areas:     NewAreasSQL(dbPool, areasTable),
		Rooms:     NewRoomsSQL(dbPool, roomsTable),
		Skills:    NewSkillsSQL(dbPool, skillsTable),
		Bookmarks: NewBookmarksSQL(dbPool, bookmarksTable),
		Photos:    NewPhotosSQL(dbPool, photosTable),
	}
}

type DbPool struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func (d *DbPool) Close() {
	d.pool.Close()
}

func NewDbPool(dsn string) (*DbPool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DbPool{pool: pool, ctx: ctx}, nil
}
