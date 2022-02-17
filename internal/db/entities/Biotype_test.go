package entities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiotype(t *testing.T) {
	var data []Biotype
	err := DB.FindAll(context.Background(), &data)

	if err != nil {
		t.Fatal(err)
	}

	assert.LessOrEqual(t, 1, len(data), "no biotypes found")
}
