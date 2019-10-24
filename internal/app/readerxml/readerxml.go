package readerxml

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sterks/XmlReader/internal/app/db"
	"github.com/sirupsen/logrus"
)

// ReaderXML ...
type ReaderXML struct {
	config *Config
	logger *logrus.Logger
	db     *db.PgDb
}

// New создание нового экземпляра
func New(config *Config) *ReaderXML {
	return &ReaderXML{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *ReaderXML) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting server")
	// if err := s.db.CheckDB(); err != nil {
	// 	s.logger.Error("Нет подключения к базе данных!")
	// }
	// s.waitForSignal()
	return nil
}

func (s *ReaderXML) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ReaderXML) waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	g := <-ch
	s.logger.Info("Сигнал на отмену/", g, "/exiting")
}
