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
	CategoryID  int
	Category    PoolCategory
	Posts       []Post `gorm:"many2many:post_pools"`
	Description string
}
