package auth

import (
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/repository/sql"
)

type loginForm struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type loginResponse struct {
	*session.Session
}

type logoutResponse struct {
	*session.Session
}

type meResponse struct {
	*session.Session
}

type signUpForm struct {
	*sql.CreateUser
}

type signUpResponse struct {
	*sql.User
}
