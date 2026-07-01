GOCACHE ?= /private/tmp/gopath-gocache
COMPOSE_ENV ?= infra/.env.example
COMPOSE_FILE ?= infra/compose.yml

.PHONY: help backend-fmt-check backend-test backend-verify frontend-lint frontend-build frontend-verify compose-config diff-check

help:
	@printf '%s\n' \
		'GoPath developer commands:' \
		'  make backend-verify  Run backend formatting check and tests' \
		'  make backend-test    Run backend Go tests' \
		'  make frontend-verify Run frontend lint and build checks' \
		'  make frontend-lint   Run frontend lint' \
		'  make frontend-build  Build frontend' \
		'  make compose-config  Validate Docker Compose config' \
		'  make diff-check      Check git diff whitespace'

backend-fmt-check:
	@if [ -n "$$(gofmt -l backend)" ]; then \
		gofmt -l backend; \
		exit 1; \
	fi

backend-test:
	cd backend && GOCACHE=$(GOCACHE) go test ./...

backend-verify: backend-fmt-check backend-test

frontend-lint:
	cd frontend && npm run lint

frontend-build:
	cd frontend && npm run build

frontend-verify: frontend-lint frontend-build

compose-config:
	docker compose --env-file $(COMPOSE_ENV) -f $(COMPOSE_FILE) config

diff-check:
	git diff --check
