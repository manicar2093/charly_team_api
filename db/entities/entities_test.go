package entities

import (
	"os"
	"testing"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/pkg/testfunc"
)

var DB rel.Repository

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../../.env")
	db := connections.PostgressConnection()
	DB = db
	os.Exit(m.Run())

}
