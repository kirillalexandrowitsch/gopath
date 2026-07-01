GOCACHE ?= /private/tmp/gopath-gocache
COMPOSE_ENV ?= infra/.env.example
COMPOSE_FILE ?= infra/compose.yml
API_BASE_URL ?= http://localhost:8080

.PHONY: help backend-fmt-check backend-test backend-verify frontend-lint frontend-build frontend-verify compose-config stack-config stack-up stack-down smoke-api diff-check

help:
	@printf '%s\n' \
		'GoPath developer commands:' \
		'  make backend-verify  Run backend formatting check and tests' \
		'  make backend-test    Run backend Go tests' \
		'  make frontend-verify Run frontend lint and build checks' \
		'  make frontend-lint   Run frontend lint' \
		'  make frontend-build  Build frontend' \
		'  make compose-config  Validate Docker Compose config' \
		'  make stack-config    Validate local Docker Compose config' \
		'  make stack-up        Start local Docker Compose stack' \
		'  make stack-down      Stop local Docker Compose stack without deleting volumes' \
		'  make smoke-api       Run local API smoke checks' \
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

stack-config: compose-config

stack-up:
	docker compose --env-file $(COMPOSE_ENV) -f $(COMPOSE_FILE) up -d

stack-down:
	docker compose --env-file $(COMPOSE_ENV) -f $(COMPOSE_FILE) down

smoke-api:
	API_BASE_URL=$(API_BASE_URL) sh scripts/smoke-api.sh

diff-check:
	git diff --check
