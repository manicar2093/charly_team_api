package connections_test

import (
	"os"
	"testing"

	"github.com/manicar2093/charly_team_api/pkg/testfunc"
)

func TestMain(m *testing.M) {

	testfunc.LoadEnvFileOrPanic("../../../.env")

	os.Exit(m.Run())
}
