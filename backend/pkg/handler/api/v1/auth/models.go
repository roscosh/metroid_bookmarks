package auth

import (
	"metroid_bookmarks/misc/session"
)

type responseMe struct {
	*session.Session
}

type formLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type responseLogin struct {
	*session.Session
}

type responseLogout struct {
	*session.Session
}
