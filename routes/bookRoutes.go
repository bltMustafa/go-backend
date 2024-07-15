package routes

import (
	"book-management/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	bookController := controllers.BookController{DB: db}
	urlController := controllers.URLController{}

	router.GET("/books", bookController.GetBooks)
	router.POST("/books", bookController.CreateBook)
	router.GET("/books/:id", bookController.GetBook)
	router.PUT("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)
	router.POST("/process_url", urlController.ProcessURL)
}
