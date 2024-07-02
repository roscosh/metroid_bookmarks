package api

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	v1 "metroid_bookmarks/internal/handler/api/v1"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
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

func (r *Router) RegisterHandlers(router *gin.RouterGroup) {
	v1Group := router.Group("/v1", r.Middleware.SessionRequired)
	v1Router := v1.NewRouter(r.Router, r.service)
	v1Router.RegisterHandlers(v1Group)
}
