package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Author      string `json:"author"`
}

var books []book

func postBook(c *gin.Context) {
	var curBook book
	if err := c.BindJSON(&curBook); err != nil {
		return
	}
	books = append(books, curBook)
	c.IndentedJSON(http.StatusCreated, curBook)

}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	title := c.Param("title")
	for _, book := range books {
		if book.Title == title {
			c.IndentedJSON(http.StatusOK, book)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()
	router.POST("/book", postBook)
	router.GET("/book", getAllBooks)
	router.GET("/book_id:title", getBook)
	router.Run("localhost:8080")
}
