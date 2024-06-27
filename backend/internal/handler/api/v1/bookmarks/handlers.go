package bookmarks

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/repository/sql/bookmarks"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary create
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param input formData createForm true "create"
// @Param photo formData file true "photo"
// @Success 200 {object} createResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /bookmarks/ [post]
func (h *Router) create(c *gin.Context) {
	session := middleware.GetSession(c)

	var form createForm

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	file, err := middleware.GetPhoto(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	format, err := middleware.ValidatePhoto(file)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	sqlForm := &bookmarks.CreateBookmark{
		UserID:  session.ID,
		AreaID:  form.AreaID,
		SkillID: form.SkillID,
		RoomID:  form.RoomID,
	}

	bookmark, err := h.bookmarksService.Create(sqlForm)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	_, err = h.photosService.Create(c, session.ID, bookmark.ID, file, h.AppConf.PhotosPath, format)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, createResponse{bookmark})
}

// @Summary delete
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /bookmarks/{id} [delete]
func (h *Router) delete(c *gin.Context) {
	session := middleware.GetSession(c)

	bookmarkID, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	bookmark, err := h.bookmarksService.Delete(bookmarkID, session.ID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, deleteResponse{bookmark})
}

// @Summary edit
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /bookmarks/{id} [put]
func (h *Router) edit(c *gin.Context) {
	session := middleware.GetSession(c)

	bookmarkID, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
	}

	var form editForm

	err = c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	bookmark, err := h.bookmarksService.Edit(bookmarkID, session.ID, form.EditBookmark)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, editResponse{bookmark})
}

// @Summary getAll
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param q query getAllForm true "getAll"
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /bookmarks/get_all [get]
func (h *Router) getAll(c *gin.Context) {
	session := middleware.GetSession(c)

	var form getAllForm

	err := c.ShouldBindWith(&form, binding.Query)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	bookmarkList, total, err := h.bookmarksService.GetAll(form.Limit, form.Page, session.ID, form.Completed, form.OrderByID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, getAllResponse{Data: bookmarkList, Total: total})
}
