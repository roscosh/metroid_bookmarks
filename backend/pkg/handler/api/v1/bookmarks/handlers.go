package bookmarks

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/repository/sql"
)

// @Summary create
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param input formData createForm true "create"
// @Param photo formData file true "photo"
// @Success 200 {object} createResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /bookmarks/ [post]
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
	format, err := baseApi.ValidatePhoto(file)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	sqlForm := &sql.CreateBookmark{
		UserId:  session.ID,
		AreaId:  form.AreaId,
		SkillId: form.SkillId,
		RoomId:  form.RoomId,
	}
	bookmark, err := h.bookmarksService.Create(sqlForm)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	_, err = h.photosService.Create(c, session.ID, bookmark.Id, file, h.Config.PhotosPath, format)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, createResponse{bookmark})
}

// @Summary delete
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /bookmarks/{id} [delete]
func (h *router) delete(c *gin.Context) {
	session := baseApi.GetSession(c)
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	bookmark, err := h.bookmarksService.Delete(id, session.ID)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{bookmark})
}

// @Summary edit
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /bookmarks/{id} [put]
func (h *router) edit(c *gin.Context) {
	session := baseApi.GetSession(c)
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
	}

	var form editForm
	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	bookmark, err := h.bookmarksService.Edit(id, session.ID, form.EditBookmark)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, editResponse{bookmark})
}

// @Summary getAll
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param q query getAllForm true "getAll"
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @router /bookmarks/get_all [get]
func (h *router) getAll(c *gin.Context) {
	session := baseApi.GetSession(c)
	var form getAllForm
	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	bookmarks, total, err := h.bookmarksService.GetAll(form.Limit, form.Page, session.ID, form.Completed, form.OrderById)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, getAllResponse{Data: bookmarks, Total: total})
}
