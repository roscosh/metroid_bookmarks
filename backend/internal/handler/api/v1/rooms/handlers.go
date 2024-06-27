package rooms

import (
	"metroid_bookmarks/internal/handler/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary create
// @Tags rooms
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /rooms/ [post]
func (h *Router) create(c *gin.Context) {
	var form createForm

	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	room, err := h.service.Create(form.CreateRoom)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, createResponse{Room: room})
}

// @Summary edit
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /rooms/{id} [put]
func (h *Router) edit(c *gin.Context) {
	roomID, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	var form editForm

	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	room, err := h.service.Edit(roomID, form.EditRoom)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, editResponse{Room: room})
}

// @Summary delete
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /rooms/{id} [delete]
func (h *Router) delete(c *gin.Context) {
	roomID, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	room, err := h.service.Delete(roomID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, deleteResponse{Room: room})
}

// @Summary getAll
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /rooms/get_all [get]
func (h *Router) getAll(c *gin.Context) {
	room, total, err := h.service.GetAll()
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, getAllResponse{Data: room, Total: total})
}
