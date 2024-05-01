package photos

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/internal/handler/api/base_api"
	"os"
	"path/filepath"
	"strconv"
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
	format, err := baseApi.ValidatePhoto(file)
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
	photo, err := h.photosService.Create(c, session.ID, form.BookmarkId, file, h.AppConf.PhotosPath, format)
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
	session := baseApi.GetSession(c)
	id, err := baseApi.GetPathID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	photo, err := h.photosService.Delete(id, session.ID)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	baseApi.Response200(c, deleteResponse{PhotoPreview: photo})
}

// @Summary download
// @Tags photos
// @Param user_id path int true "user_id"
// @Param bookmark_id path int true "bookmark_id"
// @Param name path string true "name"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} baseApi.ErrorResponse
// @Router /photos/download/{user_id}/{bookmark_id}/{name} [get]
func (h *router) download(c *gin.Context) {
	session := baseApi.GetSession(c)
	userId, err := baseApi.GetPathUserID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	if session.ID != userId {
		baseApi.AccessDenied(c)
		return
	}
	bookmarkID, err := baseApi.GetPathBookmarkID(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	name, err := baseApi.GetPathName(c)
	if err != nil {
		baseApi.Response404(c, err)
		return
	}
	path := filepath.Join(h.AppConf.PhotosPath, strconv.Itoa(userId), strconv.Itoa(bookmarkID), name)
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New("файл не существует")
			baseApi.Response404(c, err)
			return
		} else {
			errMessage := fmt.Sprintf("Ошибка при проверке файла: %s", err.Error())
			err = errors.New(errMessage)
			baseApi.Response404(c, err)
			return
		}
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}
