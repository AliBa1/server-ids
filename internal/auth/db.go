package auth

import (
	"fmt"
	"server-ids/internal/mock"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
)

// CRUD database

type AuthDBMemory struct {
	Users      []models.User
	SessionsDB *sessions.SessionsDB
}

func NewAuthDBMemory(sDB *sessions.SessionsDB) *AuthDBMemory {
	return &AuthDBMemory{
		Users:      mock.GetMockUsers(),
		SessionsDB: sDB,
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

// func (db *AuthDBMemory) AddSessionToken(token uuid.UUID, username string) {
// 	db.Sessions[token] = username
// }
