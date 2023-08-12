package connections

import (
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"
	"github.com/manicar2093/health_records/internal/config"
)

func PostgressConnection() rel.Repository {

	adapter, err := postgres.Open(config.DBConnectionURL())

	if err != nil {
		panic(err)
	}

	return rel.New(adapter)

}
