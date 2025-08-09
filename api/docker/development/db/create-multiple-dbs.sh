#!/bin/bash
set -e
for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
  psql -U "$POSTGRES_USER" -c "CREATE DATABASE $db"
done