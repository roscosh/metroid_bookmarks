package service

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"metroid_bookmarks/pkg/repository/sql"
)

type UsersService struct {
	sql *sql.UsersSQL
}

func newUsersService(sql *sql.UsersSQL) *UsersService {
	return &UsersService{sql: sql}
}

func (s *UsersService) Create(userForm *sql.CreateUser) (*sql.User, error) {
	userForm.Password = generatePasswordHash(userForm.Password)
	user, err := s.sql.Create(userForm)
	if err != nil {
		prefixError := "Ошибка БД: "
		var pgxErr *pgconn.PgError
		var errMessage string
		switch {
		case errors.As(err, &pgxErr):
			switch pgxErr.Code {
			case "23505":
				errMessage = fmt.Sprintf(`пользователь с логином "%s" уже существует!`, user.Login)
			}
		}
		if errMessage != "" {
			errMessage = prefixError + errMessage
			return nil, errors.New(errMessage)
		}
		logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Delete(id int) (*sql.User, error) {
	user, err := s.sql.Delete(id)
	if err != nil {
		prefixError := "Ошибка БД: "
		var errMessage string
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errMessage = fmt.Sprintf(`нет пользователя с id "%v"!`, id)
		}
		if errMessage != "" {
			errMessage = prefixError + errMessage
			return nil, errors.New(errMessage)
		}
		logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UsersService) Edit(id int, form sql.EditUser) (*sql.User, error) {
	if (form == sql.EditUser{}) {
		return nil, errors.New("необходимо заполнить хотя бы один параметр в форме!")
	}
	user, err := s.sql.Edit(id, form)
	if err != nil {
		prefixError := "Ошибка БД: "
		var pgxErr *pgconn.PgError
		var errMessage string
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errMessage = fmt.Sprintf(`нет пользователя с id "%v"!`, id)
		case errors.As(err, &pgxErr):
			switch pgxErr.Code {
			case "23505":
				errMessage = fmt.Sprintf(`пользователь с логином "%s" уже существует!`, *form.Login)
			}
		}
		if errMessage != "" {
			errMessage = prefixError + errMessage
			return nil, errors.New(errMessage)
		}
		logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UsersService) GetAllUsers(search string) ([]sql.User, int, error) {
	data, err := s.sql.Search(search)
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

func (s *UsersService) ChangePassword(id int, form sql.ChangePassword) (*sql.User, error) {
	token := generatePasswordHash(*form.Password)
	form.Password = &token
	user, err := s.sql.ChangePassword(id, form)
	if err != nil {
		prefixError := "Ошибка БД: "
		var errMessage string
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errMessage = fmt.Sprintf(`нет пользователя с id "%v"!`, id)
		}
		if errMessage != "" {
			errMessage = prefixError + errMessage
			return nil, errors.New(errMessage)
		}
		logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}
