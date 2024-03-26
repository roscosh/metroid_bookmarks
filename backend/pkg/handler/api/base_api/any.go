package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/misc/session"
	"mime/multipart"
	"strconv"
)

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
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
