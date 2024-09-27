package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"metroid_bookmarks/internal/handler"
	"metroid_bookmarks/internal/handler/api/middleware"
	handlerAreas "metroid_bookmarks/internal/handler/api/v1/areas"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/internal/repository/redis"
	mock_session "metroid_bookmarks/internal/repository/redis/session/mocks"
	"metroid_bookmarks/internal/repository/sql"
	"metroid_bookmarks/internal/repository/sql/areas"
	mock_areas "metroid_bookmarks/internal/repository/sql/areas/mocks"
	"metroid_bookmarks/internal/repository/sql/users"
	mock_users "metroid_bookmarks/internal/repository/sql/users/mocks"
	"metroid_bookmarks/internal/service"
	"metroid_bookmarks/pkg/session"
	_ "metroid_bookmarks/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEditArea(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	envConf, err := models.NewEnvConfig()
	require.NoError(t, err)

	a, err := models.NewAppConfig(envConf.AppConfigPath)
	require.NoError(t, err)

	usersSQL := mock_users.NewMockSQL(ctl)
	areasSQL := mock_areas.NewMockSQL(ctl)

	sqlObj := &sql.SQL{
		Users:     usersSQL,
		Areas:     areasSQL,
		Rooms:     nil,
		Skills:    nil,
		Bookmarks: nil,
		Photos:    nil,
	}
	sessionRedis := mock_session.NewMockRedis(ctl)
	redisObj := &redis.Redis{
		Session: sessionRedis,
	}
	httpService := service.NewService(sqlObj, redisObj)

	token := "23423432423sfgdfgsdfg"
	userId := 1
	exceptUser := &users.User{
		Name:    "roascosh",
		Login:   "roascosh@mail.ru",
		ID:      userId,
		IsAdmin: true,
	}
	areaID := 1
	nameEn := "rome"
	nameRu := "рим"
	editArea := &areas.EditArea{
		NameEn: &nameEn,
		NameRu: &nameRu,
	}
	form := handlerAreas.EditForm{
		EditArea: editArea,
	}

	exceptArea := &areas.Area{
		NameEn: nameEn,
		NameRu: nameRu,
		ID:     areaID,
	}
	exceptBuff, err := json.Marshal(exceptArea)
	require.NoError(t, err)

	sessionRedis.EXPECT().Get(token).Return(userId, nil).Times(1)
	usersSQL.EXPECT().GetByID(userId).Return(exceptUser, nil).Times(1)
	sessionRedis.EXPECT().Update(token, userId, session.AuthenticatedExpires).Times(1)
	areasSQL.EXPECT().Edit(areaID, editArea).Return(exceptArea, nil).Times(1)

	engine := handler.InitRoutes(httpService, a, true)

	w := httptest.NewRecorder()

	buff, err := json.Marshal(form)
	require.NoError(t, err)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/areas/%d", areaID), bytes.NewBuffer(buff))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	req.AddCookie(&http.Cookie{
		Name:  middleware.SessionCookieKey,
		Value: token,
	})

	engine.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	require.Equal(t, string(exceptBuff), string(data))
}
