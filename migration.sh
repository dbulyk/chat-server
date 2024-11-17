#!/bin/bash
source .env

export DB_HOST=${DB_HOST:-pg}
export MIGRATION_DSN="host=$DB_HOST port=5432 dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v