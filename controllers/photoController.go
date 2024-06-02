package controllers

import (
	"fmt"
	"net/http"

	"github.com/adikrnwn171/database"
	"github.com/adikrnwn171/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddPhoto(c *gin.Context) {
	// id from jwt
	user, _ := c.Get("user")

	var validate = validator.New()

	// Mengambil objek models.User dari konteks
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to convert user to User model",
		})
		return
	}

	// Mendapatkan req body data
	var body models.Photo

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// validasi data
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All data is required",
		})
		return
	}

	// Buat instance Photo dengan UserID yang sesuai
	photo := models.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoURL: body.PhotoURL,
		UserID:   userObj.ID, // Menggunakan userID dari token JWT
	}

	// Simpan foto ke database
	result := database.DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create photo",
		})
		return
	}

	// Setelah foto berhasil dibuat, tambahkan foto tersebut ke slice Photos objek User
	userObj.Photos = append(userObj.Photos, photo)

	// Simpan objek User yang telah diperbarui ke database
	database.DB.Save(&userObj)

	c.JSON(http.StatusOK, gin.H{
		"photo": userObj,
	})
}

// Menghapus foto berdasarkan ID (hanya pemilik yang bisa)
func DeletePhoto(c *gin.Context) {
	photoID := c.Param("id")

	var photo models.Photo

	// Mendapatkan foto dari database berdasarkan photoID
	result := database.DB.First(&photo, photoID)

	// Check photo ID
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Photo with ID %s not found", photoID),
		})
		return
	}

	// Mendapatkan ID pengguna dari token JWT
	user, _ := c.Get("user")

	// Mengambil objek models.User dari konteks
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to convert user to User model",
		})
		return
	}
	userID := photo.UserID

	// Memeriksa apakah pengguna adalah pemilik foto yang ingin dihapus
	if userID != userObj.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("You don't have authorization to delete photo with ID %s", photoID),
		})
		return
	}

	// Hapus foto dari database
	database.DB.Delete(&photo, photoID)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with ID %s has been deleted", photoID),
	})
}

// Memperbarui foto berdasarkan ID (hanya pemilik yang bisa)
func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("id")

	var body struct {
		Title    string
		Caption  string
		PhotoURL string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	var photo models.Photo
	database.DB.First(&photo, photoID)

	result := database.DB.First(&photo, photoID)

	// Check photo ID
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Photo with ID %s not found", photoID),
		})
		return
	}

	// Mendapatkan ID pengguna dari token JWT
	user, _ := c.Get("user")

	// Mengambil objek models.User dari konteks
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to convert user to User model",
		})
		return
	}

	// Memeriksa apakah pemilik foto adalah pengguna yang sedang login
	if photo.UserID != userObj.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("You don't have authorization to update photo with ID %s", photoID),
		})
		return
	}

	// Perbarui atribut foto
	database.DB.Model(&photo).Updates(models.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoURL: body.PhotoURL,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with ID %s has been updated", photoID),
	})
}
