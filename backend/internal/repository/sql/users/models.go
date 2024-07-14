package users

type User struct {
	Name    string `db:"name"     json:"name"`
	Login   string `db:"login"    json:"login"`
	ID      int    `db:"id"       json:"id"`
	IsAdmin bool   `db:"is_admin" json:"is_admin"`
}

type CreateUser struct {
	Name     string `binding:"required"              db:"name"     json:"name"`
	Login    string `binding:"required"              db:"login"    json:"login"`
	Password string `binding:"required,min=8,max=32" db:"password" json:"password"`
}

type EditUser struct {
	Name     *string `db:"name"`
	Login    *string `db:"login"`
	IsAdmin  *bool   `db:"is_admin"`
	Password *string `db:"password"`
}
