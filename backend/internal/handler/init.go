package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "metroid_bookmarks/docs"
	"metroid_bookmarks/internal/handler/api"
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/service"
)

func InitRoutes(service *service.Service, appConf *models.AppConfig, production bool) *gin.Engine {
	if production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	if !production {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	mwRouter := middleware.NewRouter(service, appConf)

	apiGroup := router.Group("/api/")
	apiRouter := api.NewRouter(mwRouter, service)
	apiRouter.RegisterHandlers(apiGroup)

	return router
}
