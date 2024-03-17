package users

import "metroid_bookmarks/pkg/repository/sql"

type formGetUsers struct {
	Search string `form:"search"`
}

type responseGetUsers struct {
	Data  []sql.User `json:"data"`
	Total int        `json:"total"`
}

type responseDeleteUser struct {
	*sql.User
}

type formEditUser struct {
	sql.EditUser
}

type responseEditUser struct {
	*sql.User
}

type responseChangePassword struct {
	*sql.User
}

type formChangePassword struct {
	sql.ChangePassword
}

type FormCreateUser struct {
	*sql.CreateUser
}

type ResponseCreateUser struct {
	*sql.User
}
