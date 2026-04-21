package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	log.Println("Start create a database connection...")
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Cannot find any env file: %w", err)
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "student_manager"

	// sql.Open will return a pointer for db and an error
	db, error := sql.Open("mysql", cfg.FormatDSN())

	// Handle error
	/*
		- %w is special in Go—it’s used for wrapping errors, not just formatting.
		- It uses with fmt.Errorf()
		--> To wrap original error with new message
	*/
	if error != nil {
		return nil, fmt.Errorf("Cannot get DB connection: %w", error)
	}

	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error Ping: %w", error)
	}

	fmt.Println("Database connected!!!")
	return db, nil
}