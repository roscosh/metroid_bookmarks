package areas

import (
	"metroid_bookmarks/internal/repository/sql/areas"
)

type createForm struct {
	*areas.CreateArea
}

type createResponse struct {
	*areas.Area
}
type editForm struct {
	*areas.EditArea
}
type editResponse struct {
	*areas.Area
}

type deleteResponse struct {
	*areas.Area
}

type getAllResponse struct {
	Data  []areas.Area `json:"data"`
	Total int          `json:"total"`
}
