package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/handler/api/base_api"
)

// @Summary login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body loginForm true "login"
// @Success 200 {object} loginResponse
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/login [post]
func (h *router) login(c *gin.Context) {
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)
	if sessionObj.IsAuthenticated() {
		baseApi.Response401(c, errors.New("You are already authorized!"))
		return
	}

	var form loginForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	sessionObj, err = h.service.Login(form.Login, form.Password, sessionObj)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.SetCookie(c, sessionObj)
	baseApi.Response200(c, loginResponse{Session: sessionObj})
}

// @Summary logout
// @Tags auth
// @Accept json
// @Success 200 {object} logoutResponse
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/logout [post]
func (h *router) logout(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)

	sessionObj = h.service.Logout(sessionObj)
	baseApi.SetCookie(c, sessionObj)
	baseApi.Response200(c, logoutResponse{Session: sessionObj})
}

// @Summary me
// @Tags auth
// @Accept json
// @Success 200 {object} meResponse
// @Failure 401,404 {object} baseApi.ErrorResponse
// @Router /auth/me [get]
func (h *router) me(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(baseApi.UserCtx).(*session.Session)

	baseApi.Response200(c, meResponse{Session: sessionObj})
}

// @Summary signUp (только для разработки)
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signUpForm true "signUp"
// @Success 200 {object}  signUpResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /auth/sign_up [post]
func (h *router) signUp(c *gin.Context) {
	var form signUpForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.service.SignUp(form.CreateUser)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, signUpResponse{User: user})
}
