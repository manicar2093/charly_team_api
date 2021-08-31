package connections

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLiteConnection(t *testing.T) {
	sqliteFile = "./connectionTest.db"

	conn := SQLiteConnection()

	assert.NotNil(t, conn, "connection should not be nil")

	err := conn.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Remove(sqliteFile)
	})
}
