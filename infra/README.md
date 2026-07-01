# Local Infrastructure

This directory contains the local PostgreSQL and Redis stack for GoPath.

## Environment

Copy the example file before changing local credentials or ports:

```sh
cp infra/.env.example infra/.env
```

`infra/.env` is ignored by git. Keep tracked defaults in `infra/.env.example`.

## Start

```sh
docker compose --env-file infra/.env -f infra/compose.yml up -d
```

If you did not create `infra/.env`, use the example file directly:

```sh
docker compose --env-file infra/.env.example -f infra/compose.yml up -d
```

## Apply Migrations

Apply the current schema after PostgreSQL is healthy:

```sh
./scripts/migrate-up.sh
```

The script uses `infra/.env` when it exists and falls back to `infra/.env.example`.

## Stop

Stop the local stack without deleting volumes:

```sh
docker compose --env-file infra/.env -f infra/compose.yml down
```

Do not pass `--volumes` unless you explicitly want to delete local database data.
