package sessions

import (
	"fmt"
	"net/http"
	"server-ids/internal/mock"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

type SessionsDB struct {
	Users    []models.User
	Sessions map[uuid.UUID]string
}

func NewSessionsDB() *SessionsDB {
	return &SessionsDB{
		Users:    mock.GetMockUsers(),
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

func (s *SessionsDB) GetUser(username string) (*models.User, error) {
	for _, u := range s.Users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", username)
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

func (s *SessionsDB) IsUserEmployee(r *http.Request) bool {
	keyCookie, err := r.Cookie("session_key")
	if err != nil {
		return false
	}

	key, err := uuid.Parse(keyCookie.Value)
	if err != nil {
		return false
	}

	username, err := s.GetUsername(key)
	if err != nil {
		return false
	}

	user, err := s.GetUser(username)
	if err != nil {
		return false
	}

	return user.Role == "admin" || user.Role == "employee"
}


func (s *SessionsDB) IsUserAdmin(r *http.Request) bool {
	keyCookie, err := r.Cookie("session_key")
	if err != nil {
		return false
	}

	key, err := uuid.Parse(keyCookie.Value)
	if err != nil {
		return false
	}

	username, err := s.GetUsername(key)
	if err != nil {
		return false
	}

	user, err := s.GetUser(username)
	if err != nil {
		return false
	}

	return user.Role == "admin"
}