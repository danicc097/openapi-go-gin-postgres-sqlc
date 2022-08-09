#!/bin/bash

source "${BASH_SOURCE%/*}/../.helpers.sh"

set -e

ensure_pwd_is_top_level

SQL_DIRS="internal/postgresql/queries db/migrations"
for slq_dir in $SQL_DIRS; do
  pg_format \
    --spaces 2 \
    --wrap-limit 88 \
    --function-case 2 \
    --keyword-case 1 \
    --placeholder "sqlc\\.(arg|narg)\\(:?[^)]*\\)" \
    --inplace \
    $(find "$slq_dir" -maxdepth 1 -name '*.sql')
done
