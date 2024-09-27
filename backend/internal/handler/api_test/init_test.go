package api_test

import (
	"metroid_bookmarks/internal/handler"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/repository/redis"
	"metroid_bookmarks/internal/repository/sql"
	"metroid_bookmarks/internal/service"
	_ "metroid_bookmarks/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitRoutes(t *testing.T) {
	t.Parallel()

	envConf, err := models.NewEnvConfig()
	if err != nil {
		require.NoError(t, err)
	}

	appConf, err := models.NewAppConfig(envConf.AppConfigPath)
	if err != nil {
		panic(err.Error())
	}

	sqlObj := sql.NewSQL(nil)
	redisObj := &redis.Redis{}
	httpService := service.NewService(sqlObj, redisObj)

	engine := handler.InitRoutes(httpService, appConf, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	require.Equal(t, "pong", w.Body.String())
}
