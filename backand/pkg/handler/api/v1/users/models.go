package users

import "metroid_bookmarks/pkg/repository/sql"

type FormGetUsers struct {
	Search string `form:"search"`
}

type ResponseGetUsers struct {
	Data  []sql.User `json:"data"`
	Total int        `json:"total"`
}

type ResponseDeleteUser struct {
	*sql.User
}

type FormEditUser struct {
	sql.EditUser
}

type ResponseEditUser struct {
	*sql.User
}

type ResponseChangePassword struct {
	*sql.User
}

type FormChangePassword struct {
	sql.ChangePassword
}
