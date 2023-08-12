package entities_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-rel/rel/where"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestBiotestEntity(t *testing.T) {
	var higherMuscleDensity entities.HigherMuscleDensity
	var lowerMuscleDensity entities.LowerMuscleDensity
	var skinFolds entities.SkinFolds
	var customer entities.User
	var creator entities.User
	var biotest entities.Biotest

	faker := faker.New()

	customer = entities.User{
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		UserUUID:      faker.UUID().V4(),
		Name:          "Test",
		LastName:      "Test",
		Email:         faker.Internet().Email(),
		Birthday:      time.Now(),
	}

	DB.Insert(context.Background(), &customer)
	assert.NotEmpty(t, customer.ID, "ID should not be empty")

	creator = entities.User{
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		UserUUID:      faker.UUID().V4(),
		Name:          "Test",
		LastName:      "Test",
		Email:         faker.Internet().Email(),
		Birthday:      time.Now(),
	}

	DB.Insert(context.Background(), &creator)
	assert.NotEmpty(t, creator.ID, "ID should not be empty")

	biotest = entities.Biotest{
		HigherMuscleDensity:     higherMuscleDensity,
		LowerMuscleDensity:      lowerMuscleDensity,
		SkinFolds:               skinFolds,
		WeightClasificationID:   4,
		HeartHealthID:           1,
		CustomerID:              customer.ID,
		CreatorID:               creator.ID,
		BiotestUUID:             faker.UUID().V4(),
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

	err := DB.Insert(context.Background(), &biotest)
	assert.Nil(t, err, "Unexpected error")

	higherMuscleDensity = entities.HigherMuscleDensity{
		BiotestID:            &biotest.ID,
		Neck:                 null.FloatFrom(10.0),
		Shoulders:            null.FloatFrom(42.5),
		Back:                 null.FloatFrom(25.0),
		Chest:                null.FloatFrom(20.0),
		BackChest:            null.FloatFrom(46.3),
		RightRelaxedBicep:    null.FloatFrom(15.3),
		RightContractedBicep: null.FloatFrom(14.2),
		LeftRelaxedBicep:     null.FloatFrom(12.2),
		LeftContractedBicep:  null.FloatFrom(11.6),
		RightForearm:         null.FloatFrom(12.5),
		LeftForearm:          null.FloatFrom(12.6),
		Wrists:               null.FloatFrom(18.5),
		HighAbdomen:          null.FloatFrom(106.5),
		LowerAbdomen:         null.FloatFrom(106.0),
	}

	DB.Insert(context.Background(), &higherMuscleDensity)
	assert.NotEmpty(t, higherMuscleDensity.ID, "ID should not be empty")

	lowerMuscleDensity = entities.LowerMuscleDensity{
		BiotestID: &biotest.ID,
		Hips:      null.FloatFrom(45.0),
		RightLeg:  null.FloatFrom(24.6),
		LeftLeg:   null.FloatFrom(24.1),
		RightCalf: null.FloatFrom(15.3),
		LeftCalf:  null.FloatFrom(15.0),
	}

	DB.Insert(context.Background(), &lowerMuscleDensity)
	assert.NotEmpty(t, lowerMuscleDensity.ID, "ID should not be empty")

	skinFolds = entities.SkinFolds{
		BiotestID:   &biotest.ID,
		Subscapular: null.IntFrom(32),
		Suprailiac:  null.IntFrom(100),
		Bicipital:   null.IntFrom(150),
		Tricipital:  null.IntFrom(200),
	}

	DB.Insert(context.Background(), &skinFolds)
	assert.NotEmpty(t, skinFolds.ID, "ID should not be empty")

	var biotestFound entities.Biotest
	DB.Find(context.Background(), &biotestFound, where.Eq("id", biotest.ID))
	assert.NotEmpty(t, biotestFound.Customer.ID, "Customer was not loaded correctly")
	assert.NotEmpty(t, biotestFound.ID, "ID should not be empty")
	assert.NotEmpty(t, biotestFound.HigherMuscleDensity.ID, "ID should not be empty")
	assert.NotEmpty(t, biotestFound.SkinFolds.ID, "ID should not be empty")
	assert.NotEmpty(t, biotestFound.LowerMuscleDensity.ID, "ID should not be empty")

	biotestAsBytes, _ := json.Marshal(biotestFound)
	var biotestFromJson entities.Biotest

	json.Unmarshal(biotestAsBytes, &biotestFromJson)

	biotestFromJson.CorporalAge = 50
	err = DB.Update(context.Background(), &biotestFromJson)

	DB.Find(context.Background(), &biotestFound, where.Eq("id", biotest.ID))

	assert.Equal(t, biotestFound.CorporalAge, biotestFromJson.CorporalAge, "update was not executed correctly")

	if err != nil {
		t.Fatal(err)
	}

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
