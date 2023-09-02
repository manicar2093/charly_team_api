package connections_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// testfunc.LoadEnvFileOrPanic("../../../.env")

	os.Exit(m.Run())
}
