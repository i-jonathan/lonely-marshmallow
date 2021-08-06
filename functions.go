package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func initDatabase() *gorm.DB {
	connectionLink := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connectionLink), &gorm.Config{})
	if err != nil {
		log.Println("Can't connect to db")
		log.Fatalln(err)
		return nil
	}
	err = db.AutoMigrate(&user{}, &report{})
	if err != nil {
		log.Println("error with auto migration")
		log.Fatalln(err)
		return nil
	}

	return db
}