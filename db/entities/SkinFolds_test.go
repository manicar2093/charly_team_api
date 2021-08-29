package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkinFoldsEntity(t *testing.T) {
	skinFold := SkinFolds{
		Subscapular: 12,
		Suprailiac:  12,
		Bicipital:   12,
		Tricipital:  10,
	}

	DB.Create(&skinFold)
	assert.NotEmpty(t, skinFold.ID, "ID should not be empty")
	DB.Delete(&skinFold)
}
