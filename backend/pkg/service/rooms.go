package service

import (
	"errors"
	"metroid_bookmarks/pkg/repository/sql"
)

type RoomsService struct {
	sql *sql.RoomsSQL
}

func newRoomsService(sql *sql.RoomsSQL) *RoomsService {
	return &RoomsService{sql: sql}
}

func (s *RoomsService) Create(createForm *sql.CreateRoom) (*sql.Room, error) {
	room, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) Edit(id int, editForm *sql.EditRoom) (*sql.Room, error) {
	if (editForm == &sql.EditRoom{}) {
		return nil, errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	room, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) Delete(id int) (*sql.Room, error) {
	room, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) GetAll() ([]sql.Room, int, error) {
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
