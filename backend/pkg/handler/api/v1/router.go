package v1

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/handler/api/v1/areas"
	"metroid_bookmarks/pkg/handler/api/v1/auth"
	"metroid_bookmarks/pkg/handler/api/v1/users"
	"metroid_bookmarks/pkg/service"
)

type V1Router struct {
	*baseApi.BaseAPIRouter
	service *service.Service
}

func NewV1Router(
	baseAPIHandler *baseApi.BaseAPIRouter,
	service *service.Service,
) *V1Router {
	return &V1Router{
		BaseAPIRouter: baseAPIHandler,
		service:       service,
	}
}

func (h *V1Router) RegisterRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authRouter := auth.NewRouter(h.BaseAPIRouter, h.service.Authorization)
	authRouter.RegisterRoutes(authGroup)

	usersGroup := router.Group("/users", h.Middleware.AdminRequired)
	usersRouter := users.NewRouter(h.BaseAPIRouter, h.service.Users)
	usersRouter.RegisterRoutes(usersGroup)

	areasGroup := router.Group("/areas", h.Middleware.AuthRequired)
	areasRouter := areas.NewRouter(h.BaseAPIRouter, h.service.Areas)
	areasRouter.RegisterRoutes(areasGroup)
}
