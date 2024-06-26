package rooms

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*middleware.Router
	service *service.RoomsService
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.RoomsService,
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
