package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	ftpdownloader "github.com/Sterks/XmlReader/internal/app/ftpDownloader"
	"github.com/Sterks/XmlReader/internal/app/readerxml"
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
	con := configuration.NewConfig()
	_, err2 := toml.DecodeFile(ConfigPath, &con)
	if err2 != nil {
		logrus.Error(err2)
	}
	p := ftpdownloader.New(con)
	ftp, err := p.Connect()
	if err != nil {
		logrus.Error("Нет подключения к FTP!")
	}
	fmt.Println(ftp)

	config := readerxml.NewConfig()
	_, err = toml.DecodeFile(ConfigPath, &config)
	if err != nil {
		logrus.Error(err)
	}
	s := readerxml.New(config)
	s.Start()

}
