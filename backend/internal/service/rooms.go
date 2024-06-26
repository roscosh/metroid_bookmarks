package service

import (
	"metroid_bookmarks/internal/repository/sql/rooms"
)

type RoomsService struct {
	sql *rooms.SQL
}

func newRoomsService(sql *rooms.SQL) *RoomsService {
	return &RoomsService{sql: sql}
}

func (s *RoomsService) Create(createForm *rooms.CreateRoom) (*rooms.Room, error) {
	room, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) Edit(id int, editForm *rooms.EditRoom) (*rooms.Room, error) {
	if (editForm == &rooms.EditRoom{}) {
		return nil, ErrEmptyStruct
	}
	room, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) Delete(id int) (*rooms.Room, error) {
	room, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return room, nil
}

func (s *RoomsService) GetAll() ([]rooms.Room, int, error) {
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
		data = []rooms.Room{}
	}
	return data, total, nil
}
