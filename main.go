package main

import (
	"github.com/adikrnwn171/database"
	"github.com/adikrnwn171/router"
	"github.com/gin-gonic/gin"
)

// @title My API
// @version 1.0

func main() {
	r := gin.Default()
	database.ConnectDB()

	router.UserRoutes(r)
	router.PhotoRoutes(r)

	r.Run()
}
