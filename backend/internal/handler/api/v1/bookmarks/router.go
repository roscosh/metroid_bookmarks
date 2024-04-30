package bookmarks

import (
	"github.com/gin-gonic/gin"
	baseApi "metroid_bookmarks/internal/handler/api/base_api"
	"metroid_bookmarks/internal/service"
)

type router struct {
	*baseApi.Router
	bookmarksService *service.BookmarksService
	photosService    *service.PhotosService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	bookmarksService *service.BookmarksService,
	photosService *service.PhotosService,
) baseApi.ApiRouter {
	return &router{
		Router:           baseAPIRouter,
		bookmarksService: bookmarksService,
		photosService:    photosService,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.GET("/get_all", h.getAll)
}
