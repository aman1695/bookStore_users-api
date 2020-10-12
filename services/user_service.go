package services

import (
	"github.com/aman1695/bookStore_users-api/domain/users"
	"github.com/aman1695/bookStore_users-api/utils/crypto_utils"
	"github.com/aman1695/bookStore_users-api/utils/date_utils"
	"github.com/aman1695/bookStore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)
type userService struct {}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) (*users.User, *errors.RestErr)
	FindUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err:=user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = "active"
	user.Password = crypto_utils.GetMD5(user.Password)
	return user.Save()
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <=0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	user := users.User{Id: userId}
	return user.Get()

}

func (s *userService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	current, err := current.Get()
	if err != nil {
		return nil, err
	}

		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}

		if user.Status != "" {
			current.Status = user.Status
		}

		if user.Status != "" {
			current.Status = user.Password
		}

	current, err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid id")
	}
	user := users.User{Id: userId}
	return user.Delete()
}

func (s *userService) FindUser(userCriteria string) (users.Users, *errors.RestErr) {
	user := users.User{Status: userCriteria}
	return user.FindUserByStatus()
}