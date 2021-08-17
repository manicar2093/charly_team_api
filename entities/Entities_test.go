package entities

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB
var dns = "../testing.db"

func TestMain(m *testing.M) {

	db, err := gorm.Open(sqlite.Open(dns))
	if err != nil {
		panic(err)
	}
	DB = db
	os.Exit(m.Run())

}
