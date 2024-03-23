package sql

import "metroid_bookmarks/misc"

var logger = misc.GetLogger()

type SQL struct {
	Users *UsersSQL
	Areas *AreasSQL
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
		Users: NewUsersSQL(pool),
		Areas: NewAreasSQL(pool),
	}, nil
}
