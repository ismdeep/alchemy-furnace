#!/usr/bin/env bash

set -eux
bash ./.shell/test-init.bash
go test ./handler/... -count=1