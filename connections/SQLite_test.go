package connections

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLiteConnection(t *testing.T) {
	conn := SQLiteConnection()

	assert.NotNil(t, conn, "connection should not be nil")
	db, err := conn.DB()

	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Remove(sqliteFile)
	})
}
