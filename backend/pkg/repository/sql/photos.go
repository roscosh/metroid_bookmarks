package sql

const photosTable = "photos"

type Photo struct {
	Id   int    `json:"id"      db:"id"`
	Path string `json:"path" db:"path"`
}

type PhotoWithUser struct {
	Id   int    `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

type PhotoPreview struct {
	Id         int    `json:"id"          db:"id"`
	BookmarkId int    `json:"bookmark_id" db:"bookmark_id"`
	Name       string `json:"name" db:"name"`
}

type CreatePhoto struct {
	BookmarkId int    `db:"bookmark_id"`
	Name       string `db:"name"`
}

type PhotosSQL struct {
	*baseSQL
}

func NewPhotosSQL(baseSQL *baseSQL) *PhotosSQL {
	return &PhotosSQL{baseSQL: baseSQL}
}

func (s *PhotosSQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	return insert[PhotoPreview](s.baseSQL, photosTable, *createForm)
}

func (s *PhotosSQL) Delete(id int) (*PhotoPreview, error) {
	return deleteById[PhotoPreview](s.baseSQL, photosTable, id)
}
