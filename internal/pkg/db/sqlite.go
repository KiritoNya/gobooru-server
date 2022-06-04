package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqliteDb *gorm.DB

//user:password@tcp(127.0.0.1:3306)/goslayer_db?charset=utf8
const sqliteConnectionString string = "test.db"

func NewSqliteDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqliteDb = db
}
