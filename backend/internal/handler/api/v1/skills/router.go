package skills

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"
)

type Router struct {
	*middleware.Router
	service *service.SkillsService
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.SkillsService,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: service,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.Middleware.AdminRequired, h.create)
	router.DELETE("/:id", h.Middleware.AdminRequired, h.delete)
	router.PUT("/:id", h.Middleware.AdminRequired, h.edit)
	router.GET("/get_all", h.getAll)
}
