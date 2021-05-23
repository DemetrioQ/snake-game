package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func connection() *sql.DB{
	dbinfo := getEnvVariable("DbInfo")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}else{
		log.Print("Database Connected")
	}
	return db
}


func getEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}