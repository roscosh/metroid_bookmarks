package auth

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
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

func (r *Router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/me", r.me)
	router.POST("/sign_up", r.Middleware.LogoutRequired, r.signUp)
	router.POST("/login", r.Middleware.LogoutRequired, r.login)
	router.POST("/logout", r.Middleware.AuthRequired, r.logout)
}
