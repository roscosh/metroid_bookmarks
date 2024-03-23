package auth

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.BaseAPIRouter
	service *service.AuthService
}

func NewRouter(
	baseAPIHandler *baseApi.BaseAPIRouter,
	service *service.AuthService,
) *router {
	return &router{
		BaseAPIRouter: baseAPIHandler,
		service:       service,
	}
}

func (h *router) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/me", h.me)
	router.POST("/sign_up", h.signUp)
	router.POST("/login", h.login)
	router.POST("/logout", h.Middleware.AuthRequired, h.logout)
}
