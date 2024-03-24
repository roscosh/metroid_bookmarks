package bookmarks

import (
	"github.com/gin-gonic/gin"
	baseApi "metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.BookmarksService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.BookmarksService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.GET("/get_all", h.getAll)
}
