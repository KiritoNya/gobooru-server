module github.com/KiritoNya/gobooru-server

go 1.17

replace github.com/KiritoNya/gobooru-server/internal/models  => ./internal/models

require (
	github.com/KiritoNya/gobooru-server/models v0.0.0-20220603193808-39f207d1db30
	gorm.io/driver/sqlite v1.3.2
	gorm.io/gorm v1.23.5
)

require (
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
)
