package photos

import "metroid_bookmarks/pkg/repository/sql"

type createForm struct {
	*sql.CreatePhoto
}

type createResponse struct {
	*sql.PhotoPreview
}

type deleteResponse struct {
	*sql.PhotoPreview
}
