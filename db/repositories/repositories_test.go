package repositories

import (
	"os"
	"testing"

	"github.com/manicar2093/charly_team_api/db/connections"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	DB = connections.SQLiteConnection()
	os.Exit(m.Run())
}
