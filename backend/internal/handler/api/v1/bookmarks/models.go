package bookmarks

import (
	"metroid_bookmarks/internal/repository/sql/bookmarks"
)

type createForm struct {
	AreaID  int `binding:"required" form:"area_id"`
	RoomID  int `binding:"required" form:"room_id"`
	SkillID int `binding:"required" form:"skill_id"`
}
type createResponse struct {
	*bookmarks.BookmarkPreview
}

type deleteResponse struct {
	*bookmarks.BookmarkPreview
}

type editForm struct {
	*bookmarks.EditBookmark
}

type editResponse struct {
	*bookmarks.BookmarkPreview
}

type getAllForm struct {
	Limit     int   `binding:"required" form:"limit"`
	Page      int   `binding:"required" form:"page"`
	OrderByID *bool `form:"order_by_id"`
	Completed *bool `form:"completed"`
}

type getAllResponse struct {
	Data  []bookmarks.Bookmark `json:"data"`
	Total int                  `json:"total"`
}
