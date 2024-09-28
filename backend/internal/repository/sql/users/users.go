package users

import (
	"context"
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const usersTable = "users"

type usersSQL struct {
	sql pgpool.SQL[User]
}

type SQL interface {
	Create(createForm *CreateUser) (*User, error)
	Delete(id int) (*User, error)
	Edit(id int, editForm *EditUser) (*User, error)
	GetAll(search string) ([]User, error)
	GetByCredentials(login, password string) (*User, error)
	GetByID(id int) (*User, error)
	Total() (int, error)
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[User](dbPool, usersTable)
	return &usersSQL{sql: sql}
}

func (s *usersSQL) Create(createForm *CreateUser) (*User, error) {
	entity, err := s.sql.Insert(context.Background(), createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *usersSQL) Delete(id int) (*User, error) {
	entity, err := s.sql.Delete(context.Background(), id)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *usersSQL) Edit(id int, editForm *EditUser) (*User, error) {
	entity, err := s.sql.Update(context.Background(), id, editForm)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *usersSQL) GetAll(search string) ([]User, error) {
	if search != "" {
		return s.sql.SelectAllWhere(context.Background(), "LOWER(name) LIKE $1 OR LOWER(login) LIKE $2", search, search)
	}

	return s.sql.SelectAll(context.Background())
}

func (s *usersSQL) GetByCredentials(login, password string) (*User, error) {
	return s.sql.SelectWhere(context.Background(), "login = $1 AND password = $2", login, password)
}

func (s *usersSQL) GetByID(id int) (*User, error) {
	entity, err := s.sql.Select(context.Background(), id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *usersSQL) Total() (int, error) {
	return s.sql.Total(context.Background())
}
