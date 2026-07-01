#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
COMPOSE_FILE="$ROOT_DIR/infra/compose.yml"
LOCAL_ENV_FILE="$ROOT_DIR/infra/.env"
EXAMPLE_ENV_FILE="$ROOT_DIR/infra/.env.example"
MIGRATIONS_DIR="$ROOT_DIR/backend/migrations"

if [ -f "$LOCAL_ENV_FILE" ]; then
	ENV_FILE="$LOCAL_ENV_FILE"
else
	ENV_FILE="$EXAMPLE_ENV_FILE"
fi

if [ ! -d "$MIGRATIONS_DIR" ]; then
	echo "migrations directory not found: $MIGRATIONS_DIR" >&2
	exit 1
fi

set -a
. "$ENV_FILE"
set +a

GOPATH_POSTGRES_DB=${GOPATH_POSTGRES_DB:-gopath}
GOPATH_POSTGRES_USER=${GOPATH_POSTGRES_USER:-gopath}

LC_ALL=C
export LC_ALL

set -- "$MIGRATIONS_DIR"/*.sql
if [ ! -e "$1" ]; then
	echo "migration files not found: $MIGRATIONS_DIR/*.sql" >&2
	exit 1
fi

for MIGRATION_FILE do
	MIGRATION_NAME=${MIGRATION_FILE##*/}
	printf 'applying migration: %s\n' "$MIGRATION_NAME"

	docker compose --env-file "$ENV_FILE" -f "$COMPOSE_FILE" exec -T postgres \
		psql -v ON_ERROR_STOP=1 -U "$GOPATH_POSTGRES_USER" -d "$GOPATH_POSTGRES_DB" \
		<"$MIGRATION_FILE"
done
