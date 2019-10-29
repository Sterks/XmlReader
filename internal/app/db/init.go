package db

import (
	"database/sql"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	_ "github.com/lib/pq" ///
	"github.com/sirupsen/logrus"
)

//PgDb ...
type PgDb struct {
	logger          logrus.Logger
	db              *sql.DB
	config          *configuration.Configuration
	filesRepository *FilesRepository
}

// New ...
func New(config *configuration.Configuration) *PgDb {
	return &PgDb{
		config: config,
	}
}

// Open ...
func (p *PgDb) Open() error {
	db, err := sql.Open("postgres", p.config.ConnectionString)
	if err != nil {
		logrus.Error(err)
	}

	if err := db.Ping(); err != nil {
		return err
	}

	p.db = db
	return nil
}

// Close ...
func (p *PgDb) Close() {
	p.db.Close()
}

// ConfigureLogger ...
func (p *PgDb) ConfigureLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}
	p.logger.SetLevel(level)
	return nil
}

//ConnectionDB ...
func (p *PgDb) ConnectionDB() *sql.DB {
	return p.db
}

//File ...
func (p *PgDb) File() *FilesRepository {
	if p.filesRepository != nil {
		return p.filesRepository
	}

	p.filesRepository = &FilesRepository{
		db: p,
	}

	return p.filesRepository

}

//GetLastFiles ...
func (p *PgDb) GetLastFiles() int {
	number := p.filesRepository.GetIDFile()
	return number
}
