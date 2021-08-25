package config

import (
	"fmt"
	"os"
)

var (
	DBHost, DBPort, DBName, DBUser, DBPassword, dbUrl, AWSRegion, AWSAccessKeyID, AWSSecretAccessKey string
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

func StartConfig() {
	DBHost = GetEnvOrPanic("DB_HOST")
	DBPort = GetEnvOrPanic("DB_PORT")
	DBName = GetEnvOrPanic("DB_NAME")
	DBUser = GetEnvOrPanic("DB_USER")
	DBPassword = GetEnvOrPanic("DB_PASSWORD")
	AWSRegion = GetEnvOrPanic("AWS_REGION")
	AWSAccessKeyID = GetEnvOrPanic("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey = GetEnvOrPanic("AWS_SECRET_ACCESS_KEY")
	dbUrl = os.Getenv("DB_URL")
}
