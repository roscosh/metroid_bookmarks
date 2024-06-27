package areas

import (
	"metroid_bookmarks/internal/handler/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary create
// @Tags areas
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /areas/ [post]
func (h *Router) create(c *gin.Context) {
	var form createForm

	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	area, err := h.service.Create(form.CreateArea)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, createResponse{Area: area})
}

// @Summary edit
// @Tags areas
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /areas/{id} [put]
func (h *Router) edit(c *gin.Context) {
	areaID, err := middleware.GetPathID(c)
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

	area, err := h.service.Edit(areaID, form.EditArea)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, editResponse{Area: area})
}

// @Summary delete
// @Tags areas
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /areas/{id} [delete]
func (h *Router) delete(c *gin.Context) {
	areaID, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	area, err := h.service.Delete(areaID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, deleteResponse{Area: area})
}

// @Summary getAll
// @Tags areas
// @Accept json
// @Produce json
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /areas/get_all [get]
func (h *Router) getAll(c *gin.Context) {
	area, total, err := h.service.GetAll()
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, getAllResponse{Data: area, Total: total})
}
