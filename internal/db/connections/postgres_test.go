package connections

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgressConn(t *testing.T) {
	conn := PostgressConnection()

	assert.NotNil(t, conn, "connection should not be nil")

	err := conn.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}

}
