package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	. "metroid_bookmarks/misc/session"
	"metroid_bookmarks/pkg/repository/sql"
)

const (
	salt = "cups_managment_relabs"
)

type AuthService struct {
	sql *sql.UsersSQL
}

func newAuthService(sql *sql.UsersSQL) *AuthService {
	return &AuthService{sql: sql}
}

func (s *AuthService) Login(login string, password string, session *Session) (*Session, error) {
	token := generatePasswordHash(password)
	user, err := s.sql.GetUserByCredentials(login, token)
	if err != nil {
		return nil, errors.New("Нет пользователя с таки логином/паролем!")
	}
	session.SetUser(user)
	return session, nil
}

func (s *AuthService) Logout(session *Session) *Session {
	session.ResetUser()
	return session
}

func (s *AuthService) SignUp(createForm *sql.CreateUser) (*sql.User, error) {
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
