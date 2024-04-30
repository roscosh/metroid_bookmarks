package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/internal/handler"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/repository/redis"
	"metroid_bookmarks/internal/repository/sql"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc"
	"os"
	"os/signal"
	"syscall"
)

var logger = misc.GetLogger()

func Init() {
	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
		return
	}

	logger.SetParams(envConf.LogLevel)

	config, err := models.NewConfig(envConf.DbConfig)
	if err != nil {
		logger.Error(err.Error())
		return
	}
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
	err = misc.DbMigrate(config.Db.Dsn, envConf.DbmateMigrationsDir)
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

	if envConf.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	srv := new(misc.Server)
	go func() {
		if err = srv.Run(handler.InitRoutes(newService, config, envConf)); err != nil {
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
