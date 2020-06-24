package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	err := openDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	updateArticles()
}
