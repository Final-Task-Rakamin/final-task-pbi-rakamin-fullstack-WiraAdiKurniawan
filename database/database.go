package database

import (
	"fmt"

	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/helpers"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	User := models.User{}
	helpers.LoadEnv()
	dbHost := helpers.GetEnv("DB_HOST")
	dbUser := helpers.GetEnv("DB_USER")
	dbPassword := helpers.GetEnv("DB_PASSWORD")
	dbName := helpers.GetEnv("DB_NAME")
	dbPort := helpers.GetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	db.AutoMigrate(&User, &models.Photo{})
}
