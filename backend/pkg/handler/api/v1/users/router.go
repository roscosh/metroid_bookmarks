package users

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.UsersService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	usersService *service.UsersService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: usersService,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/get_all", h.getAll)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.PUT("/change_password/:id", h.changePassword)
}
