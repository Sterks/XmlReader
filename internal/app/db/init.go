package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//PgDb ...
type PgDb struct {
	logger logrus.Logger
	dbConn *sqlx.DB
}

//CheckDB ...
func (p *PgDb) CheckDB() error {
	// if err := p.configureLogger(); err != nil {
	// 	return err
	// }
	// toml.DecodeFile("/Users/drunov/GoProject/XmlReader/configs/reader.toml", &config)
	db, err := sqlx.Connect("postgres", "postgres://postgres:596run49@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	var k []int
	if err = db.Select(&k, `select "AR_Id" from "ArchFiles" limit 1`); err != nil {
		log.Fatalln(err)
	} else {
		if k[0] > 0 {
			//p.logger.Info("База данных доступна!")
			logrus.Info("База данных доступна!")
		}
	}
	return nil
}

// configureLogger ...
// func (p *PgDb) configureLogger() error {
// 	level, err := logrus.ParseLevel("debug")
// 	if err != nil {
// 		return err
// 	}
// 	p.logger.SetLevel(level)
// 	return nil
// }
