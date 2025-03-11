package auth

import "server-ids/internal/user"

// handles buisness logic and calls database

type AuthService struct {
	db AuthDB
}

func NewAuthService(db AuthDB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(username string, password string)  {
	// OPTIONAL: hash passwords and compare to hashed

	// check if matches a user
	// if not error handle
	// if password doesn't match error handle
	// else give session token
}

func (s *AuthService) Register(username string, password string)  {
	// create new User
	newUser := user.NewUser(username, password, user.GUEST)
	// add to user db
	s.db.CreateUser(*newUser)
}