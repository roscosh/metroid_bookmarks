package areas

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*middleware.Router
	service *service.AreasService
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.AreasService,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: service,
	}
}

func (h *Router) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", h.Middleware.AdminRequired, h.create)
	routerGroup.DELETE("/:id", h.Middleware.AdminRequired, h.delete)
	routerGroup.PUT("/:id", h.Middleware.AdminRequired, h.edit)
	routerGroup.GET("/get_all", h.getAll)
}
