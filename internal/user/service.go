package user

import (
	"server-ids/internal/auth"
)

// handles buisness logic and calls database

type UserService struct {
	db auth.AuthDBMemory
}

func NewUserService(db auth.AuthDBMemory) *UserService {
	return &UserService{db: db}
}

func (u *UserService) UpdateRole(username string, newRole string) error {
	selectedUser, err := u.db.GetUser(username)
	if err != nil {
		return err
	}

	selectedUser.Role = newRole

	err = u.db.UpdateUser(*selectedUser)
	if err != nil {
		return err
	}
	return nil
}
