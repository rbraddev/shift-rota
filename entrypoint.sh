#!/bin/sh

echo "Waiting for Postgres..."

while ! nc -z db 5432; do
  sleep 0.1
done

sleep 2

echo "Postgres started"

exec "$@"