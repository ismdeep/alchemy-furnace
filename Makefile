SHELL=/bin/bash

.PHONY: create
create: clean

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	rm -f data.db

.PHONY: test
test:
	go test ./handler/... -count=1

.PHONY: swag-doc
swag-doc:
	swag init

.PHONY: docker-local
docker-local:
	docker build -t ismdeep/alchemy-furnace:local .

.PHONY: docker-build
docker-build:
	docker build -t ismdeep/alchemy-furnace:latest .

.PHONY: vendor
vendor:
	go mod tidy
	go mod download
	go mod vendor