SHELL=/bin/bash

.PHONY: help
help:
	@echo Alchemy Furnace Web
	@echo "准备工作： make install"
	@echo "本地构建： make build"
	@echo "本地启动： make serve"

.PHONY: install
install:
	yarn install

.PHONY: serve
serve:
	yarn serve

.PHONY: load-local-config
load-local-config:
	@cp -v $(CURDIR)/.data/proxy.local.conf.json $(CURDIR)/proxy.conf.json

.PHONY: load-dev-config
load-release-config:
	@cp -v $(CURDIR)/.data/proxy.release.conf.json $(CURDIR)/proxy.conf.json

.PHONY: docker
docker-dist:
	docker buildx build \
		--platform linux/amd64 \
		--pull \
		--push \
		-f Dockerfile.dist \
		-t hub.deepin.com/ljiang/alchemy-furnace-web:latest .
