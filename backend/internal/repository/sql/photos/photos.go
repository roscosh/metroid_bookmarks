package photos

import (
	"metroid_bookmarks/pkg/pgpool"
)

const photosTable = "photos"

type SQL struct {
	sql pgpool.SQL[PhotoPreview]
}

func NewSQL(dbPool *pgpool.PgPool) *SQL {
	sql := pgpool.NewSQL[PhotoPreview](dbPool, photosTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Delete(photoID, userID int) (*PhotoPreview, error) {
	query := `
	   DELETE
	   FROM photos p
	   USING bookmarks b
	   Where p.bookmark_id=b.id AND b.user_id=$1 AND p.id=$2
	   RETURNING p.id, p.bookmark_id, p.name
	`

	rows, err := s.sql.Query(query, userID, photoID)
	if err != nil {
		return nil, err
	}

	return s.sql.CollectOneRow(rows)
}
