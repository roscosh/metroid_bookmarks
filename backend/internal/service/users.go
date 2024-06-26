package service

import (
	"metroid_bookmarks/internal/repository/sql/users"
)

type UsersService struct {
	sql *users.SQL
}

func newUsersService(sql *users.SQL) *UsersService {
	return &UsersService{sql: sql}
}

func (s *UsersService) ChangePassword(id int, password string) (*users.User, error) {
	token := generatePasswordHash(password)
	editForm := users.EditUser{Password: &token}
	user, err := s.sql.Edit(id, &editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Delete(id int) (*users.User, error) {
	user, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Edit(id int, editForm *users.EditUser) (*users.User, error) {
	if (editForm == &users.EditUser{}) {
		return nil, ErrEmptyStruct
	}
	user, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return user, nil
}

func (s *UsersService) GetAll(search string) ([]users.User, int, error) {
	data, err := s.sql.GetAll(search)
	if err != nil {
		logger.Error(err.Error())

		return nil, 0, err
	}
	if data == nil {
		data = []users.User{}
	}
	return data, len(data), nil
}
