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
	EnvConf    *models.EnvConfig
}

func NewRouter(services *service.Service, appConf *models.AppConfig, envConf *models.EnvConfig) *Router {
	return &Router{
		Middleware: NewMiddleware(services.Middleware),
		AppConf:    appConf,
		EnvConf:    envConf,
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
