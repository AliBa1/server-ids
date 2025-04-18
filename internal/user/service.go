package user

import "server-ids/internal/models"

// handles buisness logic and calls database

type UserService struct {
	// authDB *auth.AuthDBMemory
	userRepo *UserRepository
}

func NewUserService(ur *UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (u *UserService) UpdateRole(username string, newRole string) error {
	user, err := u.userRepo.GetUser(username)
	if err != nil {
		return err
	}

	user.Role = newRole

	err = u.userRepo.UpdateUser(*user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) CanEditRole(user models.User) bool {
	return user.Role == "admin"
}
