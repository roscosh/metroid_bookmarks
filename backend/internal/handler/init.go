package handler

import (
	_ "metroid_bookmarks/docs" //revive:disable:blank-imports
	"metroid_bookmarks/internal/handler/api"
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRoutes(service *service.Service, appConf *models.AppConfig, production bool) *gin.Engine {
	if production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	if !production {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	mwRouter := middleware.NewRouter(service, appConf)

	apiGroup := router.Group("/api/")
	apiRouter := api.NewRouter(mwRouter, service)
	apiRouter.RegisterHandlers(apiGroup)

	return router
}
