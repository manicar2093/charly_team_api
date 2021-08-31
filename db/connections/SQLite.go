package connections

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var sqliteFile = "../../testing.db"

func SQLiteConnection() rel.Repository {
	adapter, err := sqlite3.Open(sqliteFile)
	if err != nil {
		panic(err)
	}

	return rel.New(adapter)
}
