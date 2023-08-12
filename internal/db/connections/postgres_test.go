package connections_test

import (
	"context"
	"testing"

	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/stretchr/testify/assert"
)

func TestPostgressConn(t *testing.T) {
	conn := connections.PostgressConnection()

	assert.NotNil(t, conn, "connection should not be nil")

	err := conn.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}

}
