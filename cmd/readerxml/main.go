package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Sterks/XmlReader/internal/app/readerxml"
)

var (
	configPath string
)

func main() {
	config := readerxml.NewConfig()
	_, err := toml.DecodeFile("/Users/drunov/GoProject/XmlReader/configs/reader.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	s := readerxml.New(config)
	s.Start()
}
