package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/repository/sql"
	"net/http"
)

type FormGetUsers struct {
	Search string `form:"search"`
}

type ResponseGetUsers struct {
	Data  []sql.User `json:"data"`
	Total int        `json:"total"`
}

type ResponseDeleteUser struct {
	*sql.User
}

type FormEditUser struct {
	sql.EditUser
}

type ResponseEditUser struct {
	*sql.User
}

type ResponseChangePassword struct {
	*sql.User
}

type FormChangePassword struct {
	sql.ChangePassword
}

// @Summary getAllUsers
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param q query FormGetUsers true "getAllUsers"
// @Success 200 {object} ResponseGetUsers
// @Failure 404 {object} ErrorResponse
// @Router /users/get_all [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	var form FormGetUsers
	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		h.Response404(c, err)
		return
	}
	users, total, err := h.services.Users.GetAllUsers(form.Search)
	if err != nil {
		h.Response404(c, err)
		return
	}
	c.JSON(
		http.StatusOK,
		ResponseGetUsers{
			Data:  users,
			Total: total,
		},
	)
}

// @Summary createUser
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param input body FormCreateUser true "createUser"
// @Success 200 {object} ResponseCreateUser
// @Failure 404 {object} ErrorResponse
// @Router /users/ [post]
func (h *Handler) createUser(c *gin.Context) {
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

// @Summary deleteUser
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} ResponseDeleteUser
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	id, err := h.getPathID(c)
	if err != nil {
		h.Response404(c, err)
		return
	}
	user, err := h.services.Users.Delete(id)
	if err != nil {
		h.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseDeleteUser{User: user})
}

// @Summary editUser
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body FormEditUser true "editUser"
// @Success 200 {object} ResponseEditUser
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) editUser(c *gin.Context) {
	id, err := h.getPathID(c)
	if err != nil {
		h.Response404(c, err)
		return
	}

	var form FormEditUser
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		h.Response404(c, err)
		return
	}
	user, err := h.services.Users.Edit(id, form.EditUser)
	if err != nil {
		h.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseEditUser{User: user})
}

// @Summary changePassword
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body FormChangePassword true "changePassword"
// @Success 200 {object} ResponseChangePassword
// @Failure 404 {object} ErrorResponse
// @Router /users/change_password/{id} [put]
func (h *Handler) changePassword(c *gin.Context) {
	id, err := h.getPathID(c)
	if err != nil {
		h.Response404(c, err)
		return
	}
	var form FormChangePassword
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		h.Response404(c, err)
		return
	}
	user, err := h.services.Users.ChangePassword(id, form.ChangePassword)
	if err != nil {
		h.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseChangePassword{User: user})
}
