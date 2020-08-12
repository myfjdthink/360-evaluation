package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func Setup() {
	var err error
	sqlite3FilePath := "/tmp/data.sqlite"
	DB, err = gorm.Open("sqlite3", sqlite3FilePath)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("sqlite3 conn " + sqlite3FilePath + " success...")
	initTables()
}

func initTables() {
	log.Println("Initialize tables...")
	DB.AutoMigrate(User{})
	log.Println("Initialize tables Done!")
}
