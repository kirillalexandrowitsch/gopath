#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
MIGRATIONS_DIR="$ROOT_DIR/backend/migrations"

if [ ! -d "$MIGRATIONS_DIR" ]; then
	echo "migrations directory not found: $MIGRATIONS_DIR" >&2
	exit 1
fi

set -- "$MIGRATIONS_DIR"/*.sql
if [ ! -e "$1" ]; then
	echo "migration files not found: $MIGRATIONS_DIR/*.sql" >&2
	exit 1
fi

seen_prefixes=""
count=0

for MIGRATION_FILE do
	MIGRATION_NAME=${MIGRATION_FILE##*/}

	case "$MIGRATION_NAME" in
		[0-9][0-9][0-9][0-9][0-9][0-9]_*.sql) ;;
		*)
			echo "invalid migration filename: $MIGRATION_NAME" >&2
			exit 1
			;;
	esac

	MIGRATION_SLUG=${MIGRATION_NAME#??????_}
	MIGRATION_SLUG=${MIGRATION_SLUG%.sql}
	if [ -z "$MIGRATION_SLUG" ]; then
		echo "invalid migration filename: $MIGRATION_NAME" >&2
		exit 1
	fi

	MIGRATION_PREFIX=${MIGRATION_NAME%%_*}
	case " $seen_prefixes" in
		*" $MIGRATION_PREFIX "*)
			echo "duplicate migration prefix: $MIGRATION_PREFIX" >&2
			exit 1
			;;
	esac
	seen_prefixes="${seen_prefixes}${MIGRATION_PREFIX} "

	if [ ! -s "$MIGRATION_FILE" ]; then
		echo "empty migration file: $MIGRATION_NAME" >&2
		exit 1
	fi

	FIRST_LINE=$(awk 'NF { gsub(/^[[:space:]]+|[[:space:]]+$/, ""); print; exit }' "$MIGRATION_FILE")
	LAST_LINE=$(awk 'NF { line = $0 } END { gsub(/^[[:space:]]+|[[:space:]]+$/, "", line); print line }' "$MIGRATION_FILE")

	if [ "$FIRST_LINE" != "BEGIN;" ]; then
		echo "migration must start with BEGIN;: $MIGRATION_NAME" >&2
		exit 1
	fi

	if [ "$LAST_LINE" != "COMMIT;" ]; then
		echo "migration must end with COMMIT;: $MIGRATION_NAME" >&2
		exit 1
	fi

	count=$((count + 1))
	printf 'ok - %s\n' "$MIGRATION_NAME"
done

printf 'ok - validated %s migration file(s)\n' "$count"
