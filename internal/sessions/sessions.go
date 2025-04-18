package sessions

import (
	"database/sql"
	"fmt"
	"net/http"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

type Sessions struct {
	db *sql.DB
}

func NewSessions(db *sql.DB) *Sessions {
	return &Sessions{
		db: db,
	}
}

func (s *Sessions) GetSessionUser(key uuid.UUID) (*models.User, error) {
	query := `
		SELECT u.username, u.password, u.role
		FROM users u
		JOIN sessions s ON s.username = u.username
		WHERE s.key = ?;
	`

	row := s.db.QueryRow(query, key.String())
	var user models.User
	// err := row.Scan(&user.Username, &user.Role, &user.LastLoginDate)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with key does not exist")
		}
		return nil, err
	}

	return &user, nil
}

func (s *Sessions) GetUserFromRequest(r *http.Request) (*models.User, error) {
	rawKey, err := r.Cookie("session_key")
	if err != nil {
		return nil, err
	}

	key, err := uuid.Parse(rawKey.Value)
	if err != nil {
		return nil, err
	}

	user, err := s.GetSessionUser(key)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Sessions) IsUserLoggedIn(r *http.Request) bool {
	user, err := s.GetUserFromRequest(r)
	return user != nil && err == nil
}

func (s *Sessions) IsUserEmployee(r *http.Request) bool {
	user, err := s.GetUserFromRequest(r)
	if err != nil {
		return false
	}

	if user.Role != "admin" && user.Role != "employee" {
		return false
	}

	return true
}

func (s *Sessions) IsUserAdmin(r *http.Request) bool {
	user, err := s.GetUserFromRequest(r)

	if err != nil {
		return false
	}

	if user.Role != "admin" {
		return false
	}

	return true
}
