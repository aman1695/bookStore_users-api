package users

import (
	"fmt"
	"github.com/aman1695/bookStore_users-api/domain/users"
	"github.com/aman1695/bookStore_users-api/services"
	"github.com/aman1695/bookStore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id should be a number")
		return 0,restErr
	}
	return userId,nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//	TODO: Handle error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, err := GetUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	user, saveErr := services.UserService.GetUser(userId)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusOK,user.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, err := GetUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	user.Id = userId
 	result, err := services.UserService.UpdateUser(user)
 	if err != nil {
 		c.JSON(err.StatusCode,err)
 		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, result.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.StatusCode, userErr)
		return
	}
	result, saveErr := services.UserService.DeleteUser(userId)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, map[string]string{"status":"Deleted"})
}

func Search(c *gin.Context) {
	findCriteria := c.Query("status")
	findCriteria = strings.TrimSpace(strings.ToLower(findCriteria))
	result, saveErr := services.UserService.FindUser(findCriteria)
	if saveErr != nil {
		// ToDo: Handle user creation error
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result.Marshaller(c.GetHeader("X-Public") == "true"))
}