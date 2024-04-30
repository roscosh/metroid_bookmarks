package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/repository/sql"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type PhotosService struct {
	sql *sql.PhotosSQL
}

func newPhotosService(sql *sql.PhotosSQL) *PhotosService {
	return &PhotosService{sql: sql}
}

func (s *PhotosService) Create(
	c *gin.Context,
	userId int,
	bookmarkId int,
	photoFile *multipart.FileHeader,
	photoRoot string,
	format string,
) (*sql.PhotoPreview, error) {
	var errMessage string
	var err error

	// Создание целевой директории
	saveDir := filepath.Join(photoRoot, strconv.Itoa(userId), strconv.Itoa(bookmarkId))
	if err = os.MkdirAll(saveDir, os.ModePerm); err != nil {
		errMessage = fmt.Sprintf("Ошибка при создании директории:%s", err)
		err = errors.New(errMessage)
		logger.Error(err.Error())
		return nil, err
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
				err = errors.New(errMessage)
				logger.Error(err.Error())
				return nil, err
			} else {
				success = true
				break
			}

		} else {
			filenameFormat = NewFilename + fmt.Sprintf("_%d", i) + "." + format
			path = filepath.Join(saveDir, filenameFormat)
		}
	}
	if !success {
		return nil, errors.New("file upload is overloaded, please try later")
	}
	createForm := sql.CreatePhoto{
		BookmarkId: bookmarkId,
		Name:       filenameFormat,
	}
	photo, err := s.sql.Create(&createForm)
	if err != nil {
		err = createPgError(err)
		logger.Error(err.Error())
		return nil, err
	}
	return photo, nil
}

func (s *PhotosService) Delete(id int, userId int) (*sql.PhotoPreview, error) {
	photo, err := s.sql.Delete(id, userId)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return photo, nil
}
