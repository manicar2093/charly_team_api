package repositories_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/pkg/testfunc"
)

var (
	DB  rel.Repository
	Ctx context.Context
)

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../../../.env")
	DB = connections.PostgressConnection()
	Ctx = context.Background()
	os.Exit(m.Run())
}
