package api

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/handler/api/v1"
	"metroid_bookmarks/pkg/service"
)

type ApiRouter struct {
	*baseApi.BaseAPIRouter
	service *service.Service
}

func NewApiRouter(
	baseAPIHandler *baseApi.BaseAPIRouter,
	service *service.Service,
) *ApiRouter {
	return &ApiRouter{
		BaseAPIRouter: baseAPIHandler,
		service:       service,
	}
}

func (h *ApiRouter) RegisterRoutes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1", h.Middleware.GetSession)
	v1Router := v1.NewV1Router(h.BaseAPIRouter, h.service)
	v1Router.RegisterRoutes(v1Group)
}
