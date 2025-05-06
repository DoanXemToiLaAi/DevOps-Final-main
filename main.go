package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		{ID: "1", Title: "Exploring the Cosmos", Author: "Neil Starfield", Price: 49.99},
		{ID: "2", Title: "Astrophysics for Beginners", Author: "Luna Skywalker", Price: 39.99},
		{ID: "3", Title: "The Galactic Journey", Author: "Orion Nebula", Price: 59.99},
		{ID: "4", Title: "Black Holes and Beyond", Author: "Celeste Horizon", Price: 44.99},
		{ID: "5", Title: "Interstellar Travel", Author: "Nova Lightyear", Price: 54.99},
		{ID: "6", Title: "The Quantum Universe", Author: "Quasar Photon", Price: 34.99},
		{ID: "7", Title: "Life on Mars", Author: "Red Planet", Price: 29.99},
		{ID: "8", Title: "Saturn's Rings", Author: "Titan Orbital", Price: 24.99},
		{ID: "9", Title: "The Andromeda Mystery", Author: "Galaxy Hunter", Price: 32.99},
		{ID: "10", Title: "Space Exploration 101", Author: "Cosmo Voyager", Price: 19.99},
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