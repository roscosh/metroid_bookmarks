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
	Config     *models.Config
	EnvConfig  *models.EnvConfig
}

func NewRouter(services *service.Service, config *models.Config, envConf *models.EnvConfig) *Router {
	return &Router{
		Middleware: NewMiddleware(services.Middleware),
		Config:     config,
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
