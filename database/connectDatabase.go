package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	// userName := os.Getenv("DATABASE_USERNAME")
	// password := os.Getenv("DATABASE_PASSWORD")
	// databaseHost := os.Getenv("DATABASE_HOST")
	// databaseName := os.Getenv("DATABASE_NAME")

	//connecting with postgres sql
	// dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, userName, databaseName, password)

	dbURL := os.Getenv("DB_SOURCE")

	//opening and connectiong the databse
	db, err := sql.Open("postgres", dbURL)
	//db.LogMode(true)
	// defer db.Close() // closing the open database connetion after the scope

	//checking for connection
	if err != nil {
		log.Println("Error connecting Database")
		log.Fatal(err)
	}
	//pinging for connection
	err = db.Ping()
	//error checking if the pinf was successful
	if err != nil {
		log.Println("Error pinging Database")
		log.Fatal(err)
	}

	log.Println("Database connection was succesful")
	return db
}
