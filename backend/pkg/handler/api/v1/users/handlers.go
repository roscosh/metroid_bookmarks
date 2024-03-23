package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/handler/api/base_api"
)

// @Summary changePassword
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body changePasswordForm true "changePassword"
// @Success 200 {object} changePasswordResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/change_password/{id} [put]
func (h *router) changePassword(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	var form changePasswordForm
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.service.ChangePassword(id, form.ChangePassword)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, changePasswordResponse{User: user})
}

// @Summary delete
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/{id} [delete]
func (h *router) delete(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.service.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{User: user})
}

// @Summary edit
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object} editResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/{id} [put]
func (h *router) edit(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}

	var form editForm
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	user, err := h.service.Edit(id, form.EditUser)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, editResponse{User: user})
}

// @Summary getAll
// @Tags users
// @Accept json
// @Produce json
// @Param q query getAllForm true "getAll"
// @Success 200 {object} getAllResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /users/get_all [get]
func (h *router) getAll(c *gin.Context) {
	var form getAllForm
	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	users, total, err := h.service.GetAll(form.Search)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, getAllResponse{Data: users, Total: total})
}
