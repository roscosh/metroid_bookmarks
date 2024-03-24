package service

import (
	"errors"
	"metroid_bookmarks/pkg/repository/sql"
)

type BookmarksService struct {
	sql *sql.BookmarksSQL
}

func newBookmarksService(sql *sql.BookmarksSQL) *BookmarksService {
	return &BookmarksService{sql: sql}
}

func (s *BookmarksService) Create(createForm *sql.CreateBookmark) (*sql.BookmarkPreview, error) {
	bookmark, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarksService) Delete(id int, userId int) (*sql.BookmarkPreview, error) {
	bookmark, err := s.sql.Delete(id, userId)
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
	bookmark, err := s.sql.Edit(id, userId, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarksService) GetAll(limit int, page int, userId int, completed *bool, orderById *bool) ([]sql.Bookmark, int, error) {
	offset := (page - 1) * limit
	data, err := s.sql.GetAll(limit, offset, userId, completed, orderById)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	total, err := s.sql.Total(userId, completed)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	return data, total, nil
}
