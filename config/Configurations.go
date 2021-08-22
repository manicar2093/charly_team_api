package config

import (
	"fmt"
	"os"
)

var (
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
)

func GetEnvOrPanic(envName string) string {

	value, exists := os.LookupEnv(envName)

	if !exists {
		panic(fmt.Sprintf("env varialbe '%s' is not set and is a must", envName))
	}

	if value == "" {
		panic(fmt.Sprintf("env variable '%s' is empty and is a must", envName))
	}

	return value
}

func init() {
	DBHost = GetEnvOrPanic("DB_HOST")
	DBPort = GetEnvOrPanic("DB_PORT")
	DBName = GetEnvOrPanic("DB_NAME")
	DBUser = GetEnvOrPanic("DB_USER")
	DBPassword = GetEnvOrPanic("DB_PASSWORD")
}
