package service

import (
	"errors"
	"fmt"
	"metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/repository/redis"
	"metroid_bookmarks/pkg/repository/sql"
)

type MiddlewareService struct {
	sql   *sql.UsersSQL
	redis *redis.SessionRedis
}

func newMiddlewareService(sql *sql.UsersSQL, redis *redis.SessionRedis) *MiddlewareService {
	return &MiddlewareService{sql: sql, redis: redis}
}

func (m *MiddlewareService) CreateSession() (*session.Session, error) {
	var token string
	var result bool
	var err error
	for !result {
		token = session.CreateToken()
		result, err = m.redis.Create(token, session.AnonymousExpires)
		if err != nil {
			errorMessage := fmt.Sprintf("Ошибка Redis: %s", err.Error())
			logger.Error(errorMessage)
			return nil, errors.New(errorMessage)
		}
	}
	return &session.Session{Token: token, Expires: session.AnonymousExpires}, nil
}

func (m *MiddlewareService) GetExistSession(token string) (*session.Session, error) {
	if token == "" {
		return nil, errors.New("нету токена")
	}
	id, err := m.redis.Get(token)
	if err != nil {
		return nil, err
	}
	var expires int
	var user = &sql.User{}
	if id == 0 {
		expires = session.AnonymousExpires
	} else {
		expires = session.AuthenticatedExpires
		user, err = m.sql.GetByID(id)
		if err != nil {
			return nil, err
		}
	}
	return &session.Session{User: *user, Token: token, Expires: expires}, nil
}

func (m *MiddlewareService) UpdateSession(sessionObj *session.Session) {
	m.redis.Update(sessionObj.Token, sessionObj.ID, sessionObj.Expires)
}
