#!/bin/bash
# shellcheck disable=SC1091,SC2155,SC2086

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
  UNDERSCORE="$(tput smul)"
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
  UNDERSCORE=""
  OFF=""
fi

ensure_pwd_is_top_level() {
  TOP_LEVEL="$(git rev-parse --show-toplevel)"

  if [[ -z $TOP_LEVEL ]]; then
    echo "No .git directory found, skipping top level directory check."
    return
  fi

  if [[ "$PWD" != "$TOP_LEVEL" ]] && [[ -z "$IS_TESTING" && -z "$IGNORE_PWD" ]]; then
    echo >&2 "
Please run this script from the top level of the repository.
Top level: $TOP_LEVEL
Current directory: $PWD"
    exit
  fi
}

list_descendants() {
  local desc_pids=$(ps -o pid= --ppid "$1")
  for pid in $desc_pids; do
    list_descendants "$pid"
  done
  echo "$desc_pids"
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

trim_string() {
  : "${1#"${1%%[![:space:]]*}"}"
  : "${_%"${_##*[![:space:]]}"}"
  printf '%s\n' "$_"
}

err() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
  exit 1
}

# Retrieve all environment variables from `env_file` and
# set the key-value pairs in the given associative array
get_envvars() {
  declare -n arr="$1"
  local env_file="$2"
  if [[ -f "$env_file" ]]; then
    while read -r line; do
      if [[ $line =~ ^[\#]?([A-Za-z0-9_]+)[[:space:]]*=[[:space:]]*(.*?)$ ]]; then
        key="$(trim_string ${BASH_REMATCH[1]})"
        val="$(trim_string ${BASH_REMATCH[2]})"
        arr[$key]=$val
      fi
    done <"$env_file"
  else
    err "$env_file does not exist"
  fi
}

# Drop and recreate database `db` using configuration from
# environment variables in `env_file`
drop_and_recreate_db() {
  local db="$1"

  _pg_isready

  docker exec -t postgres_db_"$PROJECT_PREFIX" \
    psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "CREATE DATABASE test OWNER $POSTGRES_USER;" || true

  echo "${RED}${BOLD}Dropping database $db.${OFF}"
  docker exec -t postgres_db_"$PROJECT_PREFIX" \
    dropdb --if-exists -f "$db"

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  docker exec -t postgres_db_"$PROJECT_PREFIX" \
    psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d test \
    -c "CREATE DATABASE $db OWNER $POSTGRES_USER;"
}

create_db_if_not_exists() {
  local db="$1"

  _pg_isready

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  docker exec -t postgres_db_"$PROJECT_PREFIX" \
    psql --no-psqlrc -U "$POSTGRES_USER" -tc "SELECT 1 FROM pg_database WHERE datname = '$db'" | grep -q 1 ||
    docker exec -t postgres_db_"$PROJECT_PREFIX" \
      psql --no-psqlrc -U "$POSTGRES_USER" -c "CREATE DATABASE $db"
}

stop_db_processes() {
  local db="$1"

  _pg_isready

  echo "${BLUE}${BOLD}Stopping any running processes for database $db.${OFF}"
  docker exec -t postgres_db_"$PROJECT_PREFIX" \
    psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "select pg_terminate_backend(pid) \
        from pg_stat_activity \
        where datname='$db'"
}

_pg_isready() {
  pg_ready=0
  while [[ ! $pg_ready -eq 1 ]]; do
    docker exec -t postgres_db_"$PROJECT_PREFIX" \
      pg_isready -U "$POSTGRES_USER" || {
      echo "${YELLOW}Waiting for postgres database to be ready...${OFF}" && sleep 1 && continue
    }
    pg_ready=1
  done
}
