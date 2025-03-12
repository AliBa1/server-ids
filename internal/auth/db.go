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
	mockUsers := []user.User{
		*user.NewUser("funguy123", "admin12345", "admin"),
		*user.NewUser("bossman", "emp12345", "employee"),
		*user.NewUser("grumpy", "guest12345", "guest"),
		*user.NewUser("jpearson", "guest12345", "guest"),
		*user.NewUser("fredrick5", "guest12345", "guest"),
		*user.NewUser("ballhoggary", "emp12345", "employee"),
		*user.NewUser("erick", "admin12345", "admin"),
		*user.NewUser("barrylarry", "emp12345", "employee"),
		*user.NewUser("twotthree", "guest12345", "guest"),
		*user.NewUser("yap", "guest12345", "guest"),
		*user.NewUser("boardman45", "guest12345", "guest"),
		*user.NewUser("1819twenty", "emp12345", "employee"),
		*user.NewUser("opi", "guest12345", "guest"),
		*user.NewUser("patrick", "guest12345", "guest"),
		*user.NewUser("fred111", "guest12345", "guest"),
		*user.NewUser("secure21", "guest12345", "guest"),
	}
	return &AuthDBMemory{
		// Users: []user.User{},
		Users: mockUsers,
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
