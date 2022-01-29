package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestimagessaver"
	"github.com/manicar2093/charly_team_api/internal/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(SaveBiotestImagesAWSLambdaTests))
}

type SaveBiotestImagesAWSLambdaTests struct {
	suite.Suite
	ctx                                                       context.Context
	biotestImagesSaver                                        *biotestimagessaver.MockBiotestImagesSaver
	saveBiotestImagesAWSLambda                                *SaveBiotestImagesAWSLambda
	request                                                   biotestimagessaver.BiotestImagesSaverRequest
	biotestUUID, frontImage, backImage, leftImage, rightImage string
	ordinaryError                                             error
}

func (c *SaveBiotestImagesAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestImagesSaver = &biotestimagessaver.MockBiotestImagesSaver{}
	c.saveBiotestImagesAWSLambda = NewSaveBiotestImagesAWSLambda(c.biotestImagesSaver)
	c.biotestUUID = "biotest_uuid"
	c.frontImage = "front/image/path"
	c.backImage = "back/image/path"
	c.leftImage = "left/image/path"
	c.rightImage = "right/image/path"
	c.request = biotestimagessaver.BiotestImagesSaverRequest{
		BiotestUUID:      c.biotestUUID,
		FrontPicture:     c.frontImage,
		BackPicture:      c.backImage,
		LeftSidePicture:  c.leftImage,
		RightSidePicture: c.rightImage,
	}
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *SaveBiotestImagesAWSLambdaTests) TearDownTest() {
	c.biotestImagesSaver.AssertExpectations(c.T())
}

func (c *SaveBiotestImagesAWSLambdaTests) TestHandler() {
	response := biotestimagessaver.BiotestImagesSaverResponse{
		BiotestImagesSaved: &c.request,
	}
	c.biotestImagesSaver.On("Run", c.ctx, &c.request).Return(
		&response,
		nil,
	)

	res, err := c.saveBiotestImagesAWSLambda.Handler(c.ctx, c.request)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusOK, res.StatusCode, "status code not correct")
	c.Nil(res.Body, "body is not correct type")
}

func (c *SaveBiotestImagesAWSLambdaTests) TestHandler_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.biotestImagesSaver.On("Run", c.ctx, &c.request).Return(
		nil,
		validationErrors,
	)

	res, err := c.saveBiotestImagesAWSLambda.Handler(c.ctx, c.request)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *SaveBiotestImagesAWSLambdaTests) TestHandler_UnhandledError() {
	c.biotestImagesSaver.On("Run", c.ctx, &c.request).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.saveBiotestImagesAWSLambda.Handler(c.ctx, c.request)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
