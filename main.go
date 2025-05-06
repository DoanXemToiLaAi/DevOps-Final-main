package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books []Book

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Sample data
	books = append(books, []Book{
		{ID: "1", Title: "Golang Programming", Author: "John Doe", Price: 29.99},
		{ID: "2", Title: "Python Basics", Author: "Jane Smith", Price: 24.99},
		{ID: "3", Title: "JavaScript Advanced", Author: "Mike Johnson", Price: 34.99},
		{ID: "4", Title: "Docker in Practice", Author: "Sarah Williams", Price: 39.99},
		{ID: "5", Title: "DevOps Handbook", Author: "Tom Davis", Price: 44.99},
		{ID: "6", Title: "Microservices Architecture", Author: "Alice Brown", Price: 49.99},
		{ID: "7", Title: "Cloud Computing", Author: "Robert Wilson", Price: 54.99},
		{ID: "8", Title: "Kubernetes Guide", Author: "Emily Clark", Price: 45.99},
		{ID: "9", Title: "API Design Patterns", Author: "David Miller", Price: 32.99},
		{ID: "10", Title: "System Design", Author: "Lisa Anderson", Price: 59.99},
	}...)

	// Define API routes
	api := router.Group("/api")
	{
		api.GET("/books", getBooks)
		api.GET("/books/:id", getBook)
		api.POST("/books", createBook)
		api.PUT("/books/:id", updateBook)
		api.DELETE("/books/:id", deleteBook)
	}

	// Start server
	router.Run(":4200")
	
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, item := range books {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	for index, item := range books {
		if item.ID == id {
			var book Book
			if err := c.ShouldBindJSON(&book); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			book.ID = id
			books = append(books[:index], books[index+1:]...)
			books = append(books, book)
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}