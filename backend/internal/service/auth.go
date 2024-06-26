package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"metroid_bookmarks/internal/repository/sql/users"

	. "metroid_bookmarks/pkg/session"
)

const (
	salt = "i3490tg4gj94jg0934jg"
)

var ErrUserDoesNotExist = errors.New("нет пользователя с таки логином/паролем")

type AuthService struct {
	sql *users.SQL
}

func newAuthService(sql *users.SQL) *AuthService {
	return &AuthService{sql: sql}
}

func (s *AuthService) Login(login, password string, session *Session) (*Session, error) {
	token := generatePasswordHash(password)
	user, err := s.sql.GetByCredentials(login, token)
	if err != nil {
		return nil, ErrUserDoesNotExist
	}
	session.SetUser(user)
	return session, nil
}

func (s *AuthService) Logout(session *Session) *Session {
	session.ResetUser()
	return session
}

func (s *AuthService) SignUp(createForm *users.CreateUser) (*users.User, error) {
	createForm.Password = generatePasswordHash(createForm.Password)
	user, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return user, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
