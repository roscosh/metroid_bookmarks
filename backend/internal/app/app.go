package app

import (
	"context"
	"errors"
	"metroid_bookmarks/internal/handler"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/repository/redis"
	"metroid_bookmarks/internal/repository/sql"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/misc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = misc.GetLogger()

type App struct {
	envConf   *models.EnvConfig
	dbPool    *sql.DbPool
	redisPool *redis.RedisPool
	srv       *misc.Server
}

func NewApp(envConf *models.EnvConfig) *App {
	return &App{envConf: envConf}
}

func (a *App) Init() {
	appConf, err := models.NewAppConfig(a.envConf.AppConfigPath)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	a.start(appConf)

	logger.Info("METROID BOOKMARKS API started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.shutdown()

	logger.Info("METROID BOOKMARKS API shutting down successfully")
}

func (a *App) start(appConf *models.AppConfig) {
	var err error

	a.dbPool, err = sql.NewDbPool(
		appConf.PostgreSQL.Dsn,
		a.envConf.MinConns,
		a.envConf.MaxConns,
		a.envConf.MaxConnLifetime,
		a.envConf.MaxConnIdleTime,
		a.envConf.HealthCheckPeriod,
	)
	if err != nil {
		logger.Errorf("failed to initialize postgreSQL: %s\n", err.Error())
		return
	}

	a.redisPool, err = redis.NewRedisPool(appConf.Redis.Dsn)
	if err != nil {
		logger.Errorf("failed to initialize redis: %s\n", err.Error())
		return
	}

	sqlObj := sql.NewSQL(a.dbPool)
	redisObj := redis.NewRedis(a.redisPool)
	newService := service.NewService(sqlObj, redisObj)

	err = misc.DbMigrate(appConf.PostgreSQL.Dsn, appConf.DbmateMigrationsDir)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	a.srv = new(misc.Server)

	go func() {
		if err = a.srv.Run(handler.InitRoutes(newService, appConf, a.envConf.Production)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Errorf("error occured while running http server: %s", err.Error())
			}
		}
	}()
}

// shutdown function for graceful close. The function will not run on develop mode
func (a *App) shutdown() {
	if err := a.srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	} else {
		logger.Info("HTTP server closed")
	}

	if err := a.dbPool.Close(); err != nil {
		logger.Errorf("error occured on redis connection close: %s", err.Error())
	} else {
		logger.Info("postgreSQL connection pool closed successfully")
	}

	if err := a.redisPool.Close(); err != nil {
		logger.Errorf("error occured on redis connection close: %s", err.Error())
	} else {
		logger.Info("redis connection pool closed successfully")
	}
}
