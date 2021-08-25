package testfunc

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/manicar2093/charly_team_api/config"
)

// LoadEnvFileOrPanic loads the requested env file and start config. If file does not exist
// it consider load variables from system.
//
// If other type of error occurs it will panic.
func LoadEnvFileOrPanic(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Printf("Loading from sys env instead %s file", envFilePath)
			return
		}
		panic(err)
	}

	config.StartConfig()
}
