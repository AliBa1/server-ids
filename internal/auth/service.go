package auth

// handles buisness logic and calls database

type AuthService struct {
	db AuthDB
}

func NewAuthService(db AuthDB) *AuthService {
	return &AuthService{db: db}
}