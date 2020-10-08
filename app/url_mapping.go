package app

import (
	"github.com/aman1695/bookStore_users-api/controllers/ping"
	"github.com/aman1695/bookStore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)

	router.GET("/users/:user_id", users.GetUser)

	router.PUT("/users/:user_id", users.UpdateUser)

	router.DELETE("users/:user_id", users.DeleteUser)
}