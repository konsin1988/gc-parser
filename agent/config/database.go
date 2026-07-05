package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    
    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  
  connStr := fmt.Sprintf(
      "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=service_info",
      dbHost, dbPort, dbUser, dbPass, dbName,
  )

	DB, err = sql.Open("postgres", connStr)
  if err != nil {
		log.Fatalf("Connection error: %v", err)
  }

  if err = DB.Ping(); err != nil {
		log.Fatalf("Database not allowed: %v", err)
  }
  log.Println("Successfully connected to database")
}
