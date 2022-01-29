package aws

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CongitoClient interface {
	AdminCreateUser(input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
}

func NewCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	session := GetAWSSession()
	return cognitoidentityprovider.New(session)
}
