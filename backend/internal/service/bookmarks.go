package service

import (
	"metroid_bookmarks/internal/repository/sql/bookmarks"
)

type BookmarksService struct {
	sqlBookmarks *bookmarks.SQL
}

func newBookmarksService(sqlBookmarks *bookmarks.SQL) *BookmarksService {
	return &BookmarksService{sqlBookmarks: sqlBookmarks}
}

func (s *BookmarksService) Create(createForm *bookmarks.CreateBookmark) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}

func (s *BookmarksService) Delete(id, userID int) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.Delete(id, userID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}

func (s *BookmarksService) Edit(bookmarkID, userID int, editForm *bookmarks.EditBookmark) (*bookmarks.BookmarkPreview, error) {
	bookmark, err := s.sqlBookmarks.Edit(bookmarkID, userID, editForm)
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

	data, err := s.sqlBookmarks.GetAll(limit, offset, userID, completed, orderByID)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}

	total, err := s.sqlBookmarks.Total(userID, completed)
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
	bookmark, err := s.sqlBookmarks.GetByID(id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return bookmark, nil
}
