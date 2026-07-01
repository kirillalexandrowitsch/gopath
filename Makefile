GOCACHE ?= /private/tmp/gopath-gocache
COMPOSE_ENV ?= infra/.env.example
COMPOSE_FILE ?= infra/compose.yml

.PHONY: help backend-test frontend-lint frontend-build compose-config diff-check

help:
	@printf '%s\n' \
		'GoPath developer commands:' \
		'  make backend-test    Run backend Go tests' \
		'  make frontend-lint   Run frontend lint' \
		'  make frontend-build  Build frontend' \
		'  make compose-config  Validate Docker Compose config' \
		'  make diff-check      Check git diff whitespace'

backend-test:
	cd backend && GOCACHE=$(GOCACHE) go test ./...

frontend-lint:
	cd frontend && npm run lint

frontend-build:
	cd frontend && npm run build

compose-config:
	docker compose --env-file $(COMPOSE_ENV) -f $(COMPOSE_FILE) config

diff-check:
	git diff --check
