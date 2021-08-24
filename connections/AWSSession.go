package connections

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/manicar2093/charly_team_api/config"
)

var awsSession *session.Session

// GetAWSSession creates a AWS Session to be use.
// If an error occured during connection it panics
func GetAWSSession() *session.Session {

	if awsSession == nil {

		awsSession = session.Must(
			session.NewSession(
				&aws.Config{
					Region: aws.String(config.AWSRegion),
				},
			),
		)
	}

	return awsSession

}
