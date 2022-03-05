package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/stdlib"
)

// Kalau bisa jangan, karena kalau testing db connect kepanggil terus
// Pake baserepo, bookrepo
// Pake yang di enigmaschool
var db = connectDB()

func main() {
	router := gin.Default()
	defer db.Close()

	router.POST("/book", InsertBook)
	router.GET("/book", ShowBooks)
	router.DELETE("/book", DeleteBook)
	router.PATCH("/book", UpdateBook)
	router.Run(os.Getenv("API_HOST"))
}
