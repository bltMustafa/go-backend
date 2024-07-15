package controllers

import (
	"book-management/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type BookController struct {
	DB *gorm.DB
}

type URLController struct{}

type URLRequest struct {
	URL       string `json:"url" bindind:"required"`
	Operation string `json:"operation" binding:"required"`
}

type URLResponse struct {
	ProcessedURL string `json:"processed_url"`
}

func (ctrl URLController) ProcessURL(c *gin.Context) {
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var processedURL string

	switch req.Operation {
	case "canonical":
		processedURL = cleanURL(req.URL)
	case "redirection":
		processedURL = redirectURL(req.URL)
	case "all":
		processedURL = cleanURL(req.URL)
		processedURL = redirectURL(cleanURL(req.URL))
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operation"})
		return
	}

	c.JSON(http.StatusOK, URLResponse{processedURL})
	println("Response:", c.JSON)
}

func cleanURL(url string) string {
	parsedURL := strings.Split(url, "?")[0]
	if strings.HasSuffix(parsedURL, "/") {
		parsedURL = strings.TrimSuffix(parsedURL, "/")
	}
	return parsedURL
}

func redirectURL(url string) string {
	lowerURL := strings.ToLower(url)
	parsedURL := strings.Split(lowerURL, "www.byfood.com")
	if len(parsedURL) == 2 {
		return "https://www.byfood.com" + parsedURL[1]
	}
	parsedURL = strings.Split(lowerURL, "byfood.com")
	if len(parsedURL) == 2 {
		return "https://byfood.com" + parsedURL[1]
	}
	return "http://www.byfood.com"
}

// Get All Books
func (ctrl BookController) GetBooks(c *gin.Context) {
	var books []models.Book
	ctrl.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Create Book
func (ctrl BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingBook models.Book

	if err := ctrl.DB.Where("title = ?", book.Title).First(&existingBook).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This book already exists"})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author is required"})
		return
	}

	ctrl.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// Get book
func (ctrl BookController) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Update a book
func (ctrl BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	fmt.Print(id)
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not found..!!"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Delete a book

func (ctrl BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not found"})
		return
	}
	ctrl.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
