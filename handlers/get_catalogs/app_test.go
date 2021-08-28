package main

import (
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/stretchr/testify/suite"
)

type AppTests struct {
	suite.Suite
	db mocks.Findable
}
