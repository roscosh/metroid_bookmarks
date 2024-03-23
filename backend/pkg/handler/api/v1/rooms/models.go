package rooms

import "metroid_bookmarks/pkg/repository/sql"

type createForm struct {
	*sql.CreateRoom
}

type createResponse struct {
	*sql.Room
}
type editForm struct {
	*sql.EditRoom
}
type editResponse struct {
	*sql.Room
}

type deleteResponse struct {
	*sql.Room
}

type getAllResponse struct {
	Data  []sql.Room `json:"data"`
	Total int        `json:"total"`
}
