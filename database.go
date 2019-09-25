package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func openDatabaseConnection() error {

	var err error

	db, err = gorm.Open("postgres", os.Getenv("POCKET_DATABASE_URL"))
	if err != nil {
		return err
	}

	db.AutoMigrate(&Article{})

	return err

}
