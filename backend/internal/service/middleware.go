package service

import (
	redisSession "metroid_bookmarks/internal/repository/redis/session"
	"metroid_bookmarks/internal/repository/sql/users"
	"metroid_bookmarks/pkg/session"
)

type MiddlewareService struct {
	sql   users.SQL
	redis redisSession.Redis
}

func newMiddlewareService(sql users.SQL, redis redisSession.Redis) *MiddlewareService {
	return &MiddlewareService{sql: sql, redis: redis}
}

func (m *MiddlewareService) CreateSession() (*session.Session, error) {
	var (
		token  string
		result bool
		err    error
	)

	for !result {
		token = session.CreateToken()

		result, err = m.redis.Create(token, session.AnonymousExpires)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return session.NewSession(nil, token, session.AnonymousExpires), nil
}

func (m *MiddlewareService) GetExistSession(token string) (*session.Session, error) {
	if token == "" {
		return nil, ErrNoToken
	}

	userID, err := m.redis.Get(token)
	if err != nil {
		return nil, err
	}

	var expires int

	user := new(users.User)

	if userID == 0 {
		expires = session.AnonymousExpires
	} else {
		expires = session.AuthenticatedExpires

		user, err = m.sql.GetByID(userID)
		if err != nil {
			return nil, err
		}
	}

	return session.NewSession(user, token, expires), nil
}

func (m *MiddlewareService) UpdateSession(sessionObj *session.Session) {
	m.redis.Update(sessionObj.Token, sessionObj.ID, sessionObj.Expires)
}
