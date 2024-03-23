package users

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.BaseAPIRouter
	service *service.UsersService
}

func NewRouter(
	baseAPIHandler *baseApi.BaseAPIRouter,
	usersService *service.UsersService,
) *router {
	return &router{
		BaseAPIRouter: baseAPIHandler,
		service:       usersService,
	}
}

func (h *router) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/get_all", h.getAll)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.PUT("/change_password/:id", h.changePassword)
}
