package entities_test

import (
	"os"
	"testing"

	"github.com/go-rel/rel"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/pkg/testfunc"
)

var DB rel.Repository

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../../../.env")
	db := connections.PostgressConnection()
	DB = db
	os.Exit(m.Run())

}
