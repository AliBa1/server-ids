package auth

import (
	"fmt"
	"server-ids/internal/mock"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

// CRUD database

type AuthDBMemory struct {
	Users    []models.User
	Sessions map[uuid.UUID]string
}

func NewAuthDBMemory() *AuthDBMemory {
	return &AuthDBMemory{
		Users:    mock.GetMockUsers(),
		Sessions: map[uuid.UUID]string{},
	}
}

func (db *AuthDBMemory) GetUsers() ([]models.User, error) {
	return db.Users, nil
}

func (db *AuthDBMemory) GetUser(username string) (*models.User, error) {
	for _, u := range db.Users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", username)
}

func (db *AuthDBMemory) CreateUser(user models.User) error {
	db.Users = append(db.Users, user)
	return nil
}

func (db *AuthDBMemory) UpdateUser(user models.User) error {
	// username can't be updated since it's the id
	for i, u := range db.Users {
		if u.Username == user.Username {
			db.Users[i] = user
			return nil
		}
	}
	return fmt.Errorf("user '%s' not found", user.Username)
}

func (db *AuthDBMemory) AddSessionToken(token uuid.UUID, username string) {
	db.Sessions[token] = username
}
