package users

import "metroid_bookmarks/pkg/repository/sql"

type changePasswordResponse struct {
	*sql.User
}

type changePasswordForm struct {
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type deleteResponse struct {
	*sql.User
}

type editForm struct {
	Name    *string `json:"name"`
	Login   *string `json:"login"`
	IsAdmin *bool   `json:"is_admin"`
}

type editResponse struct {
	*sql.User
}

type getAllForm struct {
	Search string `form:"search"`
}

type getAllResponse struct {
	Data  []sql.User `json:"data"`
	Total int        `json:"total"`
}
