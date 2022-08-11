#!/bin/bash
# shellcheck disable=SC1091,SC2155

set -Eeo pipefail

if [ -t 1 ]; then
  RED="$(tput setaf 1)"
  GREEN="$(tput setaf 2)"
  YELLOW="$(tput setaf 3)"
  BLUE="$(tput setaf 4)"
  MAGENTA="$(tput setaf 5)"
  CYAN="$(tput setaf 6)"
  WHITE="$(tput setaf 7)"
  BOLD="$(tput bold)"
  OFF="$(tput sgr0)"
else
  RED=""
  GREEN=""
  YELLOW=""
  BLUE=""
  MAGENTA=""
  CYAN=""
  WHITE=""
  BOLD=""
  OFF=""
fi

ensure_pwd_is_top_level() {
  TOP_LEVEL="$(git rev-parse --show-toplevel)"

  if [[ -z $TOP_LEVEL ]]; then
    echo "No .git directory found, skipping top level directory check."
    return
  fi

  if [[ "$PWD" != "$TOP_LEVEL" ]]; then
    echo >&2 "
Please run this script from the top level of the repository.
Top level: $TOP_LEVEL
Current directory: $PWD"
    exit
  fi
}

# Retrieve environment variable `var` from `env_file`
get_envvar() {
  local env_file="$1"
  local var="$2"

  if [[ -f "$env_file" ]]; then
    value=$(
      grep -oP "(?<=^$var=)[^ ]+" "$env_file" | head -n 1
    )
    if [[ -z "$value" ]]; then
      err "Variable $var not found in $env_file"
    fi
    echo "$value"
  else
    err "$env_file does not exist"
  fi
}

# Drop and recreate database `db`
drop_db() {
  local db="$1"
  local project_prefix="$(get_envvar .env PROJECT_PREFIX)"
  local app_env="$(get_envvar .env APP_ENV)"
  local postgres_user="$(get_envvar .env POSTGRES_USER)"
  local postgres_db="$(get_envvar .env POSTGRES_DB)"

  docker exec -t postgres_db_"$project_prefix"_"$app_env" psql --no-psqlrc -U "$postgres_user" -d "$postgres_db" -c "CREATE DATABASE test OWNER $postgres_user;" || echo "Database test already exists. Skipping"
  docker exec -t postgres_db_"$project_prefix"_"$app_env" dropdb --if-exists -f "$db"
  docker exec -t postgres_db_"$project_prefix"_"$app_env" psql --no-psqlrc -U "$postgres_user" -d test -c "CREATE DATABASE $db;"
}
