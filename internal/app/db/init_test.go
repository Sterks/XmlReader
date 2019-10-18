package db

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func TestPgDb_CheckDB(t *testing.T) {
	type fields struct {
		logger logrus.Logger
		dbConn *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDb{
				logger: tt.fields.logger,
				dbConn: tt.fields.dbConn,
			}
			if err := p.CheckDB(); (err != nil) != tt.wantErr {
				t.Errorf("PgDb.CheckDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
