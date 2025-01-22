package main


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)


type book struct {
	ID		string `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Quantity int `json:"quantity"`
}


var books = []book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 3},
	{ID: "2", Title: "A Game of Thrones", Author: "George R.R. Martin", Quantity: 5},
	{ID: "3", Title: "A Clash of Kings", Author: "George R.R. Martin", Quantity: 4},
	{ID: "4", Title: "A Storm of Swords", Author: "George R.R. Martin", Quantity: 3},
	{ID: "5", Title: "A Feast for Crows", Author: "George R.R. Martin", Quantity: 2},
	{ID: "6", Title: "A Dance with Dragons", Author: "George R.R. Martin", Quantity: 4},
}




func getBooks (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}



func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)


	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}



func checkoutBook( c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book out of stock"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}


func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}



func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}





func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}




func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}

