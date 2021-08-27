package services

import (
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
)

func TestCreateUser(t *testing.T) {

	t.Skip("No correr aun este test")

	providerMock := &mocks.CongitoClient{}
	dbMock := &mocks.Repository{}
	passGenMock := &mocks.PassGen{}

	userService := UserServiceCognito{
		provider: providerMock,
		db:       dbMock,
		passGen:  passGenMock,
	}

	userRequuest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "testing",
		Email:    "manicar2093@gmail.com",
		Birthday: time.Date(1993, time.August, 20, 0, 0, 0, 0, time.UTC),
		RoleID:   3,
	}

	err := userService.CreateUser(userRequuest)

	if err != nil {
		t.Fatal(err)
	}

}
