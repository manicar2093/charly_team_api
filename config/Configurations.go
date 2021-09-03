package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	DBHost, DBPort, DBName, DBUser, DBPassword, dbUrl, AWSRegion, AWSAccessKeyID, AWSSecretAccessKey, CognitoPoolID string
	PageSize                                                                                                        int
)

const (
	// PassLen represent the number of character will compose a password
	PassLen = 8
	// PassNumDigits indicates how many digits a password will contain
	PassNumDigits = 2
	// PassNumSymbols indicates how many symbols a password will contain
	PassNumSymbols = 1
)

func GetEnvOrPanic(envName string) string {

	value, exists := os.LookupEnv(envName)

	if !exists || value == "" {
		panic(fmt.Sprintf("env varialbe '%s' is empty or not set and is a must", envName))
	}

	return value
}

func DBConnectionURL() string {
	if dbUrl == "" {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	}

	return dbUrl
}

func getPageSize() int {
	pageSizeConverted, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))

	if err != nil {
		panic("wrong page size paramenter. should be an int")
	}
	return pageSizeConverted
}

func StartConfig() {
	DBHost = GetEnvOrPanic("DB_HOST")
	DBPort = GetEnvOrPanic("DB_PORT")
	DBName = GetEnvOrPanic("DB_NAME")
	DBUser = GetEnvOrPanic("DB_USER")
	DBPassword = GetEnvOrPanic("DB_PASSWORD")
	AWSRegion = GetEnvOrPanic("AWS_REGION")
	AWSAccessKeyID = GetEnvOrPanic("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey = GetEnvOrPanic("AWS_SECRET_ACCESS_KEY")
	CognitoPoolID = GetEnvOrPanic("COGNITO_POOL_ID")
	dbUrl = os.Getenv("DB_URL")
	PageSize = getPageSize()
}
