package users

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*middleware.Router
	service *service.UsersService
}

func NewRouter(
	mwRouter *middleware.Router,
	usersService *service.UsersService,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: usersService,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/get_all", h.getAll)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.PUT("/change_password/:id", h.changePassword)
}
