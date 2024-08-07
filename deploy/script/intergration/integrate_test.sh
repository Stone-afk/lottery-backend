#!/usr/bin/env bash

set -e
docker-compose -f docker-compose-env.yml down
docker-compose -f docker-compose-env.yml up -d
go test -race ./... -tags=e2e
docker-compose -f docker-compose-env.yml down
