package connections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgressConn(t *testing.T) {
	conn := PostgressConnection()

	assert.NotNil(t, conn, "connection should not be nil")
	db, err := conn.DB()

	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

}
