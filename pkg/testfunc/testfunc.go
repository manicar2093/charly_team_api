package testfunc

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/manicar2093/health_records/internal/config"
)

// LoadEnvFileOrPanic loads the requested env file and start config. If file does not exist
// it consider load variables from system.
//
// If other type of error occurs it will panic.
func LoadEnvFileOrPanic(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			panic(err)
		}
		log.Printf("Loading from sys env instead %s file", envFilePath)
	}

	config.StartConfig()
}

func PrintJsonIndented(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent(" ", " ")
	encoder.Encode(data)
}
