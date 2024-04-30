package service

import (
	"errors"
	"metroid_bookmarks/internal/repository/sql"
)

type AreasService struct {
	sql *sql.AreasSQL
}

func newAreasService(sql *sql.AreasSQL) *AreasService {
	return &AreasService{sql: sql}
}

func (s *AreasService) Create(createForm *sql.CreateArea) (*sql.Area, error) {
	area, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) Edit(id int, editForm *sql.EditArea) (*sql.Area, error) {
	if (editForm == &sql.EditArea{}) {
		return nil, errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	area, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) Delete(id int) (*sql.Area, error) {
	area, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return area, nil
}

func (s *AreasService) GetAll() ([]sql.Area, int, error) {
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
	return data, total, nil
}
