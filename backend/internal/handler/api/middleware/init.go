package middleware

import (
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc/log"
)

var logger = log.GetLogger()

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
