package entities

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
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
			BiotypeID:     null.IntFrom(1),
			BoneDensityID: null.IntFrom(1),
			RoleID:        1,
			Name:          "Test",
			LastName:      "Test",
			Email:         "test1@test.com",
			Birthday:      time.Now(),
		}

		DB.Insert(context.Background(), &customer)
		assert.NotEmpty(t, customer.ID, "ID should not be empty")
	})

	t.Run("Creator creation for test", func(t *testing.T) {
		creator = User{
			BiotypeID:     null.IntFrom(1),
			BoneDensityID: null.IntFrom(1),
			RoleID:        1,
			Name:          "Test",
			LastName:      "Test",
			Email:         "creator_test_1@test.com",
			Birthday:      time.Now(),
		}

		DB.Insert(context.Background(), &creator)
		assert.NotEmpty(t, creator.ID, "ID should not be empty")

	})

	t.Run("Test creation of Biotest instance", func(t *testing.T) {

		biotest = Biotest{
			LowerMuscleDensityID:    lowerMuscleDensity.ID,
			SkinFoldsID:             skinFolds.ID,
			HigherMuscleDensityID:   higherMuscleDensity.ID,
			WeightClasificationID:   4,
			HeartHealthID:           1,
			CustomerID:              customer.ID,
			CreatorID:               creator.ID,
			Weight:                  250.1,
			Height:                  195,
			BodyFatPercentage:       12.5,
			TotalBodyWater:          100.0,
			BodyMassIndex:           32.66,
			OxygenSaturationInBlood: 91.0,
			Glucose:                 null.FloatFrom(77.0),
			RestingHeartRate:        null.FloatFrom(70.0),
			MaximumHeartRate:        null.FloatFrom(70.0),
			Observations:            null.StringFrom("testing observations"),
			Recommendations:         null.StringFrom("Testing recomentations"),
			FrontPicture:            null.StringFrom("/path/to/file"),
			BackPicture:             null.StringFrom("/path/to/file"),
			RightSidePicture:        null.StringFrom("/path/to/file"),
			LeftSidePicture:         null.StringFrom("/path/to/file"),
			NextEvaluation:          null.TimeFrom(time.Now()),
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
