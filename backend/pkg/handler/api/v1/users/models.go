package users

import "metroid_bookmarks/pkg/repository/sql"

type changePasswordResponse struct {
	*sql.User
}

type changePasswordForm struct {
	sql.ChangePassword
}

type deleteResponse struct {
	*sql.User
}

type editForm struct {
	*sql.EditUser
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
