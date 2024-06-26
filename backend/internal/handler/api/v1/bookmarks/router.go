package bookmarks

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*middleware.Router
	bookmarksService *service.BookmarksService
	photosService    *service.PhotosService
}

func NewRouter(
	mwRouter *middleware.Router,
	bookmarksService *service.BookmarksService,
	photosService *service.PhotosService,
) *Router {
	return &Router{
		Router:           mwRouter,
		bookmarksService: bookmarksService,
		photosService:    photosService,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.GET("/get_all", h.getAll)
}
