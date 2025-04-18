package auth

import (
	"fmt"
	"server-ids/internal/models"
	"server-ids/internal/user"

	"github.com/google/uuid"
)

// handles buisness logic and calls database

type AuthService struct {
	authRepo *AuthRepository
	userRepo *user.UserRepository
}

func NewAuthService(ar *AuthRepository, ur *user.UserRepository) *AuthService {
	return &AuthService{authRepo: ar, userRepo: ur}
}

func (s *AuthService) Login(username string, password string) (uuid.UUID, error) {
	// OPTIONAL: hash passwords and compare to hashed

	user, err := s.userRepo.GetUser(username)
	if err != nil {
		return uuid.Nil, err
	}

	if user.Password != password {
		// user attempted login but wrong password

		// add to failed login attempts

		return uuid.Nil, fmt.Errorf("username or password doesn't match")
	}

	token := uuid.New()
	s.authRepo.AddSession(token, *user)
	return token, nil
}

func (s *AuthService) Register(username string, password string) error {
	userExists, _ := s.userRepo.GetUser(username)
	if userExists != nil {
		return fmt.Errorf("username is taken")
	}
	newUser := models.NewUser(username, password, "guest")
	s.userRepo.CreateUser(*newUser)
	return nil
}
