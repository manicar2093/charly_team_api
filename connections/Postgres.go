package connections

import (
	"github.com/manicar2093/charly_team_api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgressConnection() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DBConnectionURL()))
	if err != nil {
		panic(err)
	}

	return db

}
