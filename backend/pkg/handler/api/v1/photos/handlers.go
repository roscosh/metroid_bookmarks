package photos

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/handler/api/base_api"
)

// @Summary create
// @Tags photos
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /photos/ [post]
func (h *router) create(c *gin.Context) {
	var form createForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	photo, err := h.service.Create(form.CreatePhoto)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, createResponse{PhotoPreview: photo})
}

// @Summary delete
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /photos/{id} [delete]
func (h *router) delete(c *gin.Context) {
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	photo, err := h.service.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{PhotoPreview: photo})
}
