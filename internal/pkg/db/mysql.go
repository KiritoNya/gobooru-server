package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysqldb *gorm.DB

//user:password@tcp(127.0.0.1:3306)/goslayer_db?charset=utf8
const mysqlconnstring string = ""

func init() {
	if len(mysqlconnstring) > 0 {

		// database connection
		dbconn, err := gorm.Open(mysql.Open(mysqlconnstring), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		// create database connection pool
		sqlDB, err := dbconn.DB()
		if err != nil {
			panic(err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(20)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(200)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)

		mysqldb = dbconn
	}

	log.Println("mysql is opened successfully")
}

//NewMysql returns mysql connection instance
func NewMysql() *gorm.DB {
	return mysqldb
}
