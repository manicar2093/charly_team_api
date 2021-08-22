package testfunc

import (
	"github.com/joho/godotenv"
	"github.com/manicar2093/charly_team_api/config"
)

// LoadEnvFileOrPanic loads the requested env file and start config
// to access to env variables
func LoadEnvFileOrPanic(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		panic(err)
	}

	config.StartConfig()
}
