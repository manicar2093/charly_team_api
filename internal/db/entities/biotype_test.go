package entities_test

import (
	"context"
	"testing"

	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/stretchr/testify/assert"
)

func TestBiotype(t *testing.T) {
	var data []entities.Biotype
	err := DB.FindAll(context.Background(), &data)

	if err != nil {
		t.Fatal(err)
	}

	assert.LessOrEqual(t, 1, len(data), "no biotypes found")
}
