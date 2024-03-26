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
// @Param form formData createForm true "create"
// @Param photo formData file true "photo"
// @Success 200 {object}  createResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /photos/ [post]
func (h *router) create(c *gin.Context) {
	session := baseApi.GetSession(c)
	var form createForm
	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	file, err := baseApi.GetPhoto(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	bookmark, err := h.bookmarksService.GetByID(form.BookmarkId)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	if bookmark.UserId != session.ID {
		baseApi.AccessDenied(c)
		return
	}
	photo, err := h.photosService.Create(session.ID, form.BookmarkId, file, h.Config.PhotosPath, c)
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
	photo, err := h.photosService.Delete(id)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{PhotoPreview: photo})
}
