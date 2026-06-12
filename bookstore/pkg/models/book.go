package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book

	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

func (b *Book) UpdateBook(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	book.Name = b.Name
	book.Author = b.Author
	book.Description = b.Description
	book.Publication = b.Publication
	db.Save(&book)
	return &book
}
