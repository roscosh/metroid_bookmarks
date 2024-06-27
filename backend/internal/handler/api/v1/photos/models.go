package photos

import (
	"errors"
	"metroid_bookmarks/internal/repository/sql/photos"
)

var ErrFileDoesNotExist = errors.New("файл не существует")

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

type createForm struct {
	BookmarkID int `binding:"required" form:"bookmark_id"`
}

type createResponse struct {
	*photos.PhotoPreview
}

type deleteResponse struct {
	*photos.PhotoPreview
}
