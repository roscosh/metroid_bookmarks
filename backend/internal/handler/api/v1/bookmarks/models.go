package bookmarks

import (
	"metroid_bookmarks/internal/repository/sql/bookmarks"
)

type createForm struct {
	AreaId  int `form:"area_id"  binding:"required"`
	RoomId  int `form:"room_id"  binding:"required"`
	SkillId int `form:"skill_id" binding:"required"`
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
	Limit     int   `form:"limit"     binding:"required"`
	Page      int   `form:"page"      binding:"required"`
	OrderById *bool `form:"order_by_id"`
	Completed *bool `form:"completed"`
}

type getAllResponse struct {
	Data  []bookmarks.Bookmark `json:"data"`
	Total int                  `json:"total"`
}
