package sql

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
)

const usersTable = "users"

type UsersSQL struct {
	*baseSQL
}

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
	IsAdmin  *bool  `json:"is_admin" db:"is_admin" binding:"required"`
}

type EditUser struct {
	Name    *string `json:"name"     db:"name"`
	Login   *string `json:"login"    db:"login"`
	IsAdmin *bool   `json:"is_admin" db:"is_admin"`
}

type ChangePassword struct {
	Password *string `json:"password" db:"password" binding:"required,min=8,max=32"`
}

func NewUsersSQL(baseSQL *baseSQL) *UsersSQL {
	return &UsersSQL{baseSQL: baseSQL}
}

func (s *UsersSQL) GetUserByID(id int) (*User, error) {
	return selectById[User](s.baseSQL, usersTable, id)
}

func (s *UsersSQL) Create(userForm *CreateUser) (*User, error) {
	return create[User](s.baseSQL, usersTable, *userForm)
}

func (s *UsersSQL) GetUserByCredentials(login, password string) (*User, error) {
	var user User
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE login = $1 AND password = $2`, getDbTags(user), usersTable)
	rows, err := s.baseSQL.Query(query, login, password)
	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Delete(id int) (*User, error) {
	return deleteById[User](s.baseSQL, usersTable, id)
}

func (s *UsersSQL) Edit(id int, form EditUser) (*User, error) {
	return edit[User](s.baseSQL, usersTable, id, form)
}

func (s *UsersSQL) Search(search string) ([]User, error) {
	var user User
	tableString := fmt.Sprintf(`SELECT %s FROM %s`, getDbTags(user), usersTable)
	if search != "" {
		search = strings.ToLower(search)
		tableString += fmt.Sprintf(` WHERE LOWER(name) LIKE '%%%s%%' OR LOWER(login) LIKE '%%%s%%'`, search, search)
	}
	query := fmt.Sprintf("%s ORDER BY id DESC", tableString)
	rows, _ := s.baseSQL.Query(query)
	return pgx.CollectRows(rows, pgx.RowToStructByName[User])
}

func (s *UsersSQL) ChangePassword(id int, form ChangePassword) (*User, error) {
	return edit[User](s.baseSQL, usersTable, id, form)
}

func (s *UsersSQL) Total() (int, error) {
	query := total(usersTable)
	var count int
	return count, s.baseSQL.QueryRow(query).Scan(&count)
}
