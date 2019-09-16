package main

import (
	"log"
)

func main() {

	err := openDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	updateArticles()

}
