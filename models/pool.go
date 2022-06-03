package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Pool data model
type Pool struct {
	gorm.Model
	Version     int
	Names       pq.StringArray `gorm:"type:text[]"`
	Category    PoolCategory
	Posts       []Post
	Description string
}
