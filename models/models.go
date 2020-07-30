package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB

type Model struct {
}

func Setup() {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	fmt.Println("sqlite3 conn success...")
}
