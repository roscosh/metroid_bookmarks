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

func (r *Router) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", r.Middleware.AdminRequired, r.create)
	routerGroup.DELETE("/:id", r.Middleware.AdminRequired, r.delete)
	routerGroup.PUT("/:id", r.Middleware.AdminRequired, r.edit)
	routerGroup.GET("/get_all", r.getAll)
}
