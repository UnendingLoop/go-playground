package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User - struct for ORM
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Age   uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("sandbox.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to DB:", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Couldn't migrate:", err)
	}

	log.Println("Successfully connected to DB!")
}
