package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookRepo struct {
	BaseRepo
}

func NewBookRepo(conn *DBConn) *BookRepo {
	return &BookRepo{BaseRepo{conn: conn}}
}

// func (br *BookRepo) InsertBook(c *gin.Context) {
func InsertBook(c *gin.Context) {
	var book Books
	db.NamedExec(INSERT_BOOK, book)
	err := c.ShouldBindJSON(&book)
	// book.Desccription sudah terisi
	// fmt.Println(book.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "data": book})
}

func ShowBooks(c *gin.Context) {
	var books []Books
	db.Select(&books, GET_BOOKS)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Query row last insert (.Scan)
// db.QueryRow
func DeleteBook(c *gin.Context) {
	var book Books
	err := c.ShouldBindJSON(&book)
	db.NamedExec(DELETE_BOOK, book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "data": book})
}

func UpdateBook(c *gin.Context) {
	var book Books
	err := c.ShouldBindJSON(&book)
	db.NamedExec(UPDATE_BOOK, book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "data": book})
}
