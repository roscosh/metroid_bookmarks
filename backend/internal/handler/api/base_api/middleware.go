package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/session"
)

const (
	userCtx = "userId"
	photo   = "photo"
)

type Middleware struct {
	service *service.MiddlewareService
}

func NewMiddleware(service *service.MiddlewareService) *Middleware {
	return &Middleware{service: service}
}

func (h *Middleware) SessionRequired(c *gin.Context) {
	token, _ := c.Cookie(session.CookieSessionName)
	sessionObj, err := h.service.GetExistSession(token)
	if err != nil {
		sessionObj, err = h.service.CreateSession()
		if err != nil {
			Response404(c, err)
			c.Abort()
			return
		}
	}
	c.Set(userCtx, sessionObj)

	SetCookie(c, sessionObj)
	c.Next()

	h.service.UpdateSession(sessionObj)
}

func (h *Middleware) AdminRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAdmin() {
		Response403(c, errors.New("Нужны права администратора для этого запроса!"))
		c.Abort()
		return
	}
}

func (h *Middleware) AuthRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAuthenticated() {
		Response401(c, errors.New("Нужно залогиниться для этого запроса!"))
		c.Abort()
		return
	}
}
