package connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const sqliteFile = "../testing.db"

func SQLiteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(sqliteFile))
	if err != nil {
		panic(err)
	}

	return db
}
