package areas

import (
	"github.com/gin-gonic/gin"
	baseApi "metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.BaseAPIRouter
	service *service.AreasService
}

func NewRouter(baseAPIRouter *baseApi.BaseAPIRouter, service *service.AreasService) *router {
	return &router{BaseAPIRouter: baseAPIRouter, service: service}
}

func (h *router) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/", h.Middleware.AdminRequired, h.create)
	router.DELETE("/:id", h.Middleware.AdminRequired, h.delete)
	router.PUT("/:id", h.Middleware.AdminRequired, h.edit)
	router.GET("/get_all", h.getAll)
}
