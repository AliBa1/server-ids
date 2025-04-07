package user

import (
	"server-ids/internal/auth"
)

// handles buisness logic and calls database

type UserService struct {
	authDB *auth.AuthDBMemory
}

func NewUserService(authDB *auth.AuthDBMemory) *UserService {
	return &UserService{authDB: authDB}
}

func (u *UserService) UpdateRole(username string, newRole string) error {
	selectedUser, err := u.authDB.GetUser(username)
	if err != nil {
		return err
	}

	selectedUser.Role = newRole

	err = u.authDB.UpdateUser(*selectedUser)
	if err != nil {
		return err
	}
	return nil
}
