package app

import (
	"github.com/aman1695/bookStore_users-api/controllers/ping"
	"github.com/aman1695/bookStore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)

	router.GET("/users/:user_id", users.Get)

	router.PUT("/users/:user_id", users.Update)

	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)
}