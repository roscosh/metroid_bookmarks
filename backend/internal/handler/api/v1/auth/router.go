package auth

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"
)

type Router struct {
	*middleware.Router
	service *service.AuthService
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.AuthService,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: service,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/me", h.me)
	router.POST("/sign_up", h.Middleware.LogoutRequired, h.signUp)
	router.POST("/login", h.Middleware.LogoutRequired, h.login)
	router.POST("/logout", h.Middleware.AuthRequired, h.logout)
}
