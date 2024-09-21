package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Title       string `json "title"`
	Description string `json "description"`
	Price       int    `json "price"`
	Author      string `json author`
}

var books []book

func postBook(c *gin.Context) {
	var curBook book
	if err := c.BindJson(&curBook); err != nil {
		return
	}
	books = append(books, curBook)
	c.IndentedJson(http.StatusCreated, curBook)

}

func main() {
	router := gin.Default()
	router.POST("/book")
	router.Run("localhost:8080")
}
