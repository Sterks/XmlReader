package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/Sterks/XmlReader/cmd/readerxml/router"
	"github.com/Sterks/XmlReader/internal/app/configuration"
	ftpdownloader "github.com/Sterks/XmlReader/internal/app/ftpDownloader"
	"github.com/Sterks/XmlReader/internal/common"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"github.com/sirupsen/logrus"
)

var (
	// ConfigPath ... Нужно исправить
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "/Users/drunov/GoProject/XmlReader/configs/reader.toml", "path config file")
}

func main() {
	flag.Parse()
	router.StartServer()
	// StartServices()
}

// StartServices ...
func StartServices() {
	// Получение файлов
	con := configuration.NewConfig()
	f := ftpdownloader.New(con)
	_, err := toml.DecodeFile(ConfigPath, &con)
	if err != nil {
		logrus.Error(err)
	}
	if err := f.Start(); err != nil {
		logrus.Errorf("Не стартонул %v", err)
	}
	ftp, err := f.Connect()
	if err != nil {
		logrus.Error("Нет подключения к FTP!")
	}
	listFiles := f.GetFiles(ftp, common.FromDate(), common.ToDate())
	for _, value := range listFiles {
		// fmt.Println(value)
		f.AdderRezultDbAndFs(&value)
		f.SaveResultToDisk()
	}
}

// _, err2 := toml.DecodeFile(ConfigPath, &con)
// if err2 != nil {
// 	logrus.Error(err2)
// }
// config := readerxml.NewConfig()
// _, err := toml.DecodeFile(ConfigPath, &config)
// if err != nil {
// 	logrus.Error(err)
// }
// s := readerxml.New(config)
// s.Start()
