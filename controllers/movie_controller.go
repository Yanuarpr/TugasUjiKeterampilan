// controllers/movie_controller.go

package controllers

import (
	"movie_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Variabel untuk menyimpan koneksi database
var db *gorm.DB

// Fungsi untuk menginisialisasi koneksi database
func SetupDatabase() {
	dsn := "root:BoedakBager79@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate((&models.Movie{}))
}

// Fungsi untuk menampilkan daftar film
func GetMovies(c *gin.Context) {
	var movies []models.Movie
	db.Find(&movies)

	c.JSON(200, gin.H{
		"data": movies,
	})
}

// Fungsi untuk menambahkan film baru
func CreateMovie(c *gin.Context) {
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
	}
	db.Create(&movie)

	c.JSON(200, gin.H{
		"data": movie,
	})
}
