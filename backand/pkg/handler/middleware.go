package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/misc/session"
	"net/http"
	"strconv"
)

const (
	userCtx = "userId"
)

func (h *Handler) getSession(c *gin.Context) {
	token := c.GetHeader(session.HeadersSessionName)
	sessionObj, err := h.services.Middleware.GetExistSession(token)
	if err != nil {
		sessionObj, err = h.services.Middleware.CreateSession()
		if err != nil {
			h.Response404(c, err)
			c.Abort()
			return
		}
	}
	c.Set(userCtx, sessionObj)
	c.Header(session.HeadersSessionName, sessionObj.Token)

	c.Next()

	h.services.Middleware.UpdateSession(sessionObj)
}

func (h *Handler) adminRequired(c *gin.Context) {
	sessionObj := c.MustGet(userCtx).(*session.Session)

	if !sessionObj.IsAdmin() {
		h.Response403(c, errors.New("Нужны права администратора для этого запроса!"))
		c.Abort()
		return
	}
}

func (h *Handler) authRequired(c *gin.Context) {
	sessionObj := c.MustGet(userCtx).(*session.Session)

	if !sessionObj.IsAuthenticated() {
		h.Response401(c, errors.New("Нужно залогиниться для этого запроса!"))
		c.Abort()
		return
	}
}

func (h *Handler) Response401(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, NewErrorResponse(err))
}

func (h *Handler) Response403(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, NewErrorResponse(err))
}

func (h *Handler) Response404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, NewErrorResponse(err))
}

func (h *Handler) getPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
}
