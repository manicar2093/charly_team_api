package entities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestSkinFoldsEntity(t *testing.T) {
	skinFold := SkinFolds{
		Subscapular: null.IntFrom(12),
		Suprailiac:  null.IntFrom(12),
		Bicipital:   null.IntFrom(12),
		Tricipital:  null.IntFrom(10),
	}

	ctx := context.Background()

	DB.Insert(ctx, &skinFold)
	assert.NotEmpty(t, skinFold.ID, "ID should not be empty")
	DB.Delete(ctx, &skinFold)
}
