package rooms

import (
	"metroid_bookmarks/internal/repository/sql/rooms"
)

type createForm struct {
	*rooms.CreateRoom
}

type createResponse struct {
	*rooms.Room
}
type editForm struct {
	*rooms.EditRoom
}
type editResponse struct {
	*rooms.Room
}

type deleteResponse struct {
	*rooms.Room
}

type getAllResponse struct {
	Data  []rooms.Room `json:"data"`
	Total int          `json:"total"`
}
