package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"net/http"
)

// @Summary getAllUsers
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param q query formGetUsers true "getAllUsers"
// @Success 200 {object} responseGetUsers
// @Failure 404 {object} baseApi.ErrorResponse
func (h *UsersRouter) getAllUsers(c *gin.Context) {
	var form formGetUsers
	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	users, total, err := h.usersService.GetAllUsers(form.Search)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	c.JSON(
		http.StatusOK,
		responseGetUsers{
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
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/ [post]
func (h *UsersRouter) createUser(c *gin.Context) {
	var form FormCreateUser
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
	c.JSON(http.StatusOK, ResponseCreateUser{User: user})
}

// @Summary deleteUser
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} responseDeleteUser
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/{id} [delete]
func (h *UsersRouter) deleteUser(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.usersService.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, responseDeleteUser{User: user})
}

// @Summary editUser
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body formEditUser true "editUser"
// @Success 200 {object} responseEditUser
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/{id} [put]
func (h *UsersRouter) editUser(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}

	var form formEditUser
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.usersService.Edit(id, form.EditUser)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, responseEditUser{User: user})
}

// @Summary changePassword
// @Security HeaderAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body formChangePassword true "changePassword"
// @Success 200 {object} responseChangePassword
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/change_password/{id} [put]
func (h *UsersRouter) changePassword(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	var form formChangePassword
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.usersService.ChangePassword(id, form.ChangePassword)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	c.JSON(http.StatusOK, responseChangePassword{User: user})
}
