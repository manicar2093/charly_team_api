package entities

import (
	"os"
	"testing"

	"github.com/manicar2093/charly_team_api/db/connections"

	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {

	db := connections.SQLiteConnection()
	DB = db
	os.Exit(m.Run())

}
