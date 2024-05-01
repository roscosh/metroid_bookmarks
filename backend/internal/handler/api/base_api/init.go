package baseApi

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc"
)

var logger = misc.GetLogger()

type Router struct {
	Middleware *Middleware
	AppConf    *models.AppConfig
}

func NewRouter(services *service.Service, appConf *models.AppConfig) *Router {
	return &Router{
		Middleware: NewMiddleware(services.Middleware),
		AppConf:    appConf,
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
