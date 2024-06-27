package auth

import (
	"metroid_bookmarks/internal/repository/sql/users"
	"metroid_bookmarks/pkg/session"
)

type loginForm struct {
	Login    string `binding:"required"              json:"login"`
	Password string `binding:"required,min=8,max=32" json:"password"`
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
