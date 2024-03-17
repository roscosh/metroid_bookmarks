package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/misc/session"
	base2 "metroid_bookmarks/pkg/handler/api/base_api"
	"net/http"
)

// @Summary me
// @Security HeaderAuth
// @Tags auth
// @Accept json
// @Success 200 {object} ResponseMe
// @Failure 401,404 {object} ErrorResponse
// @Router /auth/me [get]
func (h *AuthRouter) me(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(base2.UserCtx).(*session.Session)

	c.JSON(http.StatusOK, ResponseMe{Session: sessionObj})
}

// @Summary login
// @Security HeaderAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param input body FormLogin true "login"
// @Success 200 {object} ResponseLogin
// @Failure 401,404 {object} ErrorResponse
// @Router /auth/login [post]
func (h *AuthRouter) login(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(base2.UserCtx).(*session.Session)

	if sessionObj.IsAuthenticated() {
		base2.Response401(c, errors.New("You are already authorized!"))
		return
	}

	var form FormLogin
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		base2.Response404(c, err)
		return
	}
	sessionObj, err = h.authService.Login(form.Login, form.Password, sessionObj)
	if err != nil {
		base2.Response404(c, err)
		return
	}
	c.Header(session.HeadersSessionName, sessionObj.Token)
	c.JSON(http.StatusOK, ResponseLogin{Session: sessionObj})
}

// @Summary logout
// @Security HeaderAuth
// @Tags auth
// @Accept json
// @Success 200 {object} ResponseLogout
// @Failure 401,404 {object} ErrorResponse
// @Router /auth/logout [post]
func (h *AuthRouter) logout(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(base2.UserCtx).(*session.Session)

	sessionObj = h.authService.Logout(sessionObj)
	c.Header(session.HeadersSessionName, sessionObj.Token)
	c.JSON(http.StatusOK, ResponseLogout{Session: sessionObj})
}

// @Summary signUp (только для разработки)
// @Security HeaderAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param input body FormCreateUser true "signUp"
// @Success 200 {object} ResponseCreateUser
// @Failure 404 {object} ErrorResponse
// @Router /auth/sign_up [post]
func (h *AuthRouter) signUp(c *gin.Context) {
	var form FormCreateUser
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		base2.Response404(c, err)
		return
	}
	user, err := h.usersService.Create(form.CreateUser)
	if err != nil {
		base2.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseCreateUser{User: user})
}
