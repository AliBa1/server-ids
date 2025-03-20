package auth

import (
	"fmt"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

// handles buisness logic and calls database

type AuthService struct {
	db AuthDBMemory
}

func NewAuthService(db AuthDBMemory) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(username string, password string) error {
	// OPTIONAL: hash passwords and compare to hashed

	// check if matches a user
	user, err := s.db.GetUser(username)
	if err != nil {
		return err
	}

	if user.Password != password {
		// user attempted login but wrong password

		// add to failed login attempts

		return fmt.Errorf("username or password doesn't match")
	}

	key := uuid.New()
	s.db.AddLoginKey(key, username)
	return nil
}

func (s *AuthService) Register(username string, password string) error {
	userExists, _ := s.db.GetUser(username)
	if userExists != nil {
		return fmt.Errorf("username is taken")
	}
	newUser := models.NewUser(username, password, "guest")
	s.db.CreateUser(*newUser)
	return nil
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
