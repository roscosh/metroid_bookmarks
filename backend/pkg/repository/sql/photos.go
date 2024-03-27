package sql

import "github.com/jackc/pgx/v5"

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
	*baseSQL
}

func NewPhotosSQL(pool *DbPool, table string) *PhotosSQL {
	sql := newBaseSQl(pool, table, PhotoPreview{})
	return &PhotosSQL{baseSQL: sql}
}

func (s *PhotosSQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	rows, err := s.insert(*createForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *PhotosSQL) Delete(id int, userId int) (*PhotoPreview, error) {
	query := `
	   DELETE
	   FROM photos p
	   USING bookmarks b
	   Where p.bookmark_id=b.id AND b.user_id=$1 AND p.id=$2
	   RETURNING p.id, p.bookmark_id, p.name
	`
	rows, err := s.baseSQL.query(query, userId, id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *PhotosSQL) collectOneRow(rows pgx.Rows) (*PhotoPreview, error) {
	return collectOneRow[PhotoPreview](rows)
}

func (s *PhotosSQL) collectRows(rows pgx.Rows) ([]PhotoPreview, error) {
	return collectRows[PhotoPreview](rows)
}
