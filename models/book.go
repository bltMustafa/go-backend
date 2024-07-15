package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func MigrateBooks(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	fmt.Println("Book table created!")
}
