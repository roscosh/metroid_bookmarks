package photos

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc/log"

	"github.com/gin-gonic/gin"
)

var logger = log.GetLogger()

type Router struct {
	*middleware.Router
	photosService    *service.PhotosService
	bookmarksService *service.BookmarksService
}

func NewRouter(
	mwRouter *middleware.Router,
	photosService *service.PhotosService,
	bookmarksService *service.BookmarksService,
) *Router {
	return &Router{
		Router:           mwRouter,
		photosService:    photosService,
		bookmarksService: bookmarksService,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", h.create)
	router.DELETE("/:id", h.delete)
	router.GET("/download/:user_id/:bookmark_id/:name", h.download)
}
