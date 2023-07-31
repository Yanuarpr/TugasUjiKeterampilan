// main.go

package main

import (
	"movie_api/controllers"
	"movie_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Koneksi ke database MySQL
	dsn := "root:BoedakBager79@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate model Movie ke database
	db.AutoMigrate(&models.Movie{})

	// Inisialisasi koneksi database di kontroler
	controllers.SetupDatabase()

	// Inisialisasi router Gin
	r := gin.Default()

	// Endpoint untuk menampilkan daftar film
	r.GET("/movies", controllers.GetMovies)

	// Endpoint untuk menambahkan film baru
	r.POST("/movies", controllers.CreateMovie)

	// Jalankan server pada port 8080
	r.Run(":8080")
}
