package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/repository/sql"
	"net/http"
)

type ResponseMe struct {
	*session.Session
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}

type FormLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type ResponseLogin struct {
	*session.Session
}

type ResponseLogout struct {
	*session.Session
}

type FormCreateUser struct {
	*sql.CreateUser
}

type ResponseCreateUser struct {
	*sql.User
}

// @Summary me
// @Security HeaderAuth
// @Tags auth
// @Accept json
// @Success 200 {object} ResponseMe
// @Failure 401,404 {object} ErrorResponse
// @Router /auth/me [get]
func (h *Handler) me(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(userCtx).(*session.Session)

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
func (h *Handler) login(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(userCtx).(*session.Session)

	if sessionObj.IsAuthenticated() {
		h.Response401(c, errors.New("You are already authorized!"))
		return
	}

	var form FormLogin
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		h.Response404(c, err)
		return
	}
	sessionObj, err = h.services.Authorization.Login(form.Login, form.Password, sessionObj)
	if err != nil {
		h.Response404(c, err)
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
func (h *Handler) logout(c *gin.Context) {
	//Depends
	sessionObj := c.MustGet(userCtx).(*session.Session)

	sessionObj = h.services.Authorization.Logout(sessionObj)
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
func (h *Handler) signUp(c *gin.Context) {
	var form FormCreateUser
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		h.Response404(c, err)
		return
	}
	user, err := h.services.Users.Create(form.CreateUser)
	if err != nil {
		h.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseCreateUser{User: user})
}
