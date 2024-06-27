package middleware

import (
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"metroid_bookmarks/pkg/session"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrIDType            = errors.New("id должен быть числом")
	ErrUserIDType        = errors.New("user_id должен быть числом")
	ErrBookmarkIDType    = errors.New("bookmark_id должен быть числом")
	ErrEmptyName         = errors.New("name не должен быть пустым")
	ErrAccessDenied      = errors.New("отказ в доступе")
	ErrAlreadyAuthorized = errors.New("you are already authorized")
	ErrLoginRequired     = errors.New("нужно залогиниться для этого запроса")
	ErrAdminRequired     = errors.New("нужны права администратора для этого запроса")
	ErrFileNotImage      = errors.New("файл не является изображением")
	ErrOpenFile          = errors.New("ошибка при открытии файла")
	ErrSeekImage         = errors.New("ошибка при смещении координат")
	ErrDecodeImage       = errors.New("ошибка при декодировании изображения")
)

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrIDType
	}

	return id, nil
}

func GetPathUserID(c *gin.Context) (int, error) {
	idStr := c.Param("user_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrUserIDType
	}

	return id, nil
}

func GetPathBookmarkID(c *gin.Context) (int, error) {
	idStr := c.Param("bookmark_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrBookmarkIDType
	}

	return id, nil
}

func GetPathName(c *gin.Context) (string, error) {
	name := c.Param("name")
	if name == "" {
		return "", ErrEmptyName
	}

	return name, nil
}

func GetSession(c *gin.Context) *session.Session {
	return c.MustGet(userCtx).(*session.Session) //nolint:forcetypeassert
}

func SetCookie(c *gin.Context, sessionObj *session.Session) {
	c.SetCookie(session.CookieSessionName, sessionObj.Token, sessionObj.Expires, "", "", false, false)
}

func GetPhoto(c *gin.Context) (*multipart.FileHeader, error) {
	return c.FormFile(photo)
}

func ValidatePhoto(photoFile *multipart.FileHeader) (string, error) {
	filename := photoFile.Filename
	// Проверка расширения файла
	if !strings.HasSuffix(filename, ".jpg") &&
		!strings.HasSuffix(filename, ".jpeg") &&
		!strings.HasSuffix(filename, ".png") {
		return "", ErrFileNotImage
	}
	// Попытка прочитать изображение
	file, err := photoFile.Open()
	if err != nil {
		err = fmt.Errorf("%w:%w", ErrOpenFile, err)
		return "", err
	}
	defer file.Close()

	// Попытка декодировать изображение
	format := "jpeg"

	_, err = jpeg.Decode(file)
	if err != nil {
		_, err = file.Seek(0, 0)
		if err != nil {
			err = fmt.Errorf("%w:%w", ErrSeekImage, err)
			logger.Error(err.Error())

			return "", err
		}

		format = "png"

		_, err = png.DecodeConfig(file)
		if err != nil {
			err = fmt.Errorf("%w:%w", ErrDecodeImage, err)
			logger.Error(err.Error())

			return "", err
		}
	}

	return format, nil
}
