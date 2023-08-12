package aws

import (
	"testing"

	"github.com/manicar2093/health_records/pkg/testfunc/asserts"
	"github.com/stretchr/testify/assert"
)

func TestAWSSession(t *testing.T) {

	t.Run("should connect correctly", func(t *testing.T) {
		defer asserts.ShouldNotPanic(t)
		got := GetAWSSession()
		assert.NotNil(t, got, "aws session was not created")
	})

}
