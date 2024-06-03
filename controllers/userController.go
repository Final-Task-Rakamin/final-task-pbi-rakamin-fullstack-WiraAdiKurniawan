package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/database"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/helpers"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get req body data
	var body models.User
	var validate = validator.New()

	if err := c.ShouldBind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body",
		})
		return
	}

	// validator
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All data is required",
		})
		return
	}

	// validator unique email
	var existingUser models.User
	if err := database.DB.Where("Email = ?", body.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
		})
		return
	}

	// hash password
	hash, err := helpers.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	//  post data
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create user",
		"payload": user,
	})
}

// login
func Login(c *gin.Context) {
	var body models.User

	if err := c.ShouldBind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body",
		})
		return
	}

	// find user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email or password",
		})
		return
	}

	// compare hash password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email or password",
		})
		return
	}

	// generate and signed jwt
	tokenString, err := helpers.GenerateToken(uint64(user.ID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "succesful login",
	})
}

// get user by token
func GetUserLogin(c *gin.Context) {
	user, _ := c.Get("user")

	// Cek tipe data pengguna
	if userObj, ok := user.(models.User); ok {
		// Memuat data foto-foto terkait dengan pengguna
		database.DB.Model(&userObj).Association("Photos").Find(&userObj.Photos)

		c.JSON(http.StatusOK, gin.H{
			"user": userObj,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

// update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body models.User
	var validate = validator.New()

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body",
		})
		return
	}

	// validator
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All required",
		})
		return
	}

	// find user
	var user models.User
	result := database.DB.First(&user, id)

	// check user
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("User with ID %s not found", id),
		})
		return
	}

	// hash password
	hash, err := helpers.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	database.DB.Model(&user).Updates(models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	})

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Account with ID %s has been updated", id),
	})

}

// delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	result := database.DB.First(&user, id)

	// check user
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("User with ID %s not found", id),
		})
		return
	}

	database.DB.Delete(&user, id)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Account with ID %s has been deleted", id),
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// get data all data account
func PostsIndex(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)

	for i := range user {
		database.DB.Model(&user[i]).Association("Photos").Find(&user[i].Photos)
	}

	c.JSON(200, gin.H{
		"payload": user,
	})
}

func GetPhoto(c *gin.Context) {
	var photo []models.Photo
	database.DB.Find(&photo)

	c.JSON(200, gin.H{
		"payload": photo,
	})
}
