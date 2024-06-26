package users

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
