package user

import (
	"fmt"
	"server-ids/internal/auth"
)

// handles buisness logic and calls database

type UserService struct {
	db auth.AuthDB
}

func NewUserService(db auth.AuthDB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) UpdateRole(username string, newRole string) error {
	selectedUser, err := u.db.GetUser(username)
	if err != nil {
		fmt.Printf("Error occured while getting user: %s\n", err)
		return err
	}

	selectedUser.Role = newRole

	err = u.db.UpdateUser(selectedUser)
	if err != nil {
		fmt.Printf("Error occured while updating user: %s\n", err)
		return err
	}
	return nil
}
