package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "metroid_bookmarks/docs"
	"metroid_bookmarks/internal/handler/api"
	"metroid_bookmarks/internal/handler/api/base_api"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc"
	"os"
)

var logger = misc.GetLogger()

func InitRoutes(service *service.Service, config *misc.Config) *gin.Engine {
	router := gin.New()

	PRODUCTION := os.Getenv("PRODUCTION")

	if PRODUCTION != "true" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	baseAPIRouter := baseApi.NewRouter(service, config)

	apiGroup := router.Group("/api/")
	apiRouter := api.NewRouter(baseAPIRouter, service)
	apiRouter.RegisterHandlers(apiGroup)

	return router
}
