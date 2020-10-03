package users

import (
	"fmt"
	"github.com/aman1695/bookStore_users-api/utils/errors"
)

var(
	usersDB = make(map [int64]*User)
)
func (user *User) Get() (*User,*errors.RestErr)  {
	result := usersDB[user.Id]
	if result == nil {
		return nil,errors.NewNotFoundError(fmt.Sprintf("user %d not dound", user.Id))
	}
	return usersDB[user.Id],nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewNotFoundError(fmt.Sprintf("user %s already registered", user.Email))
		}
		return errors.NewNotFoundError(fmt.Sprintf("user %d already present", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
