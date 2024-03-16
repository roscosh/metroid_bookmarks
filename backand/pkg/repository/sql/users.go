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

type User struct {
	ID       int    `json:"id"       db:"id"`
	Name     string `json:"name"     db:"name"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role"     db:"role"`
	Status   bool   `json:"status"   db:"status"`
}

type EditUser struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Role     *string `json:"role"`
	Status   *bool   `json:"status"`
}

type ChangePassword struct {
	Password *string `json:"password" binding:"required,min=8,max=32"`
}

func NewUsersSQL(pool *postgresPool) *UsersSQL {
	return &UsersSQL{postgresPool: pool}
}

func (s *UsersSQL) GetUserByID(userID int) (*User, error) {
	query := fmt.Sprintf("SELECT  id, name, username, role, status FROM %s WHERE id = $1", usersTable)
	rows, err := s.pool.Query(s.ctx, query, userID)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) Create(name, username, password, role string) (*User, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, username, password, role) VALUES ($1, $2, $3, $4) RETURNING id, name, username, role, status`, usersTable)
	rows, _ := s.pool.Query(s.ctx, query, name, username, password, role)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	return &user, err
}

func (s *UsersSQL) GetUserByCredentials(username, password string) (*User, error) {
	query := fmt.Sprintf(`SELECT id, name, username, role, status FROM %s WHERE username = $1 AND password = $2`, usersTable)
	rows, err := s.pool.Query(s.ctx, query, username, password)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *UsersSQL) Delete(id int) (*User, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id, name, username, role, status`, usersTable)
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
	tableString := fmt.Sprintf(`SELECT id, name, username, role, status FROM %s`, usersTable)
	if search != "" {
		search = strings.ToLower(search)
		tableString += fmt.Sprintf(` WHERE LOWER(name) LIKE '%%%s%%' OR LOWER(username) LIKE '%%%s%%'`, search, search)
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
