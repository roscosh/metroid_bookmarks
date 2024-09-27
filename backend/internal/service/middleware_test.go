package service

import (
	"errors"
	"fmt"
	mock_session "metroid_bookmarks/internal/repository/redis/session/mocks"
	"metroid_bookmarks/internal/repository/sql/users"
	mock_users "metroid_bookmarks/internal/repository/sql/users/mocks"
	"metroid_bookmarks/pkg/session"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMiddlewareService_CreateSession(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	var token string

	sessionRedis.EXPECT().Create(gomock.AssignableToTypeOf(token), session.AnonymousExpires).DoAndReturn(
		func(key string, _ int) (bool, error) {
			token = key
			return true, nil
		},
	)

	sessionObj, err := middlewareService.CreateSession()
	require.NoError(t, err)

	expectedSession := &session.Session{
		User:    new(users.User),
		Token:   token,
		Expires: session.AnonymousExpires,
	}

	require.Equal(t, expectedSession.User, sessionObj.User)
	require.Equal(t, expectedSession.Token, sessionObj.Token)
	require.Equal(t, expectedSession.Expires, sessionObj.Expires)
}

func TestMiddlewareService_CreateSession_10FalseSaves(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	var token string

	sessionRedis.EXPECT().Create(gomock.AssignableToTypeOf(token), session.AnonymousExpires).DoAndReturn(
		func(key string, _ int) (bool, error) {
			token = key
			return false, nil
		},
	).Times(10)
	sessionRedis.EXPECT().Create(gomock.AssignableToTypeOf(token), session.AnonymousExpires).DoAndReturn(
		func(key string, _ int) (bool, error) {
			token = key
			return true, nil
		},
	)

	sessionObj, err := middlewareService.CreateSession()
	require.NoError(t, err)

	expectedSession := &session.Session{
		User:    new(users.User),
		Token:   token,
		Expires: session.AnonymousExpires,
	}

	require.Equal(t, expectedSession.User, sessionObj.User)
	require.Equal(t, expectedSession.Token, sessionObj.Token)
	require.Equal(t, expectedSession.Expires, sessionObj.Expires)
}

func TestMiddlewareService_CreateSession_ErrorConnClosed(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	var token string

	errConnClosed := errors.New("redigo: connection closed")
	sessionRedis.EXPECT().Create(gomock.AssignableToTypeOf(token), session.AnonymousExpires).Return(false, errConnClosed)

	sessionObj, err := middlewareService.CreateSession()
	require.EqualError(t, err, errConnClosed.Error())
	require.Empty(t, sessionObj)
}

func TestMiddlewareService_GetExistSession_ErrNoToken(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	var token string

	sessionObj, err := middlewareService.GetExistSession(token)
	require.EqualError(t, err, ErrNoToken.Error())
	require.Empty(t, sessionObj)
}

func TestMiddlewareService_GetExistSession_ErrorConnClosed(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	token := "we234f32fp3f23f32f23f3f3f"
	errConnClosed := errors.New("redigo: connection closed")
	userID := 0

	sessionRedis.EXPECT().Get(token).Return(userID, errConnClosed)

	sessionObj, err := middlewareService.GetExistSession(token)
	require.EqualError(t, err, errConnClosed.Error())
	require.Empty(t, sessionObj)
}

func TestMiddlewareService_GetExistSession_Anonymous(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	token := "we234f32fp3f23f32f23f3f3f"
	userID := 0

	sessionRedis.EXPECT().Get(token).Return(userID, nil)

	expectedSession := &session.Session{
		User:    new(users.User),
		Token:   token,
		Expires: session.AnonymousExpires,
	}

	sessionObj, err := middlewareService.GetExistSession(token)
	require.NoError(t, err)
	require.Equal(t, expectedSession, sessionObj)
}

func TestMiddlewareService_GetExistSession_ErrNoID(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	const token = "we234f32fp3f23f32f23f3f3f"

	const userID = 1
	errNoID := errors.New(fmt.Sprintf(`No row with id="%v"!`, userID))

	sessionRedis.EXPECT().Get(token).Return(userID, nil)
	usersSQL.EXPECT().GetByID(userID).Return(nil, errNoID)

	sessionObj, err := middlewareService.GetExistSession(token)
	require.EqualError(t, err, errNoID.Error())
	require.Empty(t, sessionObj)
}

func TestMiddlewareService_GetExistSession_Authenticated(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	const token = "we234f32fp3f23f32f23f3f3f"

	const userID = 1
	user := &users.User{
		Name:    "roascosh",
		Login:   "roascosh@mail.ru",
		ID:      userID,
		IsAdmin: false,
	}

	sessionRedis.EXPECT().Get(token).Return(userID, nil)
	usersSQL.EXPECT().GetByID(userID).Return(user, nil)

	expectedSession := &session.Session{
		User:    user,
		Token:   token,
		Expires: session.AuthenticatedExpires,
	}

	sessionObj, err := middlewareService.GetExistSession(token)
	require.NoError(t, err)
	require.Equal(t, expectedSession, sessionObj)
}

func TestMiddlewareService_UpdateSession(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sessionRedis := mock_session.NewMockRedis(ctl)
	usersSQL := mock_users.NewMockSQL(ctl)
	middlewareService := newMiddlewareService(usersSQL, sessionRedis)

	const token = "we234f32fp3f23f32f23f3f3f"
	expectedSession := &session.Session{
		User:    new(users.User),
		Token:   token,
		Expires: session.AnonymousExpires,
	}

	sessionRedis.EXPECT().Update(expectedSession.Token, expectedSession.ID, expectedSession.Expires)

	middlewareService.UpdateSession(expectedSession)
}
