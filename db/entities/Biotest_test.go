package entities

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBiotestEntity(t *testing.T) {
	var higherMuscleDensity HigherMuscleDensity
	var lowerMuscleDensity LowerMuscleDensity
	var skinFolds SkinFolds
	var customer User
	var creator User
	var biotest Biotest

	t.Run("Test creation of HigherMuscleDensity instance", func(t *testing.T) {
		higherMuscleDensity = HigherMuscleDensity{
			Neck:                 10.0,
			Shoulders:            42.5,
			Back:                 25.0,
			Chest:                20.0,
			BackChest:            46.3,
			RightRelaxedBicep:    15.3,
			RightContractedBicep: 14.2,
			LeftRelaxedBicep:     12.2,
			LeftContractedBicep:  11.6,
			RightForearm:         12.5,
			LeftForearm:          12.6,
			Wrists:               18.5,
			HighAbdomen:          106.5,
			LowerAbdomen:         106.0,
		}

		DB.Insert(context.Background(), &higherMuscleDensity)
		assert.NotEmpty(t, higherMuscleDensity.ID, "ID should not be empty")
	})

	t.Run("Test creation of LowerMuscleDensity instance", func(t *testing.T) {
		lowerMuscleDensity = LowerMuscleDensity{
			Hips:      45.0,
			RightLeg:  24.6,
			LeftLeg:   24.1,
			RightCalf: 15.3,
			LeftCalf:  15.0,
		}

		DB.Insert(context.Background(), &lowerMuscleDensity)
		assert.NotEmpty(t, lowerMuscleDensity.ID, "ID should not be empty")
	})

	t.Run("Test creation of SkinFolds instance", func(t *testing.T) {
		skinFolds = SkinFolds{
			Subscapular: 32,
			Suprailiac:  100,
			Bicipital:   150,
			Tricipital:  200,
		}

		DB.Insert(context.Background(), &skinFolds)
		assert.NotEmpty(t, skinFolds.ID, "ID should not be empty")
	})

	t.Run("Customer creation for test", func(t *testing.T) {
		customer = User{
			Biotype:     Biotype{ID: 1},
			BoneDensity: BoneDensity{ID: 1},
			Role:        Role{ID: 1},
			Name:        "Test",
			LastName:    "Test",
			Email:       "test1@test.com",
			Birthday:    time.Now(),
		}

		DB.Insert(context.Background(), &customer)
		assert.NotEmpty(t, customer.ID, "ID should not be empty")
	})

	t.Run("Creator creation for test", func(t *testing.T) {
		creator = User{
			Biotype:     Biotype{ID: 1},
			BoneDensity: BoneDensity{ID: 1},
			Role:        Role{ID: 1},
			Name:        "Test",
			LastName:    "Test",
			Email:       "creator_test_1@test.com",
			Birthday:    time.Now(),
		}

		DB.Insert(context.Background(), &creator)
		assert.NotEmpty(t, creator.ID, "ID should not be empty")

	})

	t.Run("Test creation of Biotest instance", func(t *testing.T) {

		biotest = Biotest{
			LowerMuscleDensity:      lowerMuscleDensity,
			SkinFolds:               skinFolds,
			HigherMuscleDensity:     higherMuscleDensity,
			WeightClasification:     WeightClasification{ID: 4},
			Customer:                customer,
			Creator:                 creator,
			Weight:                  250.1,
			Height:                  195,
			BodyFatPercentage:       12.5,
			TotalBodyWater:          100.0,
			BodyMassIndex:           32.66,
			OxygenSaturationInBlood: 91.0,
			Glucose:                 77,
			RestingHeartRate:        70.0,
			MaximumHeartRate:        70.0,
			HeartHealth:             HeartHealth{ID: 1},
			Observations:            "testing observations",
			Recommendations:         "Testing recomentations",
			FrontPicture:            "/path/to/file",
			BackPicture:             "/path/to/file",
			RightSidePicture:        "/path/to/file",
			LeftSidePicture:         "/path/to/file",
			NextEvaluation:          time.Now(),
		}

		DB.Insert(context.Background(), &biotest)
		assert.NotEmpty(t, biotest.ID, "ID should not be empty")
	})

	t.Cleanup(func() {
		ctx := context.Background()
		DB.Delete(ctx, &biotest)
		DB.Delete(ctx, &higherMuscleDensity)
		DB.Delete(ctx, &lowerMuscleDensity)
		DB.Delete(ctx, &skinFolds)
		DB.Delete(ctx, &customer)
		DB.Delete(ctx, &creator)
	})
}
