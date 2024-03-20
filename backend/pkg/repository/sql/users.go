package sql

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
)

const usersTable = "users"

type UsersSQL struct {
	*postgresPool
}

type baseUser struct {
	ID   int    `json:"id"       db:"id"`
	Name string `json:"name"     db:"name"`
}

type User struct {
	baseUser
	Login   string `json:"login"    db:"login"`
	IsAdmin bool   `json:"is_admin" db:"is_admin"`
}

type CreateUser struct {
	Name     string `json:"name"     binding:"required"`
	Login    string `json:"login"    binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	IsAdmin  bool   `json:"is_admin" binding:"required"`
}

type EditUser struct {
	Name    *string `json:"name"`
	Login   *string `json:"login"`
	IsAdmin *bool   `json:"is_admin"`
}

type ChangePassword struct {
	Password *string `json:"password" binding:"required,min=8,max=32"`
}

func NewUsersSQL(pool *postgresPool) *UsersSQL {
	return &UsersSQL{postgresPool: pool}
}

func (s *UsersSQL) GetUserByID(userID int) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", getDbTags(User{}), usersTable)
	rows, err := s.pool.Query(s.ctx, query, userID)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Create(userForm *CreateUser) (*User, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, login, password, is_admin) VALUES ($1, $2, $3, $4) RETURNING id, name, login, is_admin`, usersTable)
	rows, _ := s.pool.Query(s.ctx, query, userForm.Name, userForm.Login, userForm.Password, userForm.IsAdmin)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) GetUserByCredentials(login, password string) (*User, error) {
	query := fmt.Sprintf(`SELECT id, name, login, is_admin FROM %s WHERE login = $1 AND password = $2`, usersTable)
	rows, err := s.pool.Query(s.ctx, query, login, password)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *UsersSQL) Delete(id int) (*User, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id, name, login, is_admin`, usersTable)
	rows, _ := s.pool.Query(s.ctx, query, id)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Edit(id int, form EditUser) (*User, error) {
	query, args, err := update(usersTable, id, form, User{})
	if err != nil {
		return nil, err
	}
	rows, _ := s.pool.Query(s.ctx, query, args...)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Search(search string) ([]User, error) {
	tableString := fmt.Sprintf(`SELECT id, name, login, is_admin FROM %s`, usersTable)
	if search != "" {
		search = strings.ToLower(search)
		tableString += fmt.Sprintf(` WHERE LOWER(name) LIKE '%%%s%%' OR LOWER(login) LIKE '%%%s%%'`, search, search)
	}
	query := fmt.Sprintf("%s ORDER BY id DESC", tableString)
	rows, _ := s.pool.Query(s.ctx, query)
	return pgx.CollectRows(rows, pgx.RowToStructByName[User])
}

func (s *UsersSQL) ChangePassword(userId int, form ChangePassword) (*User, error) {
	query, args, err := update(usersTable, userId, form, User{})
	if err != nil {
		return nil, err
	}
	rows, err := s.pool.Query(s.ctx, query, args...)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Total() (int, error) {
	query := total(usersTable)
	var count int
	return count, s.pool.QueryRow(s.ctx, query).Scan(&count)
}
