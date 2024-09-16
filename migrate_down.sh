#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Construct the database URL
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL}"

# Run the migration
migrate -database "${DB_URL}" -path db/migrations down
