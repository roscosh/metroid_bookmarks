package areas

import (
	"metroid_bookmarks/pkg/pgpool"
)

const areasTable = "areas"

type SQL struct {
	sql pgpool.SQL[Area]
}

func NewSQL(dbPool *pgpool.PgPool) *SQL {
	sql := pgpool.NewSQL[Area](dbPool, areasTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateArea) (*Area, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Delete(id int) (*Area, error) {
	return s.sql.Delete(id)
}

func (s *SQL) Edit(id int, editForm *EditArea) (*Area, error) {
	return s.sql.Update(id, editForm)
}

func (s *SQL) GetAll() ([]Area, error) {
	return s.sql.SelectMany()
}

func (s *SQL) GetByID(id int) (*Area, error) {
	return s.sql.SelectOne(id)
}

func (s *SQL) Total() (int, error) {
	return s.sql.Total()
}
