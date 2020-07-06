//go:generate goversioninfo -64 -icon=icon.ico -platform-specific=true

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	// SQL wrapper
	SQL gorm.DB

	// Database info
)

type Test struct {
	ID   int
	Name string
}

func Connect() {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		log.Print(err)
	}
	err = automigrate(db)

	SQL = *db

}

func automigrate(db *gorm.DB) (err error) {
	mg := db.AutoMigrate(&Test{})
	return mg.Error
}

func main() {
	Connect()
}
