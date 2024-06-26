package skills

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"metroid_bookmarks/internal/handler/api/middleware"
)

// @Summary create
// @Tags skills
// @Accept json
// @Produce json
// @Param input body createForm true "create"
// @Success 200 {object}  createResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /skills/ [post]
func (h *Router) create(c *gin.Context) {
	var form createForm
	err := c.ShouldBindWith(&form, binding.JSON)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	skill, err := h.service.Create(form.CreateSkill)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, createResponse{Skill: skill})
}

// @Summary edit
// @Tags skills
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param input body editForm true "edit"
// @Success 200 {object}  editResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /skills/{id} [put]
func (h *Router) edit(c *gin.Context) {
	id, err := middleware.GetPathID(c)
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
	skill, err := h.service.Edit(id, form.EditSkill)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, editResponse{Skill: skill})
}

// @Summary delete
// @Tags skills
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} deleteResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /skills/{id} [delete]
func (h *Router) delete(c *gin.Context) {
	id, err := middleware.GetPathID(c)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	skill, err := h.service.Delete(id)
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, deleteResponse{Skill: skill})
}

// @Summary getAll
// @Tags skills
// @Accept json
// @Produce json
// @Success 200 {object}  getAllResponse
// @Failure 404 {object} middleware.ErrorResponse
// @router /skills/get_all [get]
func (h *Router) getAll(c *gin.Context) {
	skill, total, err := h.service.GetAll()
	if err != nil {
		middleware.Response404(c, err)
		return
	}
	middleware.Response200(c, getAllResponse{Data: skill, Total: total})
}
