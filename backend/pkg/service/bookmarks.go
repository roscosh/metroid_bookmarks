package service

import (
	"errors"
	"metroid_bookmarks/pkg/repository/sql"
)

type BookmarksService struct {
	sqlBookmarks *sql.BookmarksSQL
	sqlPhotos    *sql.PhotosSQL
}

func newBookmarksService(sqlBookmarks *sql.BookmarksSQL, sqlPhotos *sql.PhotosSQL) *BookmarksService {
	return &BookmarksService{sqlBookmarks: sqlBookmarks, sqlPhotos: sqlPhotos}
}

func (s *BookmarksService) Create(createForm *sql.CreateBookmark) (*sql.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarksService) Delete(id int, userId int) (*sql.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.Delete(id, userId)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarksService) Edit(id int, userId int, editForm *sql.EditBookmark) (*sql.BookmarkPreview, error) {
	if (editForm == &sql.EditBookmark{}) {
		return nil, errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	bookmark, err := s.sqlBookmarks.Edit(id, userId, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarksService) GetAll(limit int, page int, userId int, completed *bool, orderById *bool) ([]sql.Bookmark, int, error) {
	offset := (page - 1) * limit
	data, err := s.sqlBookmarks.GetAll(limit, offset, userId, completed, orderById)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}

	total, err := s.sqlBookmarks.Total(userId, completed)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	return data, total, nil
}

func (s *BookmarksService) GetByID(id int) (*sql.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.GetByID(id)
	if err != nil {
		logger.Error(err.Error())
		err = selectPgError(err, id)
	}
	return bookmark, err
}
