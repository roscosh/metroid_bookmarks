package sql

import "metroid_bookmarks/misc"

var logger = misc.GetLogger()

type SQL struct {
	Users     *UsersSQL
	Areas     *AreasSQL
	Rooms     *RoomsSQL
	Skills    *SkillsSQL
	Bookmarks *BookmarksSQL
	Photos    *PhotosSQL
}

func (s *SQL) Close() {
	s.Users.baseSQL.pool.Close()
}

func NewSQL(dsn string) (*SQL, error) {
	pool, err := newPostgresPool(dsn)
	if err != nil {
		return nil, err
	}
	return &SQL{
		Users:     NewUsersSQL(pool),
		Areas:     NewAreasSQL(pool),
		Rooms:     NewRoomsSQL(pool),
		Skills:    NewSkillsSQL(pool),
		Bookmarks: NewBookmarksSQL(pool),
		Photos:    NewPhotosSQL(pool),
	}, nil
}
