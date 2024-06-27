package users

import (
	"metroid_bookmarks/pkg/pgpool"
)

const usersTable = "users"

type SQL struct {
	sql pgpool.SQL[User]
}

func NewSQL(dbPool *pgpool.PgPool) *SQL {
	sql := pgpool.NewSQL[User](dbPool, usersTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateUser) (*User, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Delete(id int) (*User, error) {
	return s.sql.Delete(id)
}

func (s *SQL) Edit(id int, editForm *EditUser) (*User, error) {
	return s.sql.Update(id, editForm)
}

func (s *SQL) GetAll(search string) ([]User, error) {
	if search != "" {
		return s.sql.SelectManyWhere("WHERE LOWER(name) LIKE $1 OR LOWER(login) LIKE $2", search, search)
	}

	return s.sql.SelectMany()
}

func (s *SQL) GetByCredentials(login, password string) (*User, error) {
	return s.sql.SelectWhere("login = $1 AND password = $2", login, password)
}

func (s *SQL) Get(id int) (*User, error) {
	return s.sql.SelectOne(id)
}

func (s *SQL) Total() (int, error) {
	return s.sql.Total()
}
