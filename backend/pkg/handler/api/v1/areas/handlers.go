package areas

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/handler/api/base_api"
)

// @Summary create
// @Tags areas
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /areas/ [post]
func (h *router) create(c *gin.Context) {
	var form createForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	area, err := h.service.Create(form.CreateArea)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, createResponse{Area: area})
}

// @Summary edit
// @Tags areas
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /areas/{id} [put]
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
	area, err := h.service.Edit(id, form.EditArea)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, editResponse{Area: area})
}

// @Summary delete
// @Tags areas
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /areas/{id} [delete]
func (h *router) delete(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	area, err := h.service.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{Area: area})
}

// @Summary getAll
// @Tags areas
// @Accept json
// @Produce json
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /areas/get_all [get]
func (h *router) getAll(c *gin.Context) {
	area, total, err := h.service.GetAll()
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, getAllResponse{Data: area, Total: total})
}
