package middleware

import (
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/session"

	"github.com/gin-gonic/gin"
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

func (m *Middleware) SessionRequired(c *gin.Context) {
	token, _ := c.Cookie(session.CookieSessionName)

	sessionObj, err := m.service.GetExistSession(token)
	if err != nil {
		sessionObj, err = m.service.CreateSession()
		if err != nil {
			Response404(c, err)
			c.Abort()

			return
		}
	}

	c.Set(userCtx, sessionObj)

	SetCookie(c, sessionObj)
	c.Next()

	m.service.UpdateSession(sessionObj)
}

func (m *Middleware) AdminRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAdmin() {
		Response403(c, ErrAdminRequired)
		c.Abort()

		return
	}
}

func (m *Middleware) AuthRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAuthenticated() {
		Response401(c, ErrLoginRequired)
		c.Abort()

		return
	}
}

func (m *Middleware) LogoutRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if sessionObj.IsAuthenticated() {
		Response401(c, ErrAlreadyAuthorized)
		c.Abort()

		return
	}
}
