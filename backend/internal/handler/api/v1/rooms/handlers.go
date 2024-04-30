package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/internal/handler/api/base_api"
)

// @Summary create
// @Tags rooms
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /rooms/ [post]
func (h *router) create(c *gin.Context) {
	var form createForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	room, err := h.service.Create(form.CreateRoom)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, createResponse{Room: room})
}

// @Summary edit
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /rooms/{id} [put]
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
	room, err := h.service.Edit(id, form.EditRoom)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, editResponse{Room: room})
}

// @Summary delete
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /rooms/{id} [delete]
func (h *router) delete(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	room, err := h.service.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{Room: room})
}

// @Summary getAll
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /rooms/get_all [get]
func (h *router) getAll(c *gin.Context) {
	room, total, err := h.service.GetAll()
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, getAllResponse{Data: room, Total: total})
}
