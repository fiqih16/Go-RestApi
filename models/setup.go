package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_api"))
	if err != nil {
		panic("Failed to connect to database")
	}
	
	db.AutoMigrate(&User{})
	DB = db
}