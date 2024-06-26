package api

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/handler/api/v1"
	"metroid_bookmarks/internal/service"
)

type Router struct {
	*middleware.Router
	service *service.Service
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.Service,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: service,
	}
}

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	v1Group := router.Group("/v1", h.Middleware.SessionRequired)
	v1Router := v1.NewRouter(h.Router, h.service)
	v1Router.RegisterHandlers(v1Group)
}
