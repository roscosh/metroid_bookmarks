package handler

import (
	_ "metroid_bookmarks/test"
	"testing"
)

func TestInitRoutes(t *testing.T) {
	//	envConf, err := models.NewEnvConfig()
	//	if err != nil {
	//		require.NoError(t, err)
	//	}
	//	a, err := models.NewAppConfig(envConf.AppConfigPath)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	sqlObj := sql.NewSQL(a.dbPool)
	//	redisObj := redis.NewRedis(a.redisPool)
	//	httpService := service.NewService(sqlObj, redisObj)
	//
	//	engine := InitRoutes(httpService, a, false)
	//
	//	w := httptest.NewRecorder()
	//	req, _ := http.NewRequest("GET", "/ping", nil)
	//	engine.ServeHTTP(w, req)
	//	require.Equal(t, 200, w.Code)
	//	require.Equal(t, "pong", w.Body.String())
	//}
}
