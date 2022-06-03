package models

import "github.com/jinzhu/gorm"

// PoolCategory data model
type PoolCategory struct {
	gorm.Model
	Version   int
	Name      string
	Color     string
	Usages    int
	IsDefault bool
}
