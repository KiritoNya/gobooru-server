package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Tag data model
type Tag struct {
	gorm.Model
	Version      int
	Names        pq.StringArray `gorm:"type:text[]"`
	CategoryID   int
	Category     TagCategory
	Implications []Tag `gorm:"many2many:tag_implications"`
	Suggestions  []Tag `gorm:"many2many:tag_suggestions"`
	Usages       int
	Description  string
}
