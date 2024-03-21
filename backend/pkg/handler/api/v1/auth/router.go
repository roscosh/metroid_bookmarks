package auth

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/service"
)

type AuthRouter struct {
	*baseApi.BaseAPIRouter
	authService  *service.AuthService
	usersService *service.UsersService
}

func NewAuthRouter(
	baseAPIHandler *baseApi.BaseAPIRouter,
	authService *service.AuthService,
	usersService *service.UsersService,
) *AuthRouter {
	return &AuthRouter{
		BaseAPIRouter: baseAPIHandler,
		authService:   authService,
		usersService:  usersService,
	}
}

func (h *AuthRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/me", h.me)
	if h.Config.Debug {
		router.POST("/sign_up", h.signUp) //ТОЛЬКО ДЛЯ ТЕСТОВ, на продакшене использовать /users/create
	}
	router.POST("/login", h.login)
	router.POST("/logout", h.Middleware.AuthRequired, h.logout)
}
