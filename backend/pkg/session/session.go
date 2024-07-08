package session

import (
	"crypto/rand"
	"encoding/hex"
	"metroid_bookmarks/internal/repository/sql/users"
)

const (
	AnonymousExpires     = 3600
	AuthenticatedExpires = 2592000
	TokenLength          = 24
)

type Session struct {
	*users.User
	Token   string `db:"token"   json:"token"`
	Expires int    `db:"expires" json:"expires"`
}

func NewSession(user *users.User, token string, expires int) *Session {
	if user == nil {
		user = new(users.User)
	}
	return &Session{User: user, Token: token, Expires: expires}
}

func (s *Session) IsAuthenticated() bool {
	return s.ID != 0
}

func (s *Session) IsAdmin() bool {
	return s.User.IsAdmin
}

func (s *Session) SetUser(user *users.User) {
	s.User = user
	s.Expires = AuthenticatedExpires
}

func (s *Session) ResetUser() {
	s.User = new(users.User)
	s.Expires = AnonymousExpires
}

func CreateToken() string {
	byteArray := make([]byte, TokenLength)
	if _, err := rand.Read(byteArray); err != nil {
		panic(err)
	}

	return hex.EncodeToString(byteArray)
}
