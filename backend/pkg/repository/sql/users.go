package sql

import (
	"github.com/jackc/pgx/v5"
)

const usersTable = "users"

type User struct {
	ID      int    `json:"id"       db:"id"`
	Name    string `json:"name"     db:"name"`
	Login   string `json:"login"    db:"login"`
	IsAdmin bool   `json:"is_admin" db:"is_admin"`
}

type CreateUser struct {
	Name     string `json:"name"     db:"name"     binding:"required"`
	Login    string `json:"login"    db:"login"    binding:"required"`
	Password string `json:"password" db:"password" binding:"required,min=8,max=32"`
}

type EditUser struct {
	Name     *string `db:"name"`
	Login    *string `db:"login"`
	IsAdmin  *bool   `db:"is_admin"`
	Password *string `db:"password"`
}

type UsersSQL struct {
	*baseSQL
}

func NewUsersSQL(pool *DbPool, table string) *UsersSQL {
	sql := newBaseSQl(pool, table, User{})
	return &UsersSQL{baseSQL: sql}
}

func (s *UsersSQL) Create(createForm *CreateUser) (*User, error) {
	rows, err := s.insert(*createForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *UsersSQL) GetAll(search string) ([]User, error) {
	var rows pgx.Rows
	var err error
	if search != "" {
		rows, err = s.selectWhere("WHERE LOWER(name) LIKE $1 OR LOWER(login) LIKE $2", search, search)
	} else {
		rows, err = s.selectAll()
	}
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *UsersSQL) GetByCredentials(login, password string) (*User, error) {
	rows, err := s.selectWhere("login = $1 AND password = $2", login, password)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *UsersSQL) GetByID(id int) (*User, error) {
	rows, err := s.selectById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *UsersSQL) Delete(id int) (*User, error) {
	rows, err := s.deleteById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *UsersSQL) Edit(id int, editForm *EditUser) (*User, error) {
	rows, err := s.update(id, *editForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *UsersSQL) Total() (int, error) {
	return s.total()
}

func (s *UsersSQL) collectOneRow(rows pgx.Rows) (*User, error) {
	return collectOneRow[User](rows)
}

func (s *UsersSQL) collectRows(rows pgx.Rows) ([]User, error) {
	return collectRows[User](rows)
}
