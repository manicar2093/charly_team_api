package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost, DBPort, DBName, DBUser, DBPassword, DBURL string
)

func GetEnvOrPanic(envName string) string {

	value, exists := os.LookupEnv(envName)

	if !exists || value == "" {
		panic(fmt.Sprintf("env varialbe '%s' is empty or not set and is a must", envName))
	}

	return value
}

func DBConnectionURL() string {
	if DBURL == "" {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	}

	return DBURL
}

func init() {
	err := godotenv.Load("../.env.example")
	if err != nil {
		panic(err)
	}

	DBHost = GetEnvOrPanic("DB_HOST")
	DBPort = GetEnvOrPanic("DB_PORT")
	DBName = GetEnvOrPanic("DB_NAME")
	DBUser = GetEnvOrPanic("DB_USER")
	DBPassword = GetEnvOrPanic("DB_PASSWORD")
	DBURL = os.Getenv("DB_URL")
}
