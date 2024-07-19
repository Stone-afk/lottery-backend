#!/usr/bin/env bash

set -e
docker compose -f docker-compose.env.yaml down
docker compose -f docker-compose.env.yaml up -d
go test -race ./... -tags=e2e
docker compose -f docker-compose.env.yaml down
