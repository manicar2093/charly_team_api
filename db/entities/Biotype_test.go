package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiotype(t *testing.T) {
	var data []Biotype
	result := DB.Find(&data)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	assert.LessOrEqual(t, 1, len(data), "no biotypes found")
}
