package services_test

import (
	"os"
	"testing"

	"github.com/manicar2093/health_records/pkg/testfunc"
)

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../.env")
	os.Exit(m.Run())
}
