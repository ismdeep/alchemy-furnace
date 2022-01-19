#!/usr/bin/env bash

set -eux

docker stop alchemy-furnace-db || true
docker rm   alchemy-furnace-db || true
docker run --name alchemy-furnace-db \
    -e MYSQL_ROOT_PASSWORD=liandanlu123456 \
    -e MYSQL_DATABASE=alchemy_furnace \
    -p 10006:3306 \
    -d hub.deepin.com/library/mysql:8.0
