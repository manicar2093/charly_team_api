package services

import (
	"os"
	"testing"

	"github.com/manicar2093/charly_team_api/testfunc"
)

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../.env")
	os.Exit(m.Run())
}
