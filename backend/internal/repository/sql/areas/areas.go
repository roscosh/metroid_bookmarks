package areas

import (
	"metroid_bookmarks/pkg/pgpool"
)

const areasTable = "areas"

type Sql struct {
	sql pgpool.SQL[Area]
}
type SQL interface {
	Create(createForm *CreateArea) (*Area, error)
	Delete(id int) (*Area, error)
	Edit(id int, editForm *EditArea) (*Area, error)
	GetAll() ([]Area, error)
	GetByID(id int) (*Area, error)
	Total() (int, error)
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[Area](dbPool, areasTable)
	return &Sql{sql: sql}
}

func (s *Sql) Create(createForm *CreateArea) (*Area, error) {
	return s.sql.Insert(createForm)
}

func (s *Sql) Delete(id int) (*Area, error) {
	return s.sql.Delete(id)
}

func (s *Sql) Edit(id int, editForm *EditArea) (*Area, error) {
	return s.sql.Update(id, editForm)
}

func (s *Sql) GetAll() ([]Area, error) {
	return s.sql.SelectMany()
}

func (s *Sql) GetByID(id int) (*Area, error) {
	return s.sql.SelectOne(id)
}

func (s *Sql) Total() (int, error) {
	return s.sql.Total()
}
