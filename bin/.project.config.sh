#!/bin/bash

source ".helpers.sh"
source ".project.usage.sh"
source ".project.dependencies.sh"

declare X_IGNORE_BUILD_ERRORS X_FORCE_REGEN X_NO_CONFIRMATION X_NO_GEN X_NO_CACHE X_ENV

readonly CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

readonly SPEC_PATH="openapi.yaml"

readonly BUILD_DIR="bin/build"
readonly TOOLS_DIR="bin/tools"
readonly PROTO_DIR="internal/pb"
readonly MIGRATIONS_DIR="db/migrations"
readonly CERTIFICATES_DIR="certificates"
readonly GOWRAP_TEMPLATES_DIR="internal/gowrap-templates"
readonly REPOS_DIR="internal/repos"
readonly PG_REPO_DIR="$REPOS_DIR/postgresql"
readonly XO_TEMPLATES_DIR="$PG_REPO_DIR/xo-templates"

readonly POSTGRES_TEST_DB="postgres_test"
readonly DUMPS_DIR="$HOME/openapi_go_gin_postgres_dumps"
pkg="$(head -1 go.mod)"
readonly GOMOD_PKG="${pkg#module *}"
# can only run with count=1 at most
readonly XO_TESTS_PKG="$GOMOD_PKG/internal/repos/postgresql/xo-templates/tests"
readonly OPID_AUTH="operationAuth.gen.json"
readonly PG_REPO_GEN="$PG_REPO_DIR/gen"
readonly CACHE=".generate.cache"
readonly SWAGGER_UI_DIR="internal/static/swagger-ui"
readonly MAX_FNAME_LOG_LEN=13

GEN_POSTGRES_DB="gen_db"

# determines if gen cache should be restore at program exit. 0|1
restore_cache=0
# stores the first executing function to track if caching gen is already running,
# to allow for nested xsetup.backup-gen and cache-cleanup inside multiple functions.
xsetup_backup_gen_caller=""

# stores the first executing function to determine if a migration
# is needed when running gen* functions which call each other
xsetup_gen_migrated=""

# stores the first executing function to determine if tools have been built
xsetup_tools_built=""

parse_args() {
  declare CMD="$1"

  # First comment lines automatically added to usage docs.
  while [[ "$#" -gt 0 ]]; do
    case $1 in
    --x-help)
      # Show help for a particular x function.
      COMMANDS=("$CMD")
      usage
      exit
      ;;
    --x-ignore-build-errors)
      # Proceeds with code generation regardless of tool rebuild errors.
      # Use in case of compilation errors that depend on generated code.
      # A clean gen run is required afterwards.
      export X_IGNORE_BUILD_ERRORS=1
      # export X_NO_CACHE=1 # on failed gen $CACHE is cleared, shouldn't be necessary
      ;;
    --x-force-regen)
      # Removes generation cache, forcing a new run.
      export X_FORCE_REGEN=1
      ;;
    --x-no-confirmation)
      # Bypasses confirmation messages. (WIP: Use `yes` in the meantime)
      export X_NO_CONFIRMATION=1
      ;;
    --x-no-gen)
      # Skips code generation steps.
      export X_NO_GEN=1
      ;;
    --x-no-cache)
      # Code generation backup is not restored on failure.
      # Please ensure there are no uncommitted changes in the current branch beforehand.
      export X_NO_CACHE=1
      ;;
    --x-env=*)
      # Environment to run commands in. Defaults to "dev".
      # Args: env
      export X_ENV="${1#--x-env=}"
      valid_envs="dev prod ci"
      if [[ ! " ${valid_envs[*]} " =~ " $X_ENV " ]]; then
        err "Valid environments: $valid_envs"
      fi
      ;;
    *)
      # will set everything else back
      args+=("$1")
      ;;
    esac
    shift
  done

  for arg in ${args[@]}; do
    set -- "$@" "$arg"
  done

  readonly X_IGNORE_BUILD_ERRORS X_FORCE_REGEN X_NO_CONFIRMATION X_NO_GEN X_NO_CACHE X_ENV
}

# If in completion mode, dynamically show commands to complete and exit.
# Globals:
#   COMMANDS
# Parameters:
#   None
# Outputs:
#   None
# Returns:
#   None
complete_commands() {
  if [[ -n $COMP_LINE ]]; then
    pre="${COMP_LINE##* }" # the part after the last space in the current command
    cur_commands=(${COMP_LINE%"$pre"})

    for c in "${COMMANDS[@]}"; do
      if [[ " ${cur_commands[*]} " =~ " ${c} " ]]; then
        xfn_specified=true
        break
      fi
    done

    for c in "${COMMANDS[@]}"; do
      test -z "${xfn_specified}" || break
      test -z "${pre}" -o "${c}" != "${c#"${pre}"}" -a "${pre}" != "${c}" && echo "${c} "
    done

    test -z "${xfn_specified}" && exit

    declare __x_options x_options_lines

    parse_x_options x_options_lines

    for c in "${x_options_lines[@]}"; do
      tmp="${c%%)*}"
      xopt="${tmp//\*/}"
      __x_options+=("$xopt")
    done

    declare -A __x_opts_seen
    for cmd in "${cur_commands[@]}"; do
      for opt in ${__x_options[@]}; do
        if [[ "$cmd" == *"$opt"* ]]; then
          __x_opts_seen[$opt]=true
          break
        fi
      done
    done

    for opt in ${__x_options[@]}; do
      [[ -n "${__x_opts_seen[$opt]}" ]] && continue
      if [[ ${opt:0:${#pre}} == "${pre,,}" ]]; then
        [[ "$opt" == "${pre,,}" ]] && continue # will have to be removed for inner completion
        if [[ "${opt,,}" =~ ^.*= ]]; then
          # TODO: could complete if choices found for opt (assoc array) eg --x-env has # Options: "dev prod ci"
          # which gets added to usage docs as arg: dev|prod|ci instead of using # Args: env
          # for validation, it can be generic:
          # if [[ ! " ${opts[${--x-(flag)}]} " =~ " $val " ]]; then
          #   err "Invalid value for --x-(flag): val. Allowed values: opts[${--x-(flag)}]"
          # fi`
          # and for completion itself we cycle through echo "${opt}<val1>", "${opt}<val2>",...
          # (can handle inner pre here, ie --x-env=de <TAB> outputs possible nested opts as well)
          echo "${opt}"
        else
          echo "${opt} "
        fi
      fi
    done

    exit
  fi
}
