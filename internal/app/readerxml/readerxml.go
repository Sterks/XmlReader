package readerxml

import (
	"github.com/sirupsen/logrus"
)

// ReaderXML ...
type ReaderXML struct {
	config *Config
	logger *logrus.Logger
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
}
