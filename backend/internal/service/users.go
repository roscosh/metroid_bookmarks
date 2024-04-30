package service

import (
	"errors"
	"metroid_bookmarks/internal/repository/sql"
)

type UsersService struct {
	sql *sql.UsersSQL
}

func newUsersService(sql *sql.UsersSQL) *UsersService {
	return &UsersService{sql: sql}
}

func (s *UsersService) ChangePassword(id int, password string) (*sql.User, error) {
	token := generatePasswordHash(password)
	editForm := sql.EditUser{Password: &token}
	user, err := s.sql.Edit(id, &editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Delete(id int) (*sql.User, error) {
	user, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Edit(id int, editForm *sql.EditUser) (*sql.User, error) {
	if (editForm == &sql.EditUser{}) {
		return nil, errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	user, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) GetAll(search string) ([]sql.User, int, error) {
	data, err := s.sql.GetAll(search)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	return data, len(data), nil
}
