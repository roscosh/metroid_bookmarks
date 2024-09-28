package service

import (
	"metroid_bookmarks/internal/repository/sql/rooms"
)

type RoomsService struct {
	sql rooms.SQL
}

func newRoomsService(sql rooms.SQL) *RoomsService {
	return &RoomsService{sql: sql}
}

func (s *RoomsService) Create(createForm *rooms.CreateRoom) (*rooms.Room, error) {
	room, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return room, nil
}

func (s *RoomsService) Edit(roomID int, editForm *rooms.EditRoom) (*rooms.Room, error) {
	room, err := s.sql.Edit(roomID, editForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return room, nil
}

func (s *RoomsService) Delete(roomID int) (*rooms.Room, error) {
	room, err := s.sql.Delete(roomID)
	if err != nil {
		logger.Error(err.Error())
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

	if data == nil {
		data = []rooms.Room{}
	}

	return data, len(data), nil
}
