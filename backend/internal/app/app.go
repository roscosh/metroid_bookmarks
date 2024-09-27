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
	"metroid_bookmarks/pkg/misc/log"
	"metroid_bookmarks/pkg/pgpool"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.GetLogger()

type App struct {
	envConf   *models.EnvConfig
	dbPool    *pgpool.PgPool
	redisPool *redis.Pool
	srv       *misc.Server
}

func NewApp(envConf *models.EnvConfig) *App {
	return &App{envConf: envConf}
}

func (a *App) Run() {
	appConf, err := models.NewAppConfig(a.envConf.AppConfigPath)
	if err != nil {
		logger.Fatal(err.Error())
	}

	a.startUp(appConf)

	logger.Info("METROID BOOKMARKS API started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.shutDown()

	logger.Info("METROID BOOKMARKS API shutting down successfully")
}

func (a *App) startUp(appConf *models.AppConfig) {
	var err error

	a.dbPool, err = pgpool.NewPgPool(
		a.envConf.DatabaseURL,
		a.envConf.MinConns,
		a.envConf.MaxConns,
		a.envConf.MaxConnLifetime,
		a.envConf.MaxConnIdleTime,
		a.envConf.HealthCheckPeriod,
	)
	if err != nil {
		logger.Fatalf("failed to initialize postgreSQL: %s\n", err.Error())
		return
	}

	a.redisPool, err = redis.NewRedisPool(appConf.Redis.Dsn)
	if err != nil {
		logger.Fatalf("failed to initialize redis: %s\n", err.Error())
		return
	}

	sqlObj := sql.NewSQL(a.dbPool)
	redisObj := redis.NewRedis(a.redisPool)
	httpService := service.NewService(sqlObj, redisObj)

	a.srv = new(misc.Server)

	go func() {
		err = a.srv.Run(handler.InitRoutes(httpService, appConf, a.envConf.Production))
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Fatalf("error occurred while running http server: %s", err.Error())
			}
		}
	}()
}

// shutDown function for graceful close. The function will not run in `develop` mode
func (a *App) shutDown() {
	if err := a.srv.ShutDown(context.Background()); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	} else {
		logger.Info("HTTP server closed")
	}

	if err := a.dbPool.Close(); err != nil {
		logger.Errorf("error occurred on redis connection close: %s", err.Error())
	} else {
		logger.Info("postgreSQL connection pool closed successfully")
	}

	if err := a.redisPool.Close(); err != nil {
		logger.Errorf("error occurred on redis connection close: %s", err.Error())
	} else {
		logger.Info("redis connection pool closed successfully")
	}
}
