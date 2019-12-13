.PHONY: build
build:
	go build -v ./cmd/readerxml

migration:
	migrate -database "postgres://postgres:596run49@127.0.0.1/readerxml_dev?sslmode=disable" -path ./cmd/readerxml/migrations up

migration_drop:
	migrate -database "postgres://postgres:596run49@127.0.0.1/readerxml_dev?sslmode=disable" -path ./cmd/readerxml/migrations drop

migration_create:
	migrate create -ext sql -dir ./cmd/readerxml/migrations add

.DEFAULT_GOAL := build