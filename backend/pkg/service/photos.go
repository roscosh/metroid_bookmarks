package service

import (
	"metroid_bookmarks/pkg/repository/sql"
)

type PhotosService struct {
	sql *sql.PhotosSQL
}

func newPhotosService(sql *sql.PhotosSQL) *PhotosService {
	return &PhotosService{sql: sql}
}

func (s *PhotosService) Create(createForm *sql.CreatePhoto) (*sql.PhotoPreview, error) {
	photo, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return photo, nil
}

func (s *PhotosService) Delete(id int) (*sql.PhotoPreview, error) {
	photo, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return photo, nil
}
