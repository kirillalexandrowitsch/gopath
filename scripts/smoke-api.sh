#!/bin/sh
set -eu

API_BASE_URL=${API_BASE_URL:-http://localhost:8080}
API_BASE_URL=${API_BASE_URL%/}

check_get() {
	name=$1
	path=$2
	marker=$3
	body_file=$(mktemp "${TMPDIR:-/tmp}/gopath-smoke-api.XXXXXX")

	if ! status=$(curl -sS -o "$body_file" -w "%{http_code}" "$API_BASE_URL$path"); then
		rm -f "$body_file"
		printf 'not ok - %s request failed\n' "$name" >&2
		exit 1
	fi

	if [ "$status" != "200" ]; then
		printf 'not ok - %s returned HTTP %s\n' "$name" "$status" >&2
		cat "$body_file" >&2
		rm -f "$body_file"
		exit 1
	fi

	if ! grep -Fq "$marker" "$body_file"; then
		printf 'not ok - %s response did not include marker: %s\n' "$name" "$marker" >&2
		cat "$body_file" >&2
		rm -f "$body_file"
		exit 1
	fi

	rm -f "$body_file"
	printf 'ok - %s\n' "$name"
}

check_get "healthz" "/healthz" '"status":"ok"'
check_get "tracks" "/api/v1/tracks" '"go-backend"'
check_get "levels" "/api/v1/levels" '"junior-backend"'
check_get "lesson" "/api/v1/lessons/go-errors-tests" '"go-errors-tests"'
check_get "progress" "/api/v1/progress" '"demo-user"'
check_get "profile" "/api/v1/profile" '"Alex Kim"'
