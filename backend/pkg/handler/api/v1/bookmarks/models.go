package bookmarks

import "metroid_bookmarks/pkg/repository/sql"

type createForm struct {
	AreaId  int `form:"area_id"  binding:"required"`
	RoomId  int `form:"room_id"  binding:"required"`
	SkillId int `form:"skill_id" binding:"required"`
}
type createResponse struct {
	*sql.BookmarkPreview
}

type deleteResponse struct {
	*sql.BookmarkPreview
}

type editForm struct {
	*sql.EditBookmark
}

type editResponse struct {
	*sql.BookmarkPreview
}

type getAllForm struct {
	Limit     int   `form:"limit"     binding:"required"`
	Page      int   `form:"page"      binding:"required"`
	OrderById *bool `form:"order_by_id"`
	Completed *bool `form:"completed"`
}

type getAllResponse struct {
	Data  []sql.Bookmark `json:"data"`
	Total int            `json:"total"`
}
