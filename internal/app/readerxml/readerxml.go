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
	db     *db.ConnectionString
}

// New создание нового экземпляра
func New(config *Config) *ReaderXML {
	return &ReaderXML{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *ReaderXML) Start() {
	s.logger.Info("Запуск процесса ...")
	s.waitForSignal()
}

func (s *ReaderXML) waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	g := <-ch
	s.logger.Info("Сигнал на отмену/", g, "/exiting")
	// log.("Сигнал на отмену: %v, exiting.", s)
}
