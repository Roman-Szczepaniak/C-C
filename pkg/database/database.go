package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=bd11 password=bd11 port=5432 sslmode=disable" // https://pkg.go.dev/gorm.io/driver/postgres#section-readme
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db
}
