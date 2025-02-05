package pgpool

import (
	"context"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/pkg/misc/log"
	"testing"
)

func newTestPgPool() *PgPool {
	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
	}

	logger := log.GetLogger()
	logger.SetParams(envConf.LogLevel)

	pgPool, err := NewPgPool(
		context.Background(),
		envConf.DatabaseURL,
		envConf.MinConns,
		envConf.MaxConns,
		envConf.MaxConnLifetime,
		envConf.MaxConnIdleTime,
		envConf.HealthCheckPeriod,
	)
	if err != nil {
		panic(err)
	}

	return pgPool
}

func TestNewPgPool(t *testing.T) {
	pgPool := newTestPgPool()
	defer pgPool.Close()
}
