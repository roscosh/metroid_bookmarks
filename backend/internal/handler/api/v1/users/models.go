package users

import (
	"metroid_bookmarks/internal/repository/sql/users"
)

type changePasswordResponse struct {
	*users.User
}

type changePasswordForm struct {
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type deleteResponse struct {
	*users.User
}

type editForm struct {
	Name    *string `json:"name"`
	Login   *string `json:"login"`
	IsAdmin *bool   `json:"is_admin"`
}

type editResponse struct {
	*users.User
}

type getAllForm struct {
	Search string `form:"search"`
}

type getAllResponse struct {
	Data  []users.User `json:"data"`
	Total int          `json:"total"`
}
