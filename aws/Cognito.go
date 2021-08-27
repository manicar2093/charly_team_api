package aws

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/manicar2093/charly_team_api/connections"
)

type CongitoClient interface {
	AdminCreateUser(input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
}

func NewCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	session := connections.GetAWSSession()
	return cognitoidentityprovider.New(session)
}
