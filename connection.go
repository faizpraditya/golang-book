package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DBConn struct {
	Db *sqlx.DB
}

func connectDB() *sqlx.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	datasourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Connect to databasae
	db, err := sqlx.Connect("pgx", datasourceName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connect to database!")
	}
	return db
}

// func (d *DBConn) Close() {
// 	db, err := d.Db.DB()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = db.Close()

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
