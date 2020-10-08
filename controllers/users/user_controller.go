package users

import (
	"fmt"
	"github.com/aman1695/bookStore_users-api/domain/users"
	"github.com/aman1695/bookStore_users-api/services"
	"github.com/aman1695/bookStore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//	TODO: Handle error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id should be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	user, saveErr := services.GetUser(userId)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusOK,user)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id should be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	user.Id = userId
 	result, err := services.UpdateUser(user)
 	if err != nil {
 		c.JSON(err.StatusCode,err)
 		return
	}
	println(result)
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id should be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	result, saveErr := services.DeleteUser(userId)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}