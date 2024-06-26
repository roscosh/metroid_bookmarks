package rooms

import (
	"metroid_bookmarks/pkg/pgpool"
)

const roomsTable = "rooms"

type SQL struct {
	sql pgpool.SQL[Room]
}

func NewSQL(dbPool *pgpool.DbPool) *SQL {
	sql := pgpool.NewSQL[Room](dbPool, roomsTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateRoom) (*Room, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Edit(id int, editForm *EditRoom) (*Room, error) {
	return s.sql.Update(id, editForm)
}

func (s *SQL) Delete(id int) (*Room, error) {
	return s.sql.Delete(id)
}

func (s *SQL) GetAll() ([]Room, error) {
	return s.sql.SelectMany()
}

func (s *SQL) GetByID(id int) (*Room, error) {
	return s.sql.SelectOne(id)
}

func (s *SQL) Total() (int, error) {
	return s.sql.Total()
}
