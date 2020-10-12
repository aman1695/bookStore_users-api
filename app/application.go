package app

import (
	"github.com/aman1695/bookStore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)
func StartApplication(){
	mapURLs()
	logger.Info("about to start the application.....")
	router.Run(":8080")
}

