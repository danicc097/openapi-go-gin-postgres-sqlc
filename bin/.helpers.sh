#!/usr/bin/env bash
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
  if [[ -z "$TOP_LEVEL_DIR" ]]; then
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

# Accepts flags:
#    --no-kill    Do not immediately exit.
# It does not store information on the failed command. To keep track of all failures,
# use:
# 	for pid in "${pids[@]}"; do
# 	  cmd=$(jobs -l | grep "$pid")
# 	  wait -fn "$pid" || echo "Background job failed: $cmd"
# 	done
wait_without_error() {
  local -i err=0 werr=0
  local kill=true

  while [[ $# -gt 0 ]]; do # getopts fails in CI (non interactive)
    case "$1" in
    --no-kill) kill=false ;;
    *) echo "Invalid option: $1" >&2 ;;
    esac
    shift
  done

  while
    wait -fn || werr=$?
    ((werr != 127))
  do
    err=$werr
    ((err == 0)) || break # handle as soon as it happens
  done

  if ((err != 0)); then
    sleep 0.2 # wait for all stdout/err
    echo "A job failed" >&2
    if $kill; then
      kill -s SIGUSR1 $PROC
    fi
    return 1
  fi
}

retry() {
  local retries="$1"
  local command="${@:2}"
  local options="$-" # Get the current "set" options

  # disable set -e
  if [[ $options == *e* ]]; then
    set +e
  fi
  # disable custom tracebacks
  trap ':' ERR

  $command
  local exit_code=$?

  if [[ $options == *e* ]]; then
    set -e
  fi

  if [[ $exit_code -ne 0 && $retries -gt 0 ]]; then
    retry $((retries - 1)) "$command"
  else
    return $exit_code
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

array.add_prefix() {
  local prefix="$1"
  shift
  printf "%s\n" "${@/#/$prefix}"
}

array.add_suffix() {
  local suffix="$1"
  shift
  printf "%s\n" "${@/%/$suffix}"
}

# modifies in place
array.remove_element() {
  local -n array=$1
  local value=$2
  local temp_array=()

  # Loop through the array and keep only elements that are not equal to the value
  for element in "${array[@]}"; do
    if [ "$element" != "$value" ]; then
      temp_array+=("$element")
    fi
  done

  # Assign the modified array back to the original array
  array=("${temp_array[@]}")
}
# breaks when separator has spaces, e.g. " | "
# join_by() {
#   [ "$#" -ge 1 ] || return 1
#   local IFS="$1"
#   shift
#   printf '%s\n' "$*"
# }

to_lower_sentence() {
  local kebab=$(to_kebab "$1")
  echo "${kebab//-/ }"
}

to_snake() {
  local kebab=$(to_kebab "$1")
  echo "${kebab//-/_}"
}

# https://stackoverflow.com/questions/57804252/consistent-syntax-for-obtaining-output-of-a-command-efficiently-in-bash
# also see https://github.com/dimo414/bash-cache if needed for more expensive functions
declare -Ag memoized_to_pascal

# via nameref
to_pascal() {
  local -n __to_pascal_res="$1"
  local string="$2"

  local memoized="${memoized_to_pascal[$string]}"
  if [[ -n "$memoized" ]]; then
    __to_pascal_res="$memoized"
    return
  fi

  # Replace spaces with nothing and capitalize the following letter
  string="${string// \([a-z]\)/\U\1}"

  # Replace upper letters with space + lower
  string="${string//\([A-Z]\)/ \L\1}"

  string=${string//[_-]/ }

  local exceptions=("id" "api" "url" "http" "json" "html" "css")

  for word in $string; do
    if [[ " ${exceptions[*]} " =~ " $word " ]]; then
      __to_pascal_res+="${word^^}" # Uppercase the whole word
    else
      __to_pascal_res+="${word^}" # Capitalize the first letter
    fi
  done

  memoized_to_pascal["$string"]="$__to_pascal_res"
}

# via nameref
to_camel() {
  local -n __to_camel_res="$1"
  local string="$2"

  to_pascal __to_camel_res "$string"
  __to_camel_res="${__to_camel_res,}"
}

function to_kebab() {
  echo -n "$1" |
    sed '
      s/\([^A-Z+]\)\([A-Z0-9]\)/\1-\2/g;
      s/\([0-9]\)\([A-Z]\)/\1-\2/g;
      s/\([A-Z]\)\([0-9]\)/\1-\2/g;
      s/--/-/g;
      s/\([\/]\)-/\1/g
    ' |
    tr '[:upper:]' '[:lower:]'
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

# FIXME:
get_function_name_in_line_number() {
  local line_number="$1"
  local script_file="$2"
  local function_name

  local function_name=$(awk -v line="$line_number" '
    /^ *[a-zA-Z_][a-zA-Z0-9_]* *\(\) *{/,/}/ {
      if ($0 ~ /^ *[a-zA-Z_][a-zA-Z0-9_]* *\(\) *{/) {
        current_function = $1
        gsub(/\(\)/, "", current_function)
      }
      if ($0 ~ /^}/) {
        current_function = ""
      }
    }

    NR == line { print current_function; exit }
  ' "$script_file")

  echo "$function_name"
}

# Usage: trap 'show_tracebacks' ERR
show_tracebacks() {
  local err_code=$? # do not quote!
  set +o xtrace
  local bash_command=${BASH_COMMAND}
  # function_name=$(get_function_name_in_line_number ${BASH_LINENO[0]} ${BASH_SOURCE[1]})
  # if [[ -n $function_name ]]; then
  #   function_name="[$function_name]"
  # fi
  if [ "$err_code" -eq 130 ]; then
    exit 1
  fi

  if [[ $bash_command != xlog* && $bash_command != xerr* && ${#FUNCNAME[@]} -gt 2 ]]; then
    echo >&2
    printf "${RED}%0.s-${OFF}" $(seq "80") >&2
    echo >&2
    echo "${RED}Error in ${YELLOW}${BASH_SOURCE[1]##"$TOP_LEVEL_DIR/"}:${BASH_LINENO[0]}${OFF} ${CYAN}$function_name${OFF} (exited with status $err_code)${OFF}" >&2
    echo "${RED}Traceback of ${BASH_SOURCE[1]} (most recent call last):${OFF}" >&2
    for ((i = 0; i < ${#FUNCNAME[@]} - 1; i++)); do
      local funcname="${FUNCNAME[$i]}"
      [ "$i" -eq "0" ] && funcname=$bash_command
      echo -e "  ${MAGENTA}${BASH_SOURCE[$i + 1]##"$TOP_LEVEL_DIR/"}:${BASH_LINENO[$i]}${OFF}\\t$funcname" >&2
    done
  fi
  exit 1
}

# Cache given files and return if checksums match or an error code otherwise.
# Parameters:
#   - Output .md5 file
#   - Files or directories to cache
#   - Optionally pass glob patterns to exclude via --exclude [pattern].
#   - Optionally pass --no-regen disable external cache invalidation.
cache_all() {
  local excludes=()
  local args=()
  local no_force_regen=false
  #i=0 is still program name in "$@". Therefore continue up to <= too.
  for ((i = 1; i <= ${#@}; i++)); do
    arg="${!i}"
    case $arg in
    --no-regen)
      no_force_regen=true
      ;;
    --exclude)
      ((i++))
      excludes+=("${!i}")
      ;;
    *)
      args+=("$arg")
      ;;
    esac
  done

  if [ "${#args[@]}" -lt 2 ]; then
    echo "Usage: ${FUNCNAME[0]} [--exclude <pattern>] <output_cache_md5_path> <file_or_directory> [<file_or_directory> ...]" >&2
    return 1
  fi

  local output_file="${args[0]}"
  true >"$output_file.tmp"

  exclude_args=()
  for exclude in "${excludes[@]}"; do
    exclude_args+=('!' -path "$exclude")
  done

  for arg in "${args[@]:1}"; do
    if [ -d "$arg" ]; then
      find "$arg" -type f "${exclude_args[@]}" -exec md5sum {} + >>"$output_file.tmp"
    elif [ -f "$arg" ]; then
      md5sum "$arg" >>"$output_file.tmp"
    else
      echo "Invalid argument: $arg" >&2
      return 1
    fi
  done

  if cmp -s "$output_file" "$output_file.tmp"; then
    if test -z "$X_FORCE_REGEN" || $no_force_regen; then
      echo "${CYAN}Skipping generation (cached).${OFF} Regenerate with ${RED}--x-force-regen${OFF}"
      rm "$output_file.tmp"
      return 0
    fi
  fi

  mv "$output_file.tmp" "$output_file"
  return 1
}

# Block build if magic keyword is found in any file
# Args: keyword
search_stopship() {
  { { {
    stopship_keyword="$1"
    script_path=$(realpath --relative-to="$(git rev-parse --show-toplevel)" "$0")
    local matches
    matches=$(rg --files-with-matches --glob "!$script_path" --glob '!**/.git/*' --glob '!**/.venv/*' --glob '!**/vendor/*' --glob '!**/node_modules/*' --glob '!**/build/*' --glob '!*.pyc' --regexp="$stopship_keyword" "$(git rev-parse --show-toplevel)" || true)
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
  docker.postgres psql -qtAX -v ON_ERROR_STOP=on "$@"
}

# Drop and recreate database `db`. Defaults to POSTGRES_DB.
docker.postgres.drop_and_recreate_db() {
  local db="${1:POSTGRES_DB}"

  docker.postgres.wait_until_ready

  echo "${RED}${BOLD}Dropping database $db.${OFF}"
  docker.postgres dropdb --force --if-exists -f "$db"

  docker.postgres.create_db $db
}

# Create database `db`.
docker.postgres.create_db() {
  local db="$1"

  docker.postgres.wait_until_ready

  echo "${BLUE}${BOLD}Creating database $db.${OFF}"
  docker.postgres createdb $db -U "$POSTGRES_USER" -O "$POSTGRES_USER" ||
    echo "Skipping $db database creation"
}

# Stop running processes in `db`.
docker.postgres.stop_db_processes() {
  local db="$1"

  docker.postgres.wait_until_ready

  echo "${BLUE}${BOLD}Stopping any running processes for database $db.${OFF}"
  docker.postgres psql --no-psqlrc \
    -U "$POSTGRES_USER" \
    -d "postgres" \
    -c "select pg_terminate_backend(pid) \
        from pg_stat_activity \
        where datname='$db'" >/dev/null
}

docker.postgres.wait_until_ready() {
  pg_ready=0
  while [[ ! $pg_ready -eq 1 ]]; do
    docker.postgres \
      pg_isready -U "$POSTGRES_USER" || {
      echo "${YELLOW}Waiting for postgres database to be ready...${OFF}"
      sleep 1
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

AWK_REMOVE_GO_COMMENTS='
     /\/\*/ { f=1 } # set flag that is a block comment

     f==0 && !/^[ \t]*(\/\/|\/\*)/ { # skip // or /*
        print  # print non-commented lines
     }
     /\*\// { f=0 } # reset flag at end of block comment
'

# Stores go structs (including generic instances) in pkg to a given array.
# For 100% assurance use `ast-parser find-structs`
# Parameters:
#    Struct array (nameref)
#    Package directory or file
go-utils.find_structs() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t -O ${#__arr[@]} __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -ne '/\[/!s/type\(.*\)struct.*/\1/p') # /\[/! excludes lines containing [ right away

  local generic_structs=()
  go-utils.find_generic_structs generic_structs $pkg

  for generic_struct in "${generic_structs[@]}"; do
    mapfile -t -O ${#__arr[@]} __arr < <(find "$pkg" -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
      sed -n -e "s/^type \(.*\) = ${generic_struct}\[.*\].*/\1/p")
  done

  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No structs found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" ${__arr[@]}))
}

# Stores go generic struct definitions in pkg to a given array.
# Parameters:
#    Struct array (nameref)
#    Package directory or file
go-utils.find_generic_structs() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t -O ${#__arr[@]} __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -ne 's/^type \(.*\)\[.*\] struct.*/\1/p')

  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" ${__arr[@]}))
}

# Stores xo serial or bigserial PK generated types in pkg to a given array.
go-utils.find_db_ids_int() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find "$pkg" -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -ne 's/[\s]*type[[:space:]]*\([^[:space:]]*\)ID[[:space:]]*int.*/\1ID/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No db int IDs found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort -u < <(printf "%s\n" ${__arr[@]}))
}

# Stores xo uuid PK generated types in pkg to a given array.
go-utils.find_db_ids_uuid() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find "$pkg" -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -n -E 's/func New(.*)ID\(id uuid.UUID\) (.*)ID.*/\1ID/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No db uuid IDs found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort -u < <(printf "%s\n" ${__arr[@]}))
}

# Stores go test functions in package to a given array.
# Parameters:
#    Test function array (nameref)
#    Package directory
go-utils.find_test_functions() {
  local -n __arr="$1"
  local pkg="$2"

  mapfile -t __arr < <(
    find "$pkg" -maxdepth 1 -name "*_test.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
      sed -n -E 's/^\s*func\s*(Test[a-zA-Z0-9_]*)\(.*/\1/p'
  )

  mapfile -t __arr < <(printf "%s\n" ${__arr[@]} | grep -v "TestMain")

  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No test functions found in package in directory: $pkg"
  fi

  mapfile -t __arr < <(LC_COLLATE=C sort -u < <(printf "%s\n" ${__arr[@]}))
}

# Stores go struct fields to a given array.
# Parameters:
#    Struct name
#    Filename
#    Field array (nameref)
go-utils.struct_fields() {
  struct_name="$1"
  file_name="$2"
  local -n __arr="$3"

  struct_definition=$(awk -v struct="$struct_name" '
    $1 == "type" && $2 == struct {
      in_struct = 1;
      next;
    }
    in_struct {
      if ($1 == "}") {
        in_struct = 0;
      } else if ($1 != "") {
        print " " $1;
      }
    }
  ' "$file_name")
  while read -r line; do
    field_value=$(awk -v field="$line" '$1 == field {print $1}' <<<"$struct_definition")
    __arr+=("$field_value")
  done < <(echo "$struct_definition")
}

# Stores go interfaces in package to a given array.
# Parameters:
#    Interface array (nameref)
#    Package directory
go-utils.find_interfaces() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t -O ${#__arr[@]} __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -ne 's/[\s]*type\(.*\)interface.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No interfaces found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" ${__arr[@]}))
}

# Stores go enums in package to a given array.
# Parameters:
#    Enum array (nameref)
#    Package directory
go-utils.find_enums() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t -O ${#__arr[@]} __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -ne 's/.*type[[:space:]]\+\([^=[:space:]]\+\)[[:space:]]\+string.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    echo "No enums found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" ${__arr[@]}))
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
    awk "$AWK_REMOVE_GO_COMMENTS"
}

# Stores go custom types in package to a given array.
# Deprecated: use `ast-parser find-types`
# Parameters:
#    Custom types array (nameref)
#    Package directory
go-utils.find_all_types() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t -O ${#__arr[@]} __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$AWK_REMOVE_GO_COMMENTS" {} \; |
    sed -E -n 's/^type[[:space:]]+([A-Z][A-Za-z0-9_]+)[[:space:]].*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    echo "No types found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" ${__arr[@]}))
}

# Escape regular string for sed commands
escape_sed() {
  echo "$1" | sed -e 's/[\/&]/\\&/g'
}
