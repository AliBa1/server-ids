package auth

import (
	"fmt"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

// handles buisness logic and calls database

type AuthService struct {
	db *AuthDBMemory
}

func NewAuthService(db *AuthDBMemory) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(username string, password string) (uuid.UUID, error) {
	// OPTIONAL: hash passwords and compare to hashed

	// check if matches a user
	user, err := s.db.GetUser(username)
	if err != nil {
		return uuid.Nil, err
	}

	if user.Password != password {
		// user attempted login but wrong password

		// add to failed login attempts

		return uuid.Nil, fmt.Errorf("username or password doesn't match")
	}

	// login successful, give session token
	token := uuid.New()
	s.db.AddSessionToken(token, username)
	return token, nil
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
	users, err := s.db.GetUsers()
	return users, err
}
