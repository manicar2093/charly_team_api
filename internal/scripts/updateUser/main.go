package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	session := session.Must(
		session.NewSession(
			&aws.Config{
				Region: aws.String("us-east-1"),
			},
		),
	)

	provider := cognitoidentityprovider.New(session)

	username := "micqimplementations@gmail.com"
	userPoolID := "us-east-1_Y1mSTZpCI"
	name := "email_verified"
	value := "true"

	out, err := provider.AdminUpdateUserAttributes(&cognitoidentityprovider.AdminUpdateUserAttributesInput{
		Username:   &username,
		UserPoolId: &userPoolID,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: &name, Value: &value},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out)

}
