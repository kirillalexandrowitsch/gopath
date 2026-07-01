#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
COMPOSE_FILE="$ROOT_DIR/infra/compose.yml"
LOCAL_ENV_FILE="$ROOT_DIR/infra/.env"
EXAMPLE_ENV_FILE="$ROOT_DIR/infra/.env.example"
MIGRATION_FILE="$ROOT_DIR/backend/migrations/000001_init_schema.sql"

if [ -f "$LOCAL_ENV_FILE" ]; then
	ENV_FILE="$LOCAL_ENV_FILE"
else
	ENV_FILE="$EXAMPLE_ENV_FILE"
fi

if [ ! -f "$MIGRATION_FILE" ]; then
	echo "migration file not found: $MIGRATION_FILE" >&2
	exit 1
fi

set -a
. "$ENV_FILE"
set +a

GOPATH_POSTGRES_DB=${GOPATH_POSTGRES_DB:-gopath}
GOPATH_POSTGRES_USER=${GOPATH_POSTGRES_USER:-gopath}

docker compose --env-file "$ENV_FILE" -f "$COMPOSE_FILE" exec -T postgres \
	psql -v ON_ERROR_STOP=1 -U "$GOPATH_POSTGRES_USER" -d "$GOPATH_POSTGRES_DB" \
	<"$MIGRATION_FILE"
