#!/bin/bash
# shellcheck disable=SC1091,SC2155,SC2086

# set -Eeo pipefail

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

# Prompt the user for confirmation.
confirm() {
  test -n "$NO_CONFIRMATION" && return

  local prompt="$1"
  local response

  [[ -z $prompt ]] && prompt="Are you sure?"

  prompt+=" [y/n]"

  while true; do
    read -r -p "$prompt " response
    case "${response,,}" in
    [y][e][s] | [y])
      return 0
      ;;
    [n][o] | [n])
      return 1
      ;;
    *) ;;
    esac
  done
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

join_by() {
  local d=${1-} f=${2-}
  if shift 2; then
    printf %s "$f" "${@/#/$d}"
  fi
}

to_pascal() {
  local s=${1^}
  local re='(.*_-+)([[:lower:]].*)'
  while [[ $s =~ $re ]]; do
    s=${BASH_REMATCH[1]}${BASH_REMATCH[2]^}
  done
  s=${s//[^[:alnum:]]/}
  printf '%s\n' "$s"
}

to_lower() {
  local s="$1"
  local re='([[:upper:]])'
  while [[ $s =~ $re ]]; do
    s="${s/${BASH_REMATCH[0]}/${BASH_REMATCH[0],}}"
  done
  printf '%s\n' "$s"
}

element_in_array() {
  local element=$1
  shift
  local arr=("$@")
  for item in "${arr[@]}"; do
    if [[ "$item" == "$element" ]]; then
      return 0 # element found
    fi
  done
  return 1 # element not found
}

restart_pid() {
  # get command + args
  SAVED_COMMAND="$(while IFS= read -r -d $'\0' f; do printf '%q ' "$f"; done </proc/$1/cmdline)"
  # original working directory for the command
  cd /proc/$1/cwd
  kill $1
  eval $SAVED_COMMAND &
  disown # send to background as before
}

err() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
  sleep 0.1 # while processing xerr in background
  # kill -s SIGUSR1 $PROC
  # FIXME parallel (sub-)subshell management instead of force killing
  kill 0
  exit 1 # if not using trap
}

######################## env vars ###########################

# Retrieve all environment variables from `env_file` and
# set the key-value pairs in the given associative array
get_envvars() {
  declare -n arr="$1" # pass ref by name
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

# Check all environment variables in a template are present in another.
ensure_envvars_set() {
  local env_template="$1"
  local env_file="$2"
  local -i n_missing

  test -f "$env_template" || err "File $env_template does not exist"
  test -f "$env_file" || err "File $env_file does not exist"

  while IFS= read -r envvar; do
    var=${envvar%%=}
    if [[ "$var" =~ ^\#.* ]]; then
      continue
    fi
    if ! grep -qoE "^${var}[ ]?=" "$env_file"; then
      echo "$env_file does not contain the variable $var (required by $env_template)"
      ((n_missing++))
    fi
  done <"$env_template"

  { ((n_missing != 0)) && exit 1; } || true
}

######################## db ###########################

# Drop and recreate database `db`. Defaults to POSTGRES_DB.
drop_and_recreate_db() {
  local db="${1:POSTGRES_DB}"

  _pg_isready

  dockerdb psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "CREATE DATABASE test OWNER $POSTGRES_USER;" 2>/dev/null || true

  echo "${RED}${BOLD}Dropping database $db.${OFF}"
  dockerdb \
    dropdb --if-exists -f "$db"

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  dockerdb psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d test \
    -c "CREATE DATABASE $db OWNER $POSTGRES_USER;"
}

dockerdb() {
  docker exec -i postgres_db_"$PROJECT_PREFIX" "$@"
}

dockerdb_psql() {
  docker exec -i postgres_db_"$PROJECT_PREFIX" psql -qtAX -v ON_ERROR_STOP=on "$@"
}

# Create database `db`.
create_db_if_not_exists() {
  local db="$1"

  _pg_isready

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  {
    dockerdb psql --no-psqlrc -U "$POSTGRES_USER" \
      -tc "SELECT 1 FROM pg_database WHERE datname = '$db'" |
      grep -q 1
  } ||
    dockerdb psql --no-psqlrc -U "$POSTGRES_USER" -c "CREATE DATABASE $db" ||
    echo "Skipping $db database creation"
}

# Stop running processes in `db`.
stop_db_processes() {
  local db="$1"

  _pg_isready

  echo "${BLUE}${BOLD}Stopping any running processes for database $db.${OFF}"
  dockerdb psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "select pg_terminate_backend(pid) \
        from pg_stat_activity \
        where datname='$db'" >/dev/null
}

_pg_isready() {
  pg_ready=0
  while [[ ! $pg_ready -eq 1 ]]; do
    dockerdb \
      pg_isready -U "$POSTGRES_USER" || {
      echo "${YELLOW}Waiting for postgres database to be ready...${OFF}"
      sleep 2
      continue
    }
    pg_ready=1
  done
}

######################## docker ###########################

gzip_latest_image() {
  local dir="$1"
  local image="$2"
  echo "Saving latest image $image to $dir"
  mkdir -p "$dir"
  docker save "$image":latest | gzip >"$dir/${image}_latest.tar.gz"
}

load_latest_gzip_image() {
  local dir="$1"
  local image="$2"
  echo "Loading latest image $image from $dir"
  docker load <"$dir/${image}_latest.tar.gz"
}
