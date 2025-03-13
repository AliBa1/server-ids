package auth

import (
	"fmt"
	"server-ids/internal/mock"
	"server-ids/internal/models"
)

// CRUD database

type AuthDB interface {
	GetAllUsers() ([]models.User, error)
	GetUser(username string) (models.User, error)
	CreateUser(user models.User)
	UpdateUser(user models.User) error
}

type AuthDBMemory struct {
	Users []models.User
}

func NewAuthDBMemory() *AuthDBMemory {
	// mockUsers := []models.User{
	// 	*models.NewUser("funguy123", "admin12345", "admin"),
	// 	*models.NewUser("bossman", "emp12345", "employee"),
	// 	*models.NewUser("grumpy", "guest12345", "guest"),
	// 	*models.NewUser("jpearson", "guest12345", "guest"),
	// 	*models.NewUser("fredrick5", "guest12345", "guest"),
	// 	*models.NewUser("ballhoggary", "emp12345", "employee"),
	// 	*models.NewUser("erick", "admin12345", "admin"),
	// 	*models.NewUser("barrylarry", "emp12345", "employee"),
	// 	*models.NewUser("twotthree", "guest12345", "guest"),
	// 	*models.NewUser("yap", "guest12345", "guest"),
	// 	*models.NewUser("boardman45", "guest12345", "guest"),
	// 	*models.NewUser("1819twenty", "emp12345", "employee"),
	// 	*models.NewUser("opi", "guest12345", "guest"),
	// 	*models.NewUser("patrick", "guest12345", "guest"),
	// 	*models.NewUser("fred111", "guest12345", "guest"),
	// 	*models.NewUser("secure21", "guest12345", "guest"),
	// }
	return &AuthDBMemory{
		// Users: []models.User{},
		// Users: mockUsers,
		Users: mock.GetMockUsers(),
	}
}

func (db *AuthDBMemory) GetAllUsers() ([]models.User, error) {
	return db.Users, nil
}

func (db *AuthDBMemory) GetUser(username string) (models.User, error) {
	for _, u := range db.Users {
		if u.Username == username {
			return u, nil
		}
	}
	return models.User{}, fmt.Errorf("user '%s' not found", username)
}

func (db *AuthDBMemory) CreateUser(user models.User) {
	db.Users = append(db.Users, user)
}

func (db *AuthDBMemory) UpdateUser(user models.User) error {
	for i, u := range db.Users {
		if u.Username == user.Username {
			db.Users[i] = user
			return nil
		}
	}
	return fmt.Errorf("user '%s' not found", user.Username)
}
