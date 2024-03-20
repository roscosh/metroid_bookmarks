package baseApi

import (
	"metroid_bookmarks/misc"
	"metroid_bookmarks/pkg/service"
)

var logger = misc.GetLogger()

type BaseAPIRouter struct {
	Middleware *Handler
	Config     *misc.Config
}

func NewBaseAPIRouter(services *service.Service, config *misc.Config) *BaseAPIRouter {
	return &BaseAPIRouter{
		Middleware: NewHandler(services.Middleware, config),
		Config:     config,
	}
}
