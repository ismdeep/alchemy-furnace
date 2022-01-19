#!/usr/bin/env bash

set -eux
docker stop alchemy-furnace-db || true
docker rm   alchemy-furnace-db || true
