package sql

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
	iBaseSQL[User]
}

func NewUsersSQL(dbPool *DbPool) *UsersSQL {
	sql := newIBaseSQL[User](dbPool, usersTable)
	return &UsersSQL{iBaseSQL: sql}
}

func (s *UsersSQL) Create(createForm *CreateUser) (*User, error) {
	return s.insert(createForm)
}

func (s *UsersSQL) Delete(id int) (*User, error) {
	return s.delete(id)
}

func (s *UsersSQL) Edit(id int, editForm *EditUser) (*User, error) {
	return s.update(id, editForm)
}

func (s *UsersSQL) GetAll(search string) ([]User, error) {
	if search != "" {
		return s.selectManyWhere("WHERE LOWER(name) LIKE $1 OR LOWER(login) LIKE $2", search, search)
	} else {
		return s.selectMany()
	}
}

func (s *UsersSQL) GetByCredentials(login, password string) (*User, error) {
	return s.selectWhere("login = $1 AND password = $2", login, password)
}

func (s *UsersSQL) Get(id int) (*User, error) {
	return s.selectOne(id)

}

func (s *UsersSQL) Total() (int, error) {
	return s.total()
}
