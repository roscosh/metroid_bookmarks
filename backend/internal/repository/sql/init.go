package sql

import (
	"metroid_bookmarks/internal/repository/sql/areas"
	"metroid_bookmarks/internal/repository/sql/bookmarks"
	"metroid_bookmarks/internal/repository/sql/photos"
	"metroid_bookmarks/internal/repository/sql/rooms"
	"metroid_bookmarks/internal/repository/sql/skills"
	"metroid_bookmarks/internal/repository/sql/users"
	"metroid_bookmarks/pkg/pgpool"
)

type SQL struct {
	Users     *users.SQL
	Areas     *areas.SQL
	Rooms     *rooms.SQL
	Skills    *skills.SQL
	Bookmarks *bookmarks.SQL
	Photos    *photos.SQL
}

func NewSQL(dbPool *pgpool.DbPool) *SQL {
	return &SQL{
		Users:     users.NewSQL(dbPool),
		Areas:     areas.NewSQL(dbPool),
		Rooms:     rooms.NewSQL(dbPool),
		Skills:    skills.NewSQL(dbPool),
		Bookmarks: bookmarks.NewSQL(dbPool),
		Photos:    photos.NewSQL(dbPool),
	}
}
