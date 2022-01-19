#!/usr/bin/env bash

set -eux
go mod tidy
go mod download
go mod vendor