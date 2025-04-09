package auth

import (
	"fmt"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
)

// CRUD database

type AuthDBMemory struct {
	SessionsDB *sessions.SessionsDB
}

func NewAuthDBMemory(sDB *sessions.SessionsDB) *AuthDBMemory {
	return &AuthDBMemory{
		SessionsDB: sDB,
	}
}

func (db *AuthDBMemory) GetUsers() ([]models.User, error) {
	return db.SessionsDB.Users, nil
}

func (db *AuthDBMemory) GetUser(username string) (*models.User, error) {
	for _, u := range db.SessionsDB.Users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", username)
}

func (db *AuthDBMemory) CreateUser(user models.User) error {
	db.SessionsDB.Users = append(db.SessionsDB.Users, user)
	return nil
}

func (db *AuthDBMemory) UpdateUser(user models.User) error {
	// username can't be updated since it's the id
	for i, u := range db.SessionsDB.Users {
		if u.Username == user.Username {
			db.SessionsDB.Users[i] = user
			return nil
		}
	}
	return fmt.Errorf("user '%s' not found", user.Username)
}
