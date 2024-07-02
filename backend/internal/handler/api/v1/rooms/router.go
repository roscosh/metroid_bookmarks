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

func (r *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/", r.Middleware.AdminRequired, r.create)
	router.DELETE("/:id", r.Middleware.AdminRequired, r.delete)
	router.PUT("/:id", r.Middleware.AdminRequired, r.edit)
	router.GET("/get_all", r.getAll)
}
