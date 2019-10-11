package main

import (
	"log"
	"mod/internal/app/readerxml"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	configPath = "./configs/reader.toml"
}

func main() {
	config := readerxml.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
}
