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
docker:
	docker buildx build \
		--platform linux/amd64 \
		--pull \
		--push \
		-t ismdeep/alchemy-furnace:latest .

.PHONY: vendor
vendor:
	go mod tidy
	go mod download
	go mod vendor