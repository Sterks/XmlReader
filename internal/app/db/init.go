package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ConnectionString ...
type ConnectionString struct {
	ConnectionString string
}

//PgDb ...
type PgDb struct {
	dbConn *sqlx.DB
}

// InitDb ...
func InitDb(cfg ConnectionString) (*PgDb, error) {
	dbConn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	p := &PgDb{dbConn: dbConn}
	return p, nil
}
