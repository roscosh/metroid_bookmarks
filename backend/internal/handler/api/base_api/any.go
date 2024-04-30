package baseApi

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"image/jpeg"
	"image/png"
	"metroid_bookmarks/pkg/session"
	"mime/multipart"
	"strconv"
	"strings"
)

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
}

func GetPathUserID(c *gin.Context) (int, error) {
	idStr := c.Param("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("user_id должен быть числом!")
	}
	return id, err
}

func GetPathBookmarkID(c *gin.Context) (int, error) {
	idStr := c.Param("bookmark_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("bookmark_id должен быть числом!")
	}
	return id, err
}

func GetPathName(c *gin.Context) (string, error) {
	name := c.Param("name")
	if name == "" {
		return "", errors.New("name не должен быть пустым!")
	}
	return name, nil
}

func GetSession(c *gin.Context) *session.Session {
	return c.MustGet(userCtx).(*session.Session)
}

func SetCookie(c *gin.Context, sessionObj *session.Session) {
	c.SetCookie(session.CookieSessionName, sessionObj.Token, sessionObj.Expires, "", "", false, false)
}

func GetPhoto(c *gin.Context) (*multipart.FileHeader, error) {
	return c.FormFile(photo)
}

func ValidatePhoto(photoFile *multipart.FileHeader) (string, error) {
	var errMessage string
	filename := photoFile.Filename
	// Проверка расширения файла
	if !strings.HasSuffix(filename, ".jpg") &&
		!strings.HasSuffix(filename, ".jpeg") &&
		!strings.HasSuffix(filename, ".png") {
		return "", errors.New("файл не является изображением")
	}
	// Попытка прочитать изображение
	file, err := photoFile.Open()
	if err != nil {
		errMessage = fmt.Sprintf("Ошибка при открытии файла:%s", err)
		err = errors.New(errMessage)
		logger.Error(err.Error())
		return "", err
	}
	defer file.Close()

	// Попытка декодировать изображение
	format := "jpeg"
	_, err = jpeg.Decode(file)
	if err != nil {
		file.Seek(0, 0)
		format = "png"
		_, err = png.DecodeConfig(file)
		if err != nil {
			errMessage = fmt.Sprintf("Ошибка при декодировании изображения:%s", err)
			err = errors.New(errMessage)
			logger.Error(err.Error())
			return "", err
		}
	}
	return format, nil
}
