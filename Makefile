SHELL=/bin/bash

.PHONY: create
create: clean
	docker run --name alchemy-furnace-etcd \
		--env ALLOW_NONE_AUTHENTICATION=yes \
		--env ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
		-p 2379:2379 \
		-d hub.deepin.com/library/bitnami/etcd:latest
	docker run --name alchemy-furnace-db \
		-e MYSQL_ROOT_PASSWORD=liandanlu123456 \
		-e MYSQL_DATABASE=alchemy_furnace \
		-p 10006:3306 \
		-d hub.deepin.com/library/mysql:8.0
	go test ./test/... -count=1

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	-docker stop alchemy-furnace-db
	-docker rm   alchemy-furnace-db
	-docker stop alchemy-furnace-etcd
	-docker rm   alchemy-furnace-etcd

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