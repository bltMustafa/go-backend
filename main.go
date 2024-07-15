package main

import (
	"book-management/models"
	"book-management/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {
	// Veritabanı bağlantısı
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Veritabanı migrasyonu
	models.MigrateBooks(DB)

	// Gin router
	r := gin.Default()

	// Kitap route'ları
	routes.RegisterRoutes(r, DB)

	// Sunucu
	r.Run(":7777")
}
