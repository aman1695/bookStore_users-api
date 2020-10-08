package services

import (
	"github.com/aman1695/bookStore_users-api/domain/users"
	"github.com/aman1695/bookStore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err:=user.Validate(); err != nil {
		return nil, err
	}
	RUser,err:=user.Save()
	if err != nil {
		return nil, err
	}
	return RUser, nil
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

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	if user.Id <= 0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	RUser,err:=user.Update()
	if err != nil {
		return nil, err
	}
	return RUser, nil
}

func DeleteUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	user := users.User{Id: userId}
	RUser,err:=user.Delete()
	if err != nil {
		return nil, err
	}
	return RUser, nil
}