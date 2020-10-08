package users

import (
	"fmt"
	"github.com/aman1695/bookStore_users-api/datasource/postgresql/users_db"
	"github.com/aman1695/bookStore_users-api/utils/errors"
	"log"
	"strings"
	"time"
)
const(
	queryInsertUser = "INSERT INTO users (first_name,last_name,email,date_created) VALUES ($1,$2,$3,$4) RETURNING id;"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM users Where id=$1;"
	queryUpdateUser = "UPDATE users SET first_name = $1, last_name = $2, email = $3 Where id=$4"
	queryDeleteUser = "DELETE FROM users WHERE id=$1"
)

var(
	usersDB = make(map [int]*User)
)
func (user *User) Save() (*User,*errors.RestErr) {
	now := time.Now()
	user.DateCreated=now.Format("2006-01-02T15:05:07Z")
	var userId int
	err := users_db.Client.QueryRow(queryInsertUser,user.FirstName,user.LastName,user.Email,user.DateCreated).Scan(&userId)
	if err != nil {
		if strings.Contains(err.Error(),"users_email_key") {
			log.Println(err)
			return nil,errors.NewInternalServerError(fmt.Sprintf("Email %s already registered",user.Email))
		}
		return nil,errors.NewInternalServerError(err.Error())
	}
	user.Id = int64(userId)
	return user,nil
}

func (user *User) Get() (*User,*errors.RestErr)  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return nil,errors.NewInternalServerError(fmt.Sprintf("Invalid Reques: %s",err.Error()))
	}

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(),"no rows in result set") {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user id %d not found!!",user.Id))
		}
		log.Print("error when trying to get user by id : ", getErr.Error())
		return nil,errors.NewInternalServerError(fmt.Sprintf("error when trying to get user by id %d : %s",user.Id,getErr.Error()))
	}
	stmt.Close()
	return user,nil
}

func (user *User) Update() (*User,*errors.RestErr)  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return nil,errors.NewInternalServerError(fmt.Sprintf("Invalid Request: %s",err.Error()))
	}
	var old_user User
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&old_user.Id, &old_user.FirstName, &old_user.LastName, &old_user.Email, &old_user.DateCreated); getErr != nil {
		stmt.Close()
		if strings.Contains(getErr.Error(),"no rows in result set") {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user id %d not found!!",user.Id))
		}
		log.Print("error when trying to get user by id : ", getErr.Error())
		return nil,errors.NewInternalServerError(fmt.Sprintf("error when trying to get user by id %d : %s",user.Id,getErr.Error()))
	}
	stmt.Close()
	if old_user.Email != user.Email && user.Email != ""{
		old_user.Email = user.Email
	}
	if old_user.FirstName != user.FirstName && user.FirstName != ""{
		old_user.FirstName = user.FirstName
	}
	if old_user.LastName != user.LastName && user.LastName != ""{
		old_user.LastName = user.LastName
	}
	getErr := users_db.Client.QueryRow(queryUpdateUser,old_user.FirstName,old_user.LastName,old_user.Email,old_user.Id).Err()
	if getErr != nil {
		return nil,errors.NewInternalServerError(getErr.Error())
	}
	return &old_user,nil
}

func (user *User) Delete() (*User, *errors.RestErr)  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Invalid Request: %s", err.Error()))
	}
	result := stmt.QueryRow(user.Id)
	if deleteErr:= result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); deleteErr != nil {
		stmt.Close()
		return nil, errors.NewNotFoundError(fmt.Sprintf("user id: %d not exist %s",user.Id,deleteErr.Error()))
	}
	if _, err := users_db.Client.Exec(queryDeleteUser,user.Id); err != nil {
		return nil, errors.NewInternalServerError("Internal server error")
	}
	stmt.Close()
	return user,nil
}