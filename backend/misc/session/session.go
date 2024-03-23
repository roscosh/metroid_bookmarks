package session

import (
	"crypto/rand"
	"encoding/hex"
	"metroid_bookmarks/pkg/repository/sql"
)

const (
	AnonymousExpires     = 3600
	AuthenticatedExpires = 2592000
	CookieSessionName    = "X-Session"
)

type Session struct {
	sql.User
	Token   string `json:"token" db:"token"`
	Expires int    `json:"expires" db:"expires"`
}

func (s *Session) IsAuthenticated() bool {
	if s.ID != 0 {
		return true
	} else {
		return false
	}
}

func (s *Session) IsAdmin() bool {
	return s.User.IsAdmin
}

func (s *Session) SetUser(user *sql.User) {
	s.User = *user
	s.Expires = AuthenticatedExpires
}

func (s *Session) ResetUser() {
	s.User = sql.User{}
	s.Expires = AnonymousExpires
}

func CreateToken() string {
	byteArray := make([]byte, 24)
	if _, err := rand.Read(byteArray); err != nil {
		panic(err)
	}
	return hex.EncodeToString(byteArray)
}
