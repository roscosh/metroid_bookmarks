package sql

const photosTable = "photos"

type Photo struct {
	Id   int    `json:"id"      db:"id"`
	Path string `json:"path" db:"path"`
}

type PhotoPreview struct {
	Id         int    `json:"id"      db:"id"`
	BookmarkId int    `json:"bookmark_id" db:"bookmark_id"`
	Path       string `json:"path" db:"path"`
}

// todo переделать, не должно быть path
type CreatePhoto struct {
	BookmarkId int    `json:"bookmark_id" db:"bookmark_id" binding:"required"`
	Path       string `json:"path" db:"path" binding:"required"`
}

type PhotosSQL struct {
	*baseSQL
}

func NewPhotosSQL(baseSQL *baseSQL) *PhotosSQL {
	return &PhotosSQL{baseSQL: baseSQL}
}

func (s *PhotosSQL) GetByID(id int) (*Photo, error) {
	return selectById[Photo](s.baseSQL, photosTable, id)
}

func (s *PhotosSQL) GetAll() ([]Photo, error) {
	return selectAll[Photo](s.baseSQL, photosTable)
}

func (s *PhotosSQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	return insert[PhotoPreview](s.baseSQL, photosTable, *createForm)
}

func (s *PhotosSQL) Delete(id int) (*PhotoPreview, error) {
	return deleteById[PhotoPreview](s.baseSQL, photosTable, id)
}
