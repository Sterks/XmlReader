package ftpdownloader

import (
	"log"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	"github.com/secsy/goftp"

	"github.com/sirupsen/logrus"
)

// FtpDownloader ...
type FtpDownloader struct {
	config *configuration.Configuration
	logger *logrus.Logger
	ftp    *goftp.Client
}

//New ...
func New(con *configuration.Configuration) *FtpDownloader {
	return &FtpDownloader{
		config: con,
		logger: logrus.New(),
	}
}

func (f *FtpDownloader) configureLogger() {
	level, err := logrus.ParseLevel(f.config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	f.logger.SetLevel(level)
}

//Connect ...
func (f *FtpDownloader) Connect() (*goftp.Client, error) {
	con := goftp.Config{
		User:     "free",
		Password: "free",
	}
	ftp, err := goftp.DialConfig(con, f.config.Ftp_connect)
	if err != nil {
		return nil, err
	}
	return ftp, nil
}
