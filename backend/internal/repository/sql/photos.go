package sql

const photosTable = "photos"

type Photo struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type PhotoPreview struct {
	Id         int    `json:"id"          db:"id"`
	BookmarkId int    `json:"bookmark_id" db:"bookmark_id"`
	Name       string `json:"name"        db:"name"`
}

type CreatePhoto struct {
	BookmarkId int    `db:"bookmark_id"`
	Name       string `db:"name"`
}

type PhotosSQL struct {
	iBaseSQL[PhotoPreview]
}

func NewPhotosSQL(dbPool *DbPool) *PhotosSQL {
	sql := newIBaseSQL[PhotoPreview](dbPool, photosTable)
	return &PhotosSQL{iBaseSQL: sql}
}

func (s *PhotosSQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	return s.insert(createForm)
}

func (s *PhotosSQL) Delete(id int, userId int) (*PhotoPreview, error) {
	query := `
	   DELETE
	   FROM photos p
	   USING bookmarks b
	   Where p.bookmark_id=b.id AND b.user_id=$1 AND p.id=$2
	   RETURNING p.id, p.bookmark_id, p.name
	`
	rows, err := s.iBaseSQL.query(query, userId, id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}
