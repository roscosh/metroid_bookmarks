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

func (s *UsersService) ChangePassword(userID int, password string) (*users.User, error) {
	token := generatePasswordHash(password)
	editForm := users.EditUser{Password: &token}

	user, err := s.sql.Edit(userID, &editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, userID)

		return nil, err
	}

	return user, nil
}

func (s *UsersService) Delete(userID int) (*users.User, error) {
	user, err := s.sql.Delete(userID)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, userID)

		return nil, err
	}

	return user, nil
}

func (s *UsersService) Edit(userID int, editForm *users.EditUser) (*users.User, error) {
	if (editForm == &users.EditUser{}) {
		return nil, ErrEmptyStruct
	}

	user, err := s.sql.Edit(userID, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, userID)

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
