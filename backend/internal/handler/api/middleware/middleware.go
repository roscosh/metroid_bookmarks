package middleware

import (
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	userIDQueryKey   = "userId"
	photoQueryKey    = "photo"
	SessionCookieKey = "X-Session"
)

type Middleware struct {
	service *service.MiddlewareService
}

func NewMiddleware(service *service.MiddlewareService) *Middleware {
	return &Middleware{service: service}
}

func (m *Middleware) SessionRequired(c *gin.Context) {
	token, _ := c.Cookie(SessionCookieKey)

	sessionObj, err := m.service.GetExistSession(token)
	if err != nil {
		sessionObj, err = m.service.CreateSession()
		if err != nil {
			Response404(c, err)
			c.Abort()

			return
		}
	}

	c.Set(userIDQueryKey, sessionObj)

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
