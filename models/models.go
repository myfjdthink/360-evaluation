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
	sqlite3FilePath := "/tmp/data.sqlite"
	db, err = gorm.Open("sqlite3", sqlite3FilePath)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	fmt.Println("sqlite3 conn " + sqlite3FilePath + " success...")
}
