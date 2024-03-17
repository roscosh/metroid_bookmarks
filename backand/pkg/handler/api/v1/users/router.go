package users

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type UsersRouter struct {
	*baseApi.BaseAPIRouter
	usersService *service.UsersService
}

func NewUsersRouter(
	baseAPIHandler *baseApi.BaseAPIRouter,
	usersService *service.UsersService,
) *UsersRouter {
	return &UsersRouter{
		BaseAPIRouter: baseAPIHandler,
		usersService:  usersService,
	}
}

func (h *UsersRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/get_all", h.getAllUsers)
	router.POST("/", h.Middleware.AdminRequired, h.createUser)
	router.DELETE("/:id", h.Middleware.AdminRequired, h.deleteUser)
	router.PUT("/:id", h.Middleware.AdminRequired, h.editUser)
	router.PUT("/change_password/:id", h.Middleware.AdminRequired, h.changePassword)

}
