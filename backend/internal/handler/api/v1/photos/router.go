package photos

import (
	"github.com/gin-gonic/gin"
	baseApi "metroid_bookmarks/internal/handler/api/base_api"
	"metroid_bookmarks/internal/service"
)

type router struct {
	*baseApi.Router
	photosService    *service.PhotosService
	bookmarksService *service.BookmarksService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	photosService *service.PhotosService,
	bookmarksService *service.BookmarksService,
) baseApi.ApiRouter {
	return &router{
		Router:           baseAPIRouter,
		photosService:    photosService,
		bookmarksService: bookmarksService,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.GET("/download/:user_id/:bookmark_id/:name", h.download)
}
