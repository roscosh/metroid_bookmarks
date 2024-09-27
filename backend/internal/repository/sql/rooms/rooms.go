package rooms

import (
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const roomsTable = "rooms"

type SQL interface {
	Create(createForm *CreateRoom) (*Room, error)
	Delete(id int) (*Room, error)
	Edit(id int, editForm *EditRoom) (*Room, error)
	GetAll() ([]Room, error)
	GetByID(id int) (*Room, error)
	Total() (int, error)
}

type roomsSQL struct {
	sql pgpool.SQL[Room]
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[Room](dbPool, roomsTable)
	return &roomsSQL{sql: sql}
}

func (s *roomsSQL) Create(createForm *CreateRoom) (*Room, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *roomsSQL) Delete(id int) (*Room, error) {
	entity, err := s.sql.Delete(id)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *roomsSQL) Edit(id int, editForm *EditRoom) (*Room, error) {
	entity, err := s.sql.Update(id, editForm)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *roomsSQL) GetAll() ([]Room, error) {
	return s.sql.SelectMany()
}

func (s *roomsSQL) GetByID(id int) (*Room, error) {
	entity, err := s.sql.SelectOne(id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *roomsSQL) Total() (int, error) {
	return s.sql.Total()
}
