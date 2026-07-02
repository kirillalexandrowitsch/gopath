#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
COMPOSE_FILE="$ROOT_DIR/infra/compose.yml"
LOCAL_ENV_FILE="$ROOT_DIR/infra/.env"
EXAMPLE_ENV_FILE="$ROOT_DIR/infra/.env.example"
MIGRATIONS_DIR="$ROOT_DIR/backend/migrations"

psql_exec() {
	docker compose --env-file "$ENV_FILE" -f "$COMPOSE_FILE" exec -T postgres \
		psql -v ON_ERROR_STOP=1 -U "$GOPATH_POSTGRES_USER" -d "$GOPATH_POSTGRES_DB" "$@"
}

checksum_file() {
	if command -v sha256sum >/dev/null 2>&1; then
		sha256sum "$1" | awk '{ print $1 }'
	elif command -v shasum >/dev/null 2>&1; then
		shasum -a 256 "$1" | awk '{ print $1 }'
	else
		echo "sha256sum or shasum is required to calculate migration checksums" >&2
		exit 1
	fi
}

sql_escape() {
	printf "%s" "$1" | sed "s/'/''/g"
}

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

printf 'ensuring migration tracking table\n'
psql_exec -c "CREATE TABLE IF NOT EXISTS schema_migrations (version TEXT PRIMARY KEY, filename TEXT NOT NULL UNIQUE, checksum TEXT NOT NULL, applied_at TIMESTAMPTZ NOT NULL DEFAULT now());"

for MIGRATION_FILE do
	MIGRATION_NAME=${MIGRATION_FILE##*/}
	MIGRATION_VERSION=${MIGRATION_NAME%%_*}
	MIGRATION_CHECKSUM=$(checksum_file "$MIGRATION_FILE")

	MIGRATION_VERSION_SQL=$(sql_escape "$MIGRATION_VERSION")
	MIGRATION_NAME_SQL=$(sql_escape "$MIGRATION_NAME")
	MIGRATION_CHECKSUM_SQL=$(sql_escape "$MIGRATION_CHECKSUM")

	APPLIED_CHECKSUM=$(psql_exec -q -t -A -c "SELECT checksum FROM schema_migrations WHERE version = '$MIGRATION_VERSION_SQL';")
	if [ -n "$APPLIED_CHECKSUM" ]; then
		if [ "$APPLIED_CHECKSUM" != "$MIGRATION_CHECKSUM" ]; then
			echo "migration checksum mismatch: $MIGRATION_NAME" >&2
			echo "database checksum: $APPLIED_CHECKSUM" >&2
			echo "file checksum: $MIGRATION_CHECKSUM" >&2
			exit 1
		fi

		printf 'skipping migration: %s\n' "$MIGRATION_NAME"
		continue
	fi

	printf 'applying migration: %s\n' "$MIGRATION_NAME"
	psql_exec <"$MIGRATION_FILE"

	psql_exec -c "INSERT INTO schema_migrations (version, filename, checksum) VALUES ('$MIGRATION_VERSION_SQL', '$MIGRATION_NAME_SQL', '$MIGRATION_CHECKSUM_SQL');"
done
