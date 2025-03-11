package auth

import (
	"fmt"
	"server-ids/internal/user"
)

// CRUD database

type AuthDB interface {
	GetAllUsers() ([]user.User, error)
	GetUser(username string) (user.User, error)
	CreateUser(user user.User)
}

type AuthDBMemory struct {
	Users []user.User
}

func NewAuthDBMemory() *AuthDBMemory {
	return &AuthDBMemory{
		Users: []user.User{},
	}
}

func (db *AuthDBMemory) GetAllUsers() ([]user.User, error) {
	return db.Users, nil
}

func (db *AuthDBMemory) GetUser(username string) (user.User, error) {
	for _, u := range db.Users {
		if u.Username == username {
			return u, nil
		}
	}
	return user.User{}, fmt.Errorf("user '%s' not found", username)
}

func (db *AuthDBMemory) CreateUser(user user.User) {
	db.Users = append(db.Users, user)
}