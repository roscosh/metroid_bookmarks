package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/handler/api/v1/users"
	"net/http"
)

// @Summary me
// @Tags auth
// @Accept json
// @Success 200 {object} responseMe
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/me [get]
func (h *AuthRouter) me(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)

	c.JSON(http.StatusOK, responseMe{Session: sessionObj})
}

// @Summary login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body formLogin true "login"
// @Success 200 {object} responseLogin
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/login [post]
func (h *AuthRouter) login(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)

	if sessionObj.IsAuthenticated() {
		baseApi.Response401(c, errors.New("You are already authorized!"))
		return
	}

	var form formLogin
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	sessionObj, err = h.authService.Login(form.Login, form.Password, sessionObj)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.SetCookie(c, sessionObj)
	c.JSON(http.StatusOK, responseLogin{Session: sessionObj})
}

// @Summary logout
// @Tags auth
// @Accept json
// @Success 200 {object} responseLogout
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/logout [post]
func (h *AuthRouter) logout(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)

	sessionObj = h.authService.Logout(sessionObj)
	baseApi.SetCookie(c, sessionObj)
	c.JSON(http.StatusOK, responseLogout{Session: sessionObj})
}

// @Summary signUp (только для разработки)
// @Tags auth
// @Accept json
// @Produce json
// @Param input body  users.FormCreateUser true "signUp"
// @Success 200 {object}  users.ResponseCreateUser
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /auth/sign_up [post]
func (h *AuthRouter) signUp(c *gin.Context) {
	var form users.FormCreateUser
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.usersService.Create(form.CreateUser)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, users.ResponseCreateUser{User: user})
}
