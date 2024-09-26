package service

import (
	"fmt"
	"metroid_bookmarks/internal/repository/sql/photos"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PhotosService struct {
	sql *photos.SQL
}

func newPhotosService(sql *photos.SQL) *PhotosService {
	return &PhotosService{sql: sql}
}

func (s *PhotosService) Create(
	c *gin.Context,
	userID int,
	bookmarkID int,
	photoFile *multipart.FileHeader,
	photoRoot string,
	format string,
) (*photos.PhotoPreview, error) {
	var (
		errMessage string
		err        error
	)

	// Создание целевой директории
	saveDir := filepath.Join(photoRoot, strconv.Itoa(userID), strconv.Itoa(bookmarkID))
	if err = os.MkdirAll(saveDir, os.ModePerm); err != nil {
		errMessage = fmt.Sprintf("Ошибка при создании директории:%s", err)
		logger.Error(errMessage)

		return nil, NewErr(errMessage)
	}

	NewFilename := time.Now().UTC().Format("20060102_150405")
	filenameFormat := NewFilename + "." + format
	path := filepath.Join(saveDir, filenameFormat)

	var success bool

	for i := 1; i < 100; i++ {
		if _, err = os.Stat(path); os.IsNotExist(err) {
			err = c.SaveUploadedFile(photoFile, path)
			if err != nil {
				errMessage = fmt.Sprintf("Ошибка при сохранении файла:%s", err)
				logger.Error(errMessage)

				return nil, NewErr(errMessage)
			}

			success = true

			break
		}

		filenameFormat = NewFilename + fmt.Sprintf("_%d", i) + "." + format
		path = filepath.Join(saveDir, filenameFormat)
	}

	if !success {
		return nil, ErrFileUploadOverload
	}

	createForm := photos.CreatePhoto{
		BookmarkID: bookmarkID,
		Name:       filenameFormat,
	}

	photo, err := s.sql.Create(&createForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return photo, nil
}

func (s *PhotosService) Delete(photoID, userID int) (*photos.PhotoPreview, error) {
	photo, err := s.sql.Delete(photoID, userID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return photo, nil
}
