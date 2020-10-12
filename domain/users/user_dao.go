package users

import (
	"fmt"
	"github.com/aman1695/bookStore_users-api/datasource/postgresql/users_db"
	"github.com/aman1695/bookStore_users-api/utils/errors"
	"log"
	"strings"
)
const(
	queryInsertUser = "INSERT INTO users (first_name,last_name,email,date_created,status,user_pass) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM users Where id=$1;"
	queryUpdateUser = "UPDATE users SET first_name = $1, last_name = $2, email = $3, status = $4 Where id=$5"
	queryDeleteUser = "DELETE FROM users WHERE id=$1"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users Where status=$1;"
 )

var(
	usersDB = make(map [int]*User)
)
func (user *User) Save() (*User,*errors.RestErr) {
	var userId int
	err := users_db.Client.QueryRow(queryInsertUser,user.FirstName,user.LastName,user.Email,user.DateCreated,user.Status,user.Password).Scan(&userId)
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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
	stmt, Perr := users_db.Client.Prepare(queryUpdateUser)
	if Perr != nil {
		stmt.Close()
		return nil, errors.NewBadRequestError(Perr.Error())
	}

	_,getErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.Status,user.Id)
	if getErr != nil {
		stmt.Close()
		return nil,errors.NewInternalServerError(getErr.Error())
	}
	stmt.Close()
	log.Println(user)
	return user,nil
}

func (user *User) Delete() (*User, *errors.RestErr)  {
	user,err := user.Get()
	if err != nil {
		return nil, err
	}
	if _, err := users_db.Client.Exec(queryDeleteUser,user.Id); err != nil {
		return nil, errors.NewInternalServerError("Internal server error")
	}
	return user,nil
}

func (user *User) FindUserByStatus() ([]User, *errors.RestErr) {
	rows, err := users_db.Client.Query(queryFindUserByStatus, user.Status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]User,0)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id,&u.FirstName,&u.LastName,&u.Email,&u.DateCreated,&u.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, u)
		rows.NextResultSet()
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s",user.Status))
	}
	return results, nil
}