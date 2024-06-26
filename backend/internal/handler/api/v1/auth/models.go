package auth

import (
	"metroid_bookmarks/internal/repository/sql/users"
	"metroid_bookmarks/pkg/session"
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
	*users.CreateUser
}

type signUpResponse struct {
	*users.User
}
