package main

import (
	"github.com/KiritoNya/gobooru-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tagCategory := models.TagCategory{
		Name:      "Prova",
		Color:     "#6544D",
		Order:     1,
		IsDefault: true,
	}

	db.Create(&tagCategory)
}
