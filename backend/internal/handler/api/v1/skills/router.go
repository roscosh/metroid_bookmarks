package skills

import (
	"github.com/gin-gonic/gin"
	baseApi "metroid_bookmarks/internal/handler/api/base_api"
	"metroid_bookmarks/internal/service"
)

type router struct {
	*baseApi.Router
	service *service.SkillsService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.SkillsService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.Middleware.AdminRequired, h.create)
	router.DELETE("/:id", h.Middleware.AdminRequired, h.delete)
	router.PUT("/:id", h.Middleware.AdminRequired, h.edit)
	router.GET("/get_all", h.getAll)
}
