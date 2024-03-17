package auth

import (
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/repository/sql"
)

type ResponseMe struct {
	*session.Session
}

type FormLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type ResponseLogin struct {
	*session.Session
}

type ResponseLogout struct {
	*session.Session
}

type FormCreateUser struct {
	*sql.CreateUser
}

type ResponseCreateUser struct {
	*sql.User
}
