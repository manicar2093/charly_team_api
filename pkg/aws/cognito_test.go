package aws

import (
	"testing"

	"github.com/manicar2093/health_records/pkg/testfunc/asserts"
	"github.com/stretchr/testify/assert"
)

func TestNewCognitoClient(t *testing.T) {
	defer asserts.ShouldNotPanic(t)
	cognitoClient := NewCognitoClient()
	assert.NotNil(t, cognitoClient, "cognito client was not created")
}
