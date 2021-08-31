package repositories

import (
	"context"
	"os"
	"testing"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/connections"
)

var (
	DB  rel.Repository
	Ctx context.Context
)

func TestMain(m *testing.M) {
	DB = connections.SQLiteConnection()
	Ctx = context.Background()
	os.Exit(m.Run())
}
