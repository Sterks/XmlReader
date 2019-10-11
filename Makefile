.PHONY: build
build:
	go build -v ./cmd/readerxml

.DEFAULT_GOAL := build