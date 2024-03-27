package sql

import "github.com/jackc/pgx/v5"

const areasTable = "areas"

type Area struct {
	Id     int    `json:"id"      db:"id"`
	NameEn string `json:"name_en" db:"name_en"`
	NameRu string `json:"name_ru" db:"name_ru"`
}

type CreateArea struct {
	NameEn string `json:"name_en" db:"name_en" binding:"required"`
	NameRu string `json:"name_ru" db:"name_ru" binding:"required"`
}

type EditArea struct {
	NameEn *string `json:"name_en" db:"name_en"`
	NameRu *string `json:"name_ru" db:"name_ru"`
}

type AreasSQL struct {
	*baseSQL
}

func NewAreasSQL(pool *DbPool, table string) *AreasSQL {
	sql := newBaseSQl(pool, table, Area{})
	return &AreasSQL{baseSQL: sql}
}

func (s *AreasSQL) Create(createForm *CreateArea) (*Area, error) {
	rows, err := s.insert(*createForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *AreasSQL) Delete(id int) (*Area, error) {
	rows, err := s.deleteById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *AreasSQL) Edit(id int, editForm *EditArea) (*Area, error) {
	rows, err := s.update(id, *editForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *AreasSQL) GetAll() ([]Area, error) {
	rows, err := s.selectAll()
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *AreasSQL) GetByID(id int) (*Area, error) {
	rows, err := s.selectById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *AreasSQL) Total() (int, error) {
	return s.total()
}

func (s *AreasSQL) collectOneRow(rows pgx.Rows) (*Area, error) {
	return collectOneRow[Area](rows)
}

func (s *AreasSQL) collectRows(rows pgx.Rows) ([]Area, error) {
	return collectRows[Area](rows)
}
