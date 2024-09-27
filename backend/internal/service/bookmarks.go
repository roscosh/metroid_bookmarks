package service

import (
	"metroid_bookmarks/internal/repository/sql/bookmarks"
)

type BookmarksService struct {
	sql bookmarks.SQL
}

func newBookmarksService(sqlBookmarks bookmarks.SQL) *BookmarksService {
	return &BookmarksService{sql: sqlBookmarks}
}

func (s *BookmarksService) Create(createForm *bookmarks.CreateBookmark) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}

func (s *BookmarksService) Delete(id, userID int) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sql.Delete(id, userID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}

func (s *BookmarksService) Edit(bookmarkID, userID int, editForm *bookmarks.EditBookmark) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sql.Edit(bookmarkID, userID, editForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}

func (s *BookmarksService) GetAll(
	limit int,
	page int,
	userID int,
	completed *bool,
	orderByID *bool,
) ([]bookmarks.Bookmark, int, error) {
	offset := (page - 1) * limit

	data, err := s.sql.GetAll(limit, offset, userID, completed, orderByID)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}

	total, err := s.sql.Total(userID, completed)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}

	if data == nil {
		data = []bookmarks.Bookmark{}
	}

	return data, total, nil
}

func (s *BookmarksService) GetByID(id int) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sql.GetByID(id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}
