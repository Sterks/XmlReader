package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ConnectionString ...
type ConnectionString struct {
	ConnectionString string
}

type pgDb struct {
	dbConn *sqlx.DB
}

// InitDb ...
func InitDb(cfg ConnectionString) (*pgDb, error) {
	dbConn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	p := &pgDb{dbConn: dbConn}
	return p, nil
}
