package entities

import (
	"os"
	"testing"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/connections"
)

var DB rel.Repository

func TestMain(m *testing.M) {

	db := connections.SQLiteConnection()
	DB = db
	os.Exit(m.Run())

}
