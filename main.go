package main

import (
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/database"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	router.UserRoutes(r)
	router.PhotoRoutes(r)

	r.Run()

}
