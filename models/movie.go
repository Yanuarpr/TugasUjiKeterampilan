// models/movie.go

package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Movie struct {
	gorm.Model
	Title       string
	Description string
	Year        int
}
