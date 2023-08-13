#!/bin/bash
# shellcheck disable=SC1091,SC2155,SC2086

TOP_LEVEL_DIR="$(git rev-parse --show-toplevel)"

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

# redirects a command/script to tty in order to bypass xlog and xerr pipeline redirection.
with_tty() {
  if test -t 0; then
    "$@" </dev/tty >/dev/tty
  else
    "$@"
  fi
}

ensure_pwd_is_top_level() {
  if [[ -z $TOP_LEVEL_DIR ]]; then
    echo "No .git directory found, skipping top level directory check."
    return
  fi

  cd "$TOP_LEVEL_DIR" || true
}

# Prompt the user for confirmation.
# Most likely will want to always run as `with_tty confirm ...`
confirm() {
  test -n "$CI" && return
  test -n "$NO_CONFIRMATION" && return

  local prompt="$1"
  local response

  [[ -z $prompt ]] && prompt="Are you sure?"

  prompt+=" [y/n]"

  sleep 0.3 # for some reason there's a race between xlog/xerr and this prompt with VSCode terminal
  while true; do
    printf "\n%s " "$prompt"
    read -r response
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

# waits for parallel processes to finish sucessfully, signalling SIGUSR1 otherwise.
wait_without_error() {
  local -i err=0 werr=0
  while
    wait -fn || werr=$?
    ((werr != 127)) # 127: not found
  do
    err=$werr
    ((err == 0)) || break # handle error as soon as it happens
  done
  #trap 'wait || :' EXIT # wait for all jobs before exiting (regardless of handling above)
  if ((err != 0)); then
    echo "A job failed"
    kill -s SIGUSR1 $PROC
    exit 1
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
  local string=$1
  local pascal_case=""

  # Replace spaces with nothing and capitalize the following letter
  string=$(echo "$string" | sed 's/ \([a-z]\)/\U\1/g')

  # Replace underscores and hyphens with spaces
  string=${string//[_-]/ }

  # Split the string into words and capitalize the first letter of each word
  for word in $string; do
    pascal_case+="${word^}"
  done

  echo "$pascal_case"
}

to_lower() {
  local s="$1"
  local re='([[:upper:]])'
  while [[ $s =~ $re ]]; do
    s="${s/${BASH_REMATCH[0]}/${BASH_REMATCH[0],}}"
  done
  printf '%s\n' "$s"
}

# splits a string by the first instance of a separator
function cut_first() {
  local str="$1"
  local separator="$2"
  local first_part="${str%%"$separator"*}"
  local second_part="${str#*"$separator"}"
  echo "$first_part"
  echo "$second_part"
}

# returns 0 if an element has been found
element_in_array() {
  local element=$1
  shift
  local arr=("$@")
  for item in "${arr[@]}"; do
    if [[ "$item" == "$element" ]]; then
      return 0
    fi
  done
  return 1
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
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')] (${YELLOW}${BASH_SOURCE[1]##"$TOP_LEVEL_DIR/"}:${BASH_LINENO[0]}${OFF}): ${RED}$*${OFF}" >&2
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
  local -n __arr="$1" # pass ref by name
  local env_file="$2"
  if [[ -f "$env_file" ]]; then
    while read -r line; do
      if [[ $line =~ ^[\#]?([A-Za-z0-9_]+)[[:space:]]*=[[:space:]]*(.*?)$ ]]; then
        key="$(trim_string ${BASH_REMATCH[1]})"
        val="$(trim_string ${BASH_REMATCH[2]})"
        __arr[$key]=$val
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

# Usage: trap 'show_tracebacks' ERR
show_tracebacks() {
  local err_code="$?"
  set +o xtrace
  local bash_command=${BASH_COMMAND}
  echo "${RED}Error in ${BASH_SOURCE[1]##"$TOP_LEVEL_DIR/"}:${BASH_LINENO[0]} ('$bash_command' exited with status $err_code)${OFF}" >&2

  if [[ $bash_command != xlog* && $bash_command != xerr* && ${#FUNCNAME[@]} -gt 2 ]]; then
    # Print out the stack trace described by $function_stack
    echo "${RED}Traceback of ${BASH_SOURCE[1]} (most recent call last):${OFF}" >&2
    for ((i = 0; i < ${#FUNCNAME[@]} - 1; i++)); do
      local funcname="${FUNCNAME[$i]}"
      [ "$i" -eq "0" ] && funcname=$bash_command
      echo -e "  ${MAGENTA}${BASH_SOURCE[$i + 1]##*\/}:${BASH_LINENO[$i]}${OFF}\\t$funcname" >&2
    done
  fi
  exit 1
}

# Cache given files and return if checksums match or an error code otherwise.
cache_all() {
  if [ $# -lt 2 ]; then
    err "Usage: ${FUNCNAME[0]} <output_cache_md5_path> <file_or_directory> [<file_or_directory> ...]"
  fi

  output_file="$1"
  shift

  if md5sum -c "$output_file" &>/dev/null && [[ $X_FORCE_REGEN -eq 0 ]]; then
    echo "Skipping generation (cached). Force regen with --x-force-regen"
    return 0
  fi

  >"$output_file"

  for arg in "$@"; do
    if [ -d "$arg" ]; then
      find "$arg" -type f -exec md5sum {} + >>"$output_file"
    elif [ -f "$arg" ]; then
      md5sum "$arg" >>"$output_file"
    else
      err "Invalid argument: $arg"
    fi
  done

  return 1
}

# Block build if magic keyword is found in any file
# Args: keyword
search_stopship() {
  { { {
    stopship_keyword="$1"
    local matches
    matches=$(find "$(git rev-parse --show-toplevel)" \
      -type f \
      -not -path "$0" \
      -not -path '**/.git/*' \
      -not -path '**/.venv/*' \
      -not -path '**/node_modules/*' \
      -not -path '**/build/*' \
      -not -path '**/*.pyc' \
      -not -exec git check-ignore -q --no-index {} \; \
      -exec grep --files-with-matches --regexp="$stopship_keyword" {} \;)
    if [[ -n $matches ]]; then
      echo "${RED}'$stopship_keyword'${OFF} found in tracked files."
      echo "Please fix all related issues in the following files:"
      printf "\t %s\n" $matches
      exit 1
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

######################## docker ###########################

docker.redis() {
  docker exec -i redis_"$PROJECT_PREFIX" "$@"
}

docker.postgres() {
  docker exec -i postgres_db_"$PROJECT_PREFIX" "$@"
}

docker.postgres.psql() {
  docker exec -i postgres_db_"$PROJECT_PREFIX" psql -qtAX -v ON_ERROR_STOP=on "$@"
}

# Drop and recreate database `db`. Defaults to POSTGRES_DB.
docker.postgres.drop_and_recreate_db() {
  local db="${1:POSTGRES_DB}"

  docker.postgres.isready

  docker.postgres psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "CREATE DATABASE test OWNER $POSTGRES_USER;" 2>/dev/null || true

  echo "${RED}${BOLD}Dropping database $db.${OFF}"
  docker.postgres \
    dropdb --if-exists -f "$db"

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  docker.postgres psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d test \
    -c "CREATE DATABASE $db OWNER $POSTGRES_USER;"
}

# Create database `db`.
docker.postgres.create_db() {
  local db="$1"

  docker.postgres.isready

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  {
    docker.postgres psql --no-psqlrc -U "$POSTGRES_USER" \
      -tc "SELECT 1 FROM pg_database WHERE datname = '$db'" |
      grep -q 1
  } ||
    docker.postgres psql --no-psqlrc -U "$POSTGRES_USER" -c "CREATE DATABASE $db" ||
    echo "Skipping $db database creation"
}

# Stop running processes in `db`.
docker.postgres.stop_db_processes() {
  local db="$1"

  docker.postgres.isready

  echo "${BLUE}${BOLD}Stopping any running processes for database $db.${OFF}"
  docker.postgres psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "select pg_terminate_backend(pid) \
        from pg_stat_activity \
        where datname='$db'" >/dev/null
}

docker.postgres.isready() {
  pg_ready=0
  while [[ ! $pg_ready -eq 1 ]]; do
    docker.postgres \
      pg_isready -U "$POSTGRES_USER" || {
      echo "${YELLOW}Waiting for postgres database to be ready...${OFF}"
      sleep 2
      continue
    }
    pg_ready=1
  done
}

# Saves latest image to destination.
# Parameters:
#   Output directory
#   Image name
docker.images.save_latest() {
  local dir="$1"
  local image="$2"
  echo "Saving latest image $image to $dir"
  mkdir -p "$dir"
  docker save "$image":latest | gzip >"$dir/${image}_latest.tar.gz"
}

# Loads latest image from destination.
# Parameters:
#   Input directory
#   Image name
docker.images.load_latest() {
  local dir="$1"
  local image="$2"
  echo "Loading latest image $image from $dir"
  docker load <"$dir/${image}_latest.tar.gz"
}

######################## go ###########################

# Stores go structs in package to a given array.
# Parameters:
#    Struct array (nameref)
#    Package directory
go-utils.find_structs() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/[\s]*type\(.*\)struct.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No structs found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

# Stores go interfaces in package to a given array.
# Parameters:
#    Interface array (nameref)
#    Package directory
go-utils.find_interfaces() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/[\s]*type\(.*\)interface.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No interfaces found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

# Stores go enums in package to a given array.
# Parameters:
#    Enum array (nameref)
#    Package directory
go-utils.find_enums() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/.*type[[:space:]]\+\([^=[:space:]]\+\)[[:space:]]\+string.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    echo "No enums found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

# Returns go interface methods in file.
# Parameters:
#    Interface name
#    Go file
go-utils.get_interface_methods() {
  local interface="$1"
  local file="$2"
  awk "/^type $interface /{flag=1; print; next} flag && /^}/{flag=0} flag" $file |
    sed -e '1d' |
    awk "$awk_remove_go_comments"
}
