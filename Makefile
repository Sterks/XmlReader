.PHONY: build
build:
	go build -v ./cmd/readerxml

migration:
	migrate -database "postgres://postgres:596run49@127.0.0.1/readerxml_dev?sslmode=disable" -path ./cmd/readerxml/migrations up

.DEFAULT_GOAL := build