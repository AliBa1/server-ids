package sessions

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type SessionsDB struct {
	Sessions map[uuid.UUID]string
}

func NewSessionsDB() *SessionsDB {
	return &SessionsDB{
		Sessions: make(map[uuid.UUID]string),
	}
}

func (s *SessionsDB) AddSession(token uuid.UUID, username string) {
	s.Sessions[token] = username
}

func (s *SessionsDB) GetUsername(token uuid.UUID) (string, error) {
	username, ok := s.Sessions[token]
	if username == "" || !ok {
		return "", fmt.Errorf("user not found")
	}
	return username, nil
}

func (s *SessionsDB) IsUserLoggedIn(r *http.Request) bool {
	keyCookie, err := r.Cookie("session_key")
	if err != nil {
		return false
	}

	key, err := uuid.Parse(keyCookie.Value)
	if err != nil {
		return false
	}

	_, err = s.GetUsername(key)
	return err == nil
}
