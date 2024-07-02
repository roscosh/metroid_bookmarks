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

func (r *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/get_all", r.getAll)
	router.DELETE("/:id", r.delete)
	router.PUT("/:id", r.edit)
	router.PUT("/change_password/:id", r.changePassword)
}
