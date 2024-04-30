package areas

import "metroid_bookmarks/internal/repository/sql"

type createForm struct {
	*sql.CreateArea
}

type createResponse struct {
	*sql.Area
}
type editForm struct {
	*sql.EditArea
}
type editResponse struct {
	*sql.Area
}

type deleteResponse struct {
	*sql.Area
}

type getAllResponse struct {
	Data  []sql.Area `json:"data"`
	Total int        `json:"total"`
}
