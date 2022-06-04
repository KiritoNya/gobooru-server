package model

import "github.com/jinzhu/gorm"

// TagCategory data model
type TagCategory struct {
	gorm.Model
	Version   int `gorm:"default:1"`
	Name      string
	Color     string
	Usages    int
	Order     int
	IsDefault bool
}
