package areas

import (
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const areasTable = "areas"

type areasSQL struct {
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
	return &areasSQL{sql: sql}
}

func (s *areasSQL) Create(createForm *CreateArea) (*Area, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *areasSQL) Delete(id int) (*Area, error) {
	entity, err := s.sql.Delete(id)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *areasSQL) Edit(id int, editForm *EditArea) (*Area, error) {
	entity, err := s.sql.Update(id, editForm)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *areasSQL) GetAll() ([]Area, error) {
	return s.sql.SelectMany()
}

func (s *areasSQL) GetByID(id int) (*Area, error) {
	entity, err := s.sql.SelectOne(id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *areasSQL) Total() (int, error) {
	return s.sql.Total()
}
