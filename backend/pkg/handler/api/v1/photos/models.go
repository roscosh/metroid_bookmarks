package photos

import "metroid_bookmarks/pkg/repository/sql"

type createForm struct {
	BookmarkId int `form:"bookmark_id" binding:"required"`
}

type createResponse struct {
	*sql.PhotoPreview
}

type deleteResponse struct {
	*sql.PhotoPreview
}
