package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/internal/handler/api/middleware"
)

// @Summary login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body loginForm true "login"
// @Success 200 {object} loginResponse
// @Failure 401,404 {object} middleware.ErrorResponse
// @Router /auth/login [post]
func (h *Router) login(c *gin.Context) {
	session := middleware.GetSession(c)

	var form loginForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	session, err = h.service.Login(form.Login, form.Password, session)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.SetCookie(c, session)
	middleware.Response200(c, loginResponse{Session: session})
}

// @Summary logout
// @Tags auth
// @Accept json
// @Success 200 {object} logoutResponse
// @Failure 401,404 {object} middleware.ErrorResponse
// @Router /auth/logout [post]
func (h *Router) logout(c *gin.Context) {
	session := middleware.GetSession(c)

	session = h.service.Logout(session)
	middleware.SetCookie(c, session)
	middleware.Response200(c, logoutResponse{Session: session})
}

// @Summary me
// @Tags auth
// @Accept json
// @Success 200 {object} meResponse
// @Failure 401,404 {object} middleware.ErrorResponse
// @Router /auth/me [get]
func (h *Router) me(c *gin.Context) {
	session := middleware.GetSession(c)

	middleware.Response200(c, meResponse{Session: session})
}

// @Summary signUp
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signUpForm true "signUp"
// @Success 200 {object}  signUpResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /auth/sign_up [post]
func (h *Router) signUp(c *gin.Context) {
	var form signUpForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	user, err := h.service.SignUp(form.CreateUser)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, signUpResponse{User: user})
}
