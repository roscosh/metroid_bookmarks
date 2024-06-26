package service

import (
	"metroid_bookmarks/internal/repository/sql/areas"
)

type AreasService struct {
	sql *areas.SQL
}

func newAreasService(sql *areas.SQL) *AreasService {
	return &AreasService{sql: sql}
}

func (s *AreasService) Create(createForm *areas.CreateArea) (*areas.Area, error) {
	area, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) Edit(id int, editForm *areas.EditArea) (*areas.Area, error) {
	if (editForm == &areas.EditArea{}) {
		return nil, ErrEmptyStruct
	}
	area, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) Delete(id int) (*areas.Area, error) {
	area, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) GetAll() ([]areas.Area, int, error) {
	data, err := s.sql.GetAll()
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	total, err := s.sql.Total()
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	if data == nil {
		data = []areas.Area{}
	}
	return data, total, nil
}
