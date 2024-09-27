package photos

import (
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const photosTable = "photos"

type SQL interface {
	Create(createForm *CreatePhoto) (*PhotoPreview, error)
	Delete(photoID, userID int) (*PhotoPreview, error)
}

type photoSQL struct {
	sql pgpool.SQL[PhotoPreview]
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[PhotoPreview](dbPool, photosTable)
	return &photoSQL{sql: sql}
}

func (s *photoSQL) Create(createForm *CreatePhoto) (*PhotoPreview, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *photoSQL) Delete(photoID, userID int) (*PhotoPreview, error) {
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

	entity, err := s.sql.CollectOneRow(rows)
	if err != nil {
		err = pgerr.DeletePgError(err, photoID)
		return nil, err
	}

	return entity, nil
}
