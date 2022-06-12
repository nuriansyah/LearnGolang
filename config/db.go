package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"log"

	"gorm.io/gorm"
)

func ConDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection is Failure, %v\n", err.Error())
	}
	fmt.Println("DB Connected")
	return db
}
