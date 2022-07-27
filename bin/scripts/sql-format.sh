#!/bin/bash

set -e

REPO_NAME="$(basename "$(git rev-parse --show-toplevel)")"
if [[ $(basename "$PWD") != "$REPO_NAME" ]]; then
  echo "Please run this script from the root repo's directory: '$REPO_NAME'"
  echo "Current directory: $PWD"
  exit 1
fi

SQL_DIRS="internal/services/queries db/migrations"
for slq_dir in $SQL_DIRS; do
  pg_format \
    --spaces 2 \
    --wrap-limit 88 \
    --function-case 2 \
    --keyword-case 1 \
    --placeholder "sqlc\\.(arg|narg)\\(:?[^)]*\\)" \
    --inplace \
    $(find "$slq_dir" -maxdepth 1 -name '*.sql' | tr '\n' ' ')
done
