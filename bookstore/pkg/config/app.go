package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(
		"mysql",
		"root:secret@tcp(127.0.0.1:3306)/bookstore_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
