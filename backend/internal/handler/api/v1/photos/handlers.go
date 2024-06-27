package photos

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary create
// @Tags photos
// @Accept json
// @Produce json
// @Param form formData createForm true "create"
// @Param photo formData file true "photo"
// @Success 200 {object}  createResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /photos/ [post]
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
	bookmark, err := h.bookmarksService.GetByID(form.BookmarkID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	if bookmark.UserID != session.ID {
		middleware.AccessDenied(c)
		return
	}
	photo, err := h.photosService.Create(c, session.ID, form.BookmarkID, file, h.AppConf.PhotosPath, format)
	if err != nil {
		middleware.Response404(c, err)
		return
	}

	middleware.Response200(c, createResponse{PhotoPreview: photo})
}

// @Summary delete
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /photos/{id} [delete]
func (h *Router) delete(c *gin.Context) {
	session := middleware.GetSession(c)
	id, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	photo, err := h.photosService.Delete(id, session.ID)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, deleteResponse{PhotoPreview: photo})
}

// @Summary download
// @Tags photos
// @Param user_id path int true "user_id"
// @Param bookmark_id path int true "bookmark_id"
// @Param name path string true "name"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /photos/download/{user_id}/{bookmark_id}/{name} [get]
func (h *Router) download(c *gin.Context) {
	session := middleware.GetSession(c)
	userID, err := middleware.GetPathUserID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	if session.ID != userID {
		middleware.AccessDenied(c)
		return
	}
	bookmarkID, err := middleware.GetPathBookmarkID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	name, err := middleware.GetPathName(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	path := filepath.Join(h.AppConf.PhotosPath, strconv.Itoa(userID), strconv.Itoa(bookmarkID), name)
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			middleware.Response404(c, ErrFileDoesNotExist)
		} else {
			errMessage := "ошибка при проверке файла: " + err.Error()
			logger.Error(errMessage)
			middleware.Response404(c, &Error{message: errMessage})
		}

		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}
