package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/Sterks/XmlReader/internal/app/readerxml"
	"github.com/sirupsen/logrus"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "/Users/drunov/GoProject/XmlReader/configs/reader.toml", "path config file")
}

func main() {
	flag.Parse()
	config := readerxml.NewConfig()
	_, err := toml.DecodeFile(ConfigPath, &config)
	if err != nil {
		logrus.Error(err)
	}
	s := readerxml.New(config)
	s.Start()
}
