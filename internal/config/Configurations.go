package config

import (
	"fmt"
	"os"
)

var (
	DBHost, DBPort, DBName, DBUser, DBPassword, dbUrl, AWSRegion, AWSAccessKeyID, AWSSecretAccessKey, CognitoPoolID string
)

const (
	// PassLen represent the number of character will compose a password
	PassLen = 8
	// PassNumDigits indicates how many digits a password will contain
	PassNumDigits = 2
	// PassNumSymbols indicates how many symbols a password will contain
	PassNumSymbols = 1
	PageSize       = 10
	AvatarURLSrc   = "https://avatars.dicebear.com/api/jdenticon/"
)

func DBConnectionURL() string {
	if dbUrl == "" {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	}

	return dbUrl
}

func StartConfig() {
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	AWSRegion = os.Getenv("AWS_REGION")
	AWSAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	CognitoPoolID = os.Getenv("COGNITO_POOL_ID")
	dbUrl = os.Getenv("DB_URL")
}
