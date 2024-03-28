package sql

import "github.com/jackc/pgx/v5"

const roomsTable = "rooms"

type Room struct {
	Id     int    `json:"id"      db:"id"`
	NameEn string `json:"name_en" db:"name_en"`
	NameRu string `json:"name_ru" db:"name_ru"`
}

type CreateRoom struct {
	NameEn string `json:"name_en" db:"name_en" binding:"required"`
	NameRu string `json:"name_ru" db:"name_ru" binding:"required"`
}

type EditRoom struct {
	NameEn *string `json:"name_en" db:"name_en"`
	NameRu *string `json:"name_ru" db:"name_ru"`
}

type RoomsSQL struct {
	*baseSQL
}

func NewRoomsSQL(pool *DbPool, table string) *RoomsSQL {
	sql := newBaseSQl(pool, table, Room{})
	return &RoomsSQL{baseSQL: sql}
}

func (s *RoomsSQL) Create(createForm *CreateRoom) (*Room, error) {
	rows, err := s.insert(*createForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *RoomsSQL) Edit(id int, editForm *EditRoom) (*Room, error) {
	rows, err := s.update(id, *editForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *RoomsSQL) Delete(id int) (*Room, error) {
	rows, err := s.deleteById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *RoomsSQL) GetAll() ([]Room, error) {
	rows, err := s.selectAll()
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *RoomsSQL) GetByID(id int) (*Room, error) {
	rows, err := s.selectById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *RoomsSQL) Total() (int, error) {
	return s.total()
}

func (s *RoomsSQL) collectOneRow(rows pgx.Rows) (*Room, error) {
	return collectOneRow[Room](rows)
}

func (s *RoomsSQL) collectRows(rows pgx.Rows) ([]Room, error) {
	return collectRows[Room](rows)
}
