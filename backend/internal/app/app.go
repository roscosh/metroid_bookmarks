package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/handler"
	"metroid_bookmarks/internal/repository/redis"
	"metroid_bookmarks/internal/repository/sql"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc"
	"os"
	"os/signal"
	"syscall"
)

var logger = misc.GetLogger()

// @title METROID BOOKMARKS API
// @version 1.0
// @description API Server for metroid bookmarks
// @host localhost:3000
// @BasePath /api/v1
func Init() {
	config := misc.GetConfig()
	dbPool, err := sql.NewDbPool(config.Db.Dsn)
	if err != nil {
		logger.Errorf("failed to create db dbPool: %s\n", err.Error())
		return
	}
	SQL := sql.NewSQL(dbPool)
	if err != nil {
		logger.Errorf("failed to initialize db: %s\n", err.Error())
		return
	}
	err = misc.DbMigrate()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	redisClient, err := redis.NewRedisPool(config.Redis.Dsn)
	if err != nil {
		logger.Errorf("failed to initialize redis: %s\n", err.Error())
		return
	}
	newRedis := redis.NewRedis(redisClient)
	newService := service.NewService(SQL, newRedis)

	PRODUCTION := os.Getenv("PRODUCTION")
	if PRODUCTION == "true" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	srv := new(misc.Server)
	go func() {
		if err = srv.Run(handler.InitRoutes(newService, config)); err != nil {
			logger.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Info("METROID BOOKMARKS API started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("METROID BOOKMARKS API shutting down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s\n", err.Error())
	}

	dbPool.Close()
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("Recovered from panic: %s", r)
		}
	}()

	if err = redisClient.Close(); err != nil {
		logger.Errorf("error occured on redis connection close: %s\n", err.Error())
	}
}
