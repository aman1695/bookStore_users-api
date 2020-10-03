package services

import (
	"github.com/aman1695/bookStore_users-api/domain/users"
	"github.com/aman1695/bookStore_users-api/urls/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err:=user.Validate(); err != nil {
		return nil, err
	}
	if err:=user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <=0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	user := users.User{Id: userId}
	resUser,err:=user.Get()
	if err != nil {
		return resUser,err
	}
	return resUser,nil

}
