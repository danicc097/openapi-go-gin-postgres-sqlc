#!/bin/bash

source ".helpers.sh"
source ".project.usage.sh"
source ".project.dependencies.sh"

source ".project.config.sh"

######################### x-functions setup #########################

xsetup.build-tools() {
  test -n "$xsetup_tools_built" && return

  xsetup_tools_built="${FUNCNAME[1]}"

  x.gen.build-tools
}

xsetup.backup-gen() {
  test -n "$xsetup_backup_gen_caller" && return

  xsetup_backup_gen_caller="${FUNCNAME[1]}"

  mkdir -p "$CACHE"

  gen_backup_branch="backup-gen-$(uuidgen)"
  gen_backup_stash_name="backup-stash-$gen_backup_branch"

  uuidgen >backup-gen-stash-dummy.txt
  git stash push -m "$gen_backup_stash_name" --include-untracked
  git checkout -b "$gen_backup_branch" &>/dev/null
  git stash apply "stash^{/$gen_backup_stash_name}" &>/dev/null

  restore_cache=1
}

xsetup.backup-gen.cleanup() {
  if [[ "$xsetup_backup_gen_caller" = "${FUNCNAME[1]}" ]]; then
    restore_cache=0
  fi
}

xsetup.backup-gen.restore() {
  echo "Restoring previously generated code (branch $gen_backup_branch)..."
  wait

  git reset --hard && git clean -df &>/dev/null
  git stash apply "stash^{/$gen_backup_stash_name}" &>/dev/null

  rm -rf "$CACHE/*.md5" # cache is not up to date due to forced exit and restore
  # TODO: delete specific folders like $CACHE/{counterfeiter,gowrap} only if
  # failure on those tools specifically
}

xsetup.drop-and-migrate-gen-db() {
  test -n "$xsetup_gen_migrated" && return
  xsetup_gen_migrated=1

  (
    POSTGRES_DB="$GEN_POSTGRES_DB"

    x.db.drop
    x.migrate up
  )
}

######################### x-functions #########################

# Check build dependencies are met and prompt to install if missing.
x.check-build-deps() {
  # TODO: should check go installs in gobin are the version we expect when running x.gen
  # via x.install-tools.check
  { { {
    mkdir -p $TOOLS_DIR

    while IFS= read -r line; do
      [[ $line =~ ^declare\ -f\ check\.bin\. ]] && BIN_CHECKS+=("${line##declare -f check.bin.}")
      [[ $line =~ ^declare\ -f\ install\.bin\. ]] && BIN_INSTALLS+=("${line##declare -f install.bin.}")
    done < <(declare -F)

    echo "Checking dependencies..."
    for bin in "${BIN_CHECKS[@]}"; do
      # local r
      # r="$(...)" # redirect to var while also streaming unbuffered output with | tee /dev/tty
      if "check.bin.$bin"; then
        continue
      fi

      if ! element_in_array "$bin" "${BIN_INSTALLS[@]}"; then
        echo "No automatic installation available. Please install $bin manually and retry"
        exit 1
      fi

      with_tty confirm "Do you want to install $bin now?" || exit 1

      echo "Installing $bin..."
      if ! "install.bin.$bin"; then
        err "$bin installation failed"
      fi
      echo "Installed $bin..."
    done
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Check dependencies and fetch required tools.
x.bootstrap() {
  { { {
    git submodule update --init --recursive # sync later on with git `git submodule update --force --recursive --remote`

    x.check-build-deps
    x.backend.sync-deps
    x.install-tools
    x.setup.swagger-ui
    x.gen.build-tools

    cd frontend
    pnpm i --frozen-lockfile
    cd -

    cd e2e
    pnpm i --frozen-lockfile
    cd -

    traefik_dir="$HOME/traefik-bootstrap"
    with_tty confirm "Do you want to setup and run traefik (install dir: $traefik_dir)?" && x.setup.traefik "$traefik_dir"
    echo "${RED}Make sure to add \`complete -o nospace -C project project\` to your ~/.bashrc for completion.${OFF}"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Install miscellaneous tool binaries locally.
x.install-tools() {
  { { {
    echo "Installing tools..."

    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 &
    # go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0 &
    go install github.com/danicc097/sqlc/cmd/sqlc@custom &
    # for easier test search.
    # NOTE: unrelated run test broken discovery: https://github.com/golang/vscode-go/issues/2719
    go install github.com/danicc097/go-test-renamer@v0.2.0 &
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 &
    go install github.com/joho/godotenv/cmd/godotenv@latest &
    go install github.com/tufin/oasdiff@latest &
    go install golang.org/x/tools/cmd/goimports@latest &
    go install mvdan.cc/gofumpt@latest &
    go install github.com/danicc097/air@latest &
    go install github.com/danicc097/xo/v5@v5.1.0 &
    go install github.com/mikefarah/yq/v4@v4.34.2 &
    go install github.com/hexdigest/gowrap/cmd/gowrap@latest &
    go install golang.org/x/tools/cmd/stringer@latest &

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 &
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 &

    GO111MODULE=off go get -u github.com/maxbrunsfeld/counterfeiter &

    # install node libs with --prefix $TOOLS_DIR, if any
    # ...

    wait
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Fetch latest Swagger UI bundle.
x.setup.swagger-ui() {
  { { {
    local name

    name="$(curl --silent "https://api.github.com/repos/swagger-api/swagger-ui/releases/latest" | jq -r ".. .tag_name? // empty")"
    curl -fsSL "github.com/swagger-api/swagger-ui/archive/refs/tags/$name.tar.gz" -o swagger-ui.tar.gz
    tar xf swagger-ui.tar.gz swagger-ui-"${name#*v}"/dist --one-top-level=swagger-ui --strip-components=2
    rm swagger-ui.tar.gz
    mkdir -p $SWAGGER_UI_DIR
    mv swagger-ui/* $SWAGGER_UI_DIR
    rm -r swagger-ui
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Run pre-generation scripts in the internal package.
x.gen.pregen() {
  xsetup.backup-gen
  xsetup.drop-and-migrate-gen-db
  xsetup.build-tools
  { { {
    POSTGRES_DB="$GEN_POSTGRES_DB"

    echo "Running generation"

    codegen validate-spec -env=".env.$X_ENV"

    sync_spec_with_db

    ######## Ensure consistent style for future codegen

    echo "Applying PascalCase to operation IDs in $SPEC_PATH"
    spec_content=$(<$SPEC_PATH)

    # outputs safe double-quoted paths for yq
    # https://github.com/mikefarah/yq/issues/1295
    mapfile -t opid_paths < <(yq e '
      .paths[][].operationId
      | path
      | with(.[] | select(contains(".") or contains("/") or contains("{")); . = "\"" + . + "\"")
      | join(".")
    ' $SPEC_PATH)
    mapfile -t opids < <(yq e ".paths[][].operationId" $SPEC_PATH)

    # construct single yq call
    local ops=()
    for i in ${!opids[@]}; do
      new_opid="$(to_pascal ${opids[$i]})"
      ops+=(".${opid_paths[$i]}=\"${new_opid}\"") # cant have space
    done

    yq_op=$(join_by " | " ${ops[*]})

    spec_content=$(yq e "$yq_op" < <(echo "$spec_content"))
    echo "$spec_content" >$SPEC_PATH

    echo "Updating roles and scopes in $SPEC_PATH"
    ######## Sync enums with external sources and validate
    ######## external json files are the source of truth
    # arrays can't be nested
    declare -A enum_src_files=(
      [Scope]="scopes.json"
      [Role]="roles.json"
    )
    declare -A enum_vext=(
      [Scope]="x-required-scopes"
      [Role]="x-required-role"
    )
    declare -A enum_values=()
    for enum in ${!enum_src_files[@]}; do
      [[ $(yq e ".components.schemas | has(\"$enum\")" $SPEC_PATH) = "false" ]] &&
        yq e ".components.schemas.$enum.type = \"string\"" -i $SPEC_PATH

      local src_file="${enum_src_files[$enum]}"
      vendor_ext="${enum_vext[$enum]}"

      enums=$(yq -P --output-format=yaml '.[] | key' $src_file)
      mapfile -t enums <<<$enums

      src_comment="$src_file keys"
      replace_enum_in_spec "$enum" enums "$src_comment"

      mapfile spec_enums < <(yq e ".paths[][].$vendor_ext | select(length > 0)" $SPEC_PATH)
      spec_enums=("${spec_enums[*]//- /}")
      mapfile -t spec_enums < <(printf "\"%s\"\n" ${spec_enums[*]})
      mapfile -t clean_enums < <(printf "\"%s\"\n" ${enums[*]})
      # ensure only existing enums from src_file are used
      for spec_enum in "${spec_enums[@]}"; do
        [[ ! " ${clean_enums[*]} " =~ " ${spec_enum} " ]] && err "$spec_enum is not a valid '$enum'"
      done

      enum_list=$(printf ",\"%s\"" "${enums[@]}")
      enum_list="[${enum_list:1}]"
      enum_values[$enum]=$enum_list
    done

    yq -e ".definitions.Operation.properties +=
        {
          \"${enum_vext[Role]}\": {
            \"type\": \"string\",
            \"enum\": ${enum_values[Role]}
          },
          \"${enum_vext[Scope]}\": {
            \"type\": \"array\",
            \"items\": {\"enum\": ${enum_values[Scope]}}
          }
        }" -i -oj .vscode/openapi-schema.json

    ######## Generate shared policies once the spec has been validated
    echo "Writing shared auth policies"

    yq -o=json e "
    .paths[][]
    | explode(.)
    | {
      .operationId: {
        \"scopes\": .x-required-scopes,
        \"role\": .x-required-role,
        \"requiresAuthentication\": has(\"security\")
        }
      }
    | select(.[]) as \$i ireduce ({}; . + \$i)
  " $SPEC_PATH >$OPID_AUTH

    codegen pre -env=".env.$X_ENV" -op-id-auth="$OPID_AUTH"

    update_spec_with_structs

    remove_schemas_marked_to_delete
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

sync_spec_with_db() {
  cache_all "$CACHE/db.md5" bin/project bin/.helpers.sh .env.$X_ENV db/ $SPEC_PATH && return 0

  sync_db_enums_with_spec

  ######## Sync projects and related project info

  mapfile -t db_projects < <(dockerdb_psql -d $POSTGRES_DB -c "select name from projects;" 2>/dev/null)

  [[ ${#db_projects[@]} -gt 0 ]] || err "No projects found in database $POSTGRES_DB"
  replace_enum_in_spec "Project" db_projects "projects table"

  for project in ${db_projects[@]}; do
    ### kanban steps
    mapfile -t kanban_steps < <(dockerdb_psql -d $POSTGRES_DB -c "
          select name from kanban_steps where project_id = (
            select project_id from projects where name = '$project'
            );" 2>/dev/null)
    [[ ${#kanban_steps[@]} -gt 0 ]] || {
      echo "${YELLOW}[WARNING] No kanban steps found for project '$project' in database $POSTGRES_DB${OFF}" && continue
    }
    schema_name="$(to_pascal $project)KanbanSteps"
    replace_enum_in_spec "$schema_name" kanban_steps "kanban_steps table"

    ### work item types
    mapfile -t work_item_types < <(dockerdb_psql -d $POSTGRES_DB -c "
          select name from work_item_types where project_id = (
            select project_id from projects where name = '$project'
            );" 2>/dev/null)
    [[ ${#work_item_types[@]} -gt 0 ]] || {
      echo "${YELLOW}[WARNING] No work item types found for project '$project' in database $POSTGRES_DB${OFF}" && continue
    }
    schema_name="$(to_pascal $project)WorkItemTypes"
    replace_enum_in_spec "$schema_name" work_item_types "work_item_types table"
  done

  generate_models_mappings
}

# for manually inserted elements via migrations, e.g. projects, kanban_steps, work_item_type,
# generate 2-way maps id<- ->name to save up useless db calls and make logic switching
# and repos usage much easier
generate_models_mappings() {
  local model_mappings_path="internal/models_mappings.gen.go"
  cat <<EOF >$model_mappings_path
// Code generated by project. DO NOT EDIT.

package internal

import "$GOMOD_PKG/internal/models"

EOF

  mapfile -t projects_rows < <(dockerdb_psql -d $POSTGRES_DB -c "select project_id,name from projects;" 2>/dev/null)
  generate_model_mappings_dicts Project projects_rows

  for project in ${db_projects[@]}; do
    mapfile -t kanban_steps_rows < <(dockerdb_psql -d $POSTGRES_DB -c "
          select kanban_step_id,name from kanban_steps where project_id = (
            select project_id from projects where name = '$project'
            );" 2>/dev/null)
    [[ ${#kanban_steps_rows[@]} -gt 0 ]] || continue
    prefix="$(to_pascal "$project")KanbanSteps"
    generate_model_mappings_dicts $prefix kanban_steps_rows

    kanban_steps_rows=()
    mapfile -t kanban_steps_rows < <(dockerdb_psql -d $POSTGRES_DB -c "
          select kanban_step_id,step_order from kanban_steps where project_id = (
            select project_id from projects where name = '$project'
            );" 2>/dev/null)
    [[ ${#kanban_steps_rows[@]} -gt 0 ]] || continue
    prefix="$(to_pascal "$project")KanbanSteps"
    echo "var (
    ${prefix}StepOrderByID = map[int]int{
  " >>$model_mappings_path
    for row in "${kanban_steps_rows[@]}"; do
      first=$(cut_first "$row" "|") # always safe
      mapfile -t arr <<<"${first}"
      local id="${arr[0]}"
      local kanban_step="${arr[1]}"
      echo "${id}: ${kanban_step}," >>$model_mappings_path
    done
    echo "})" >>$model_mappings_path
  done

  for project in ${db_projects[@]}; do
    mapfile -t work_item_types_rows < <(dockerdb_psql -d $POSTGRES_DB -c "
          select work_item_type_id,name from work_item_types where project_id = (
            select project_id from projects where name = '$project'
            );" 2>/dev/null)
    [[ ${#work_item_types_rows[@]} -gt 0 ]] || continue
    prefix="$(to_pascal "$project")WorkItemTypes"
    generate_model_mappings_dicts $prefix work_item_types_rows
  done

  gofumpt -w $model_mappings_path
}

# generates dictionaries for existing database elements, meant for those
# inserted exclusively via migrations
generate_model_mappings_dicts() {
  local prefix="$1"
  local -n __arr="$2" # db rows
  echo "var (
	${prefix}NameByID = map[int]models.${prefix}{
  " >>$model_mappings_path
  for row in "${__arr[@]}"; do
    first=$(cut_first "$row" "|") # always safe
    mapfile -t arr <<<"${first}"
    local id="${arr[0]}"
    local name="${arr[1]}"
    echo "${id}: models.${prefix}$(to_pascal "$name")," >>$model_mappings_path
  done
  echo "}
	${prefix}IDByName = map[models.${prefix}]int{
  " >>$model_mappings_path
  for row in "${__arr[@]}"; do
    first=$(cut_first "$row" "|") # always safe
    mapfile -t arr <<<"${first}"
    local id="${arr[0]}"
    local name="${arr[1]}"
    echo "models.${prefix}$(to_pascal "$name"): ${id}," >>$model_mappings_path
  done
  echo "})
  " >>$model_mappings_path
}

clean_yq_array() {
  local -n __arr="$1"
  __arr=("${__arr[*]//- /}")
  mapfile -t __arr < <(printf "\"%s\"\n" ${__arr[*]})
  echo ${__arr[@]}
}

# Run post-generation scripts in the internal package.
x.gen.postgen() {
  xsetup.backup-gen
  xsetup.drop-and-migrate-gen-db
  xsetup.build-tools
  { { {
    POSTGRES_DB="$GEN_POSTGRES_DB"

    echo "Running generation"

  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Generate type-safe Go code from SQL.
x.gen.sqlc() {
  xsetup.backup-gen
  xsetup.drop-and-migrate-gen-db
  { { {
    echo "Running generation"
    rm -f "$PG_REPO_GEN"/db/*.sqlc.go
    sqlc generate --experimental -f "$PG_REPO_DIR"/sqlc.yaml || err "Failed sqlc generation"
    rm -f "$PG_REPO_GEN"/db/models.go # sqlc enums
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Automatically generate CRUD and index queries with joins based on existing indexes from a Postgres schema.
x.gen.xo() {
  xsetup.backup-gen
  xsetup.drop-and-migrate-gen-db
  { { {
    cache_all "$CACHE/xo.md5" bin/project bin/.helpers.sh .env.$X_ENV go.mod db/ $XO_TEMPLATES_DIR/ && return 0

    echo "Running generation"

    rm -rf "$PG_REPO_GEN"/db/*.xo.go

    mkdir -p "$PG_REPO_GEN"/db
    xo_schema -o "$PG_REPO_GEN"/db --debug \
      --schema public \
      --ignore "*.created_at" \
      --ignore "*.updated_at" || err "Failed xo public schema generation" &

    mkdir -p "$PG_REPO_GEN"/db/v
    xo_schema -o "$PG_REPO_GEN"/db/v \
      --schema v \
      --main-schema-pkg $GOMOD_PKG/"$PG_REPO_GEN"/db ||
      err "Failed xo v schema generation" &

    mkdir -p "$PG_REPO_GEN"/db/cache
    xo_schema -o "$PG_REPO_GEN"/db/cache \
      --schema cache \
      --main-schema-pkg $GOMOD_PKG/"$PG_REPO_GEN"/db ||
      err "Failed xo cache schema generation" &

    wait

    files=$(find "$PG_REPO_GEN/db" \
      -name "*.go")
    goimports -w $files
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Generate a type-safe SQL builder.
x.gen.jet() {
  xsetup.backup-gen
  xsetup.drop-and-migrate-gen-db
  xsetup.build-tools
  { { {
    # results may be combined with xo's *Public structs and not reinvent the wheel for jet.
    # should not be hard to generate all adapters at once jet->xo *Public in a new file alongside jet gen.
    # in the end fields are the same name if goName conventions are followed (configurable via custom jet cmd)
    # if it gives problems for some fields (ID, API and the like)
    echo "Running generation"

    local gen_path="$PG_REPO_GEN/jet"
    local schema=public
    rm -rf "$gen_path"
    {
      jet -dbname="$GEN_POSTGRES_DB" -env=.env."$X_ENV" --out=./"$gen_path" --schema=$schema
      mv "./$gen_path"/$GEN_POSTGRES_DB/* "$gen_path"
      rm -r "./$gen_path/$GEN_POSTGRES_DB/"
    }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Generate interface wrappers with common logic: tracing, timeout...
x.gen.gowrap() {
  xsetup.backup-gen
  { { {
    local repos="$REPOS_DIR/repos.go"

    echo "Running generation"

    local cache="$CACHE/gowrap"
    local suffixes=(
      "retry:with_retry"
      "timeout:with_timeout"
      "otel:with_otel"
      "prometheus:with_prometheus" # TODO: https://last9.io/blog/native-support-for-opentelemetry-metrics-in-prometheus/ https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/pkg/translator/prometheus
    )

    local repo_interfaces=()
    find_go_interfaces repo_interfaces $repos

    mkdir -p "$cache"

    local updated_ifaces=()
    for iface in ${repo_interfaces[@]}; do
      iface_content="$(find_go_interface_content $iface $repos)"
      if diff "$cache/$iface" <(echo "$iface_content") &>/dev/null && [[ $X_FORCE_REGEN -eq 0 ]]; then
        continue
      fi

      for suffix in ${suffixes[@]}; do
        {
          IFS=":" read -r -a arr <<<${suffix}
          local tmpl="${arr[0]}"
          local suffix="${arr[1]}"
          gowrap gen \
            -g \
            -p $GOMOD_PKG/$REPOS_DIR \
            -i $iface \
            -t "$GOWRAP_TEMPLATES_DIR/$tmpl.tmpl" \
            -o "$REPOS_DIR/reposwrappers/${iface,,}_$suffix.gen.go"
        } &
      done

      echo "$iface_content" >"$cache/$iface"
      updated_ifaces+=("$iface")
    done

    wait_without_error

    if [[ ${#updated_ifaces[@]} -gt 0 ]]; then
      echo "Updated repo interfaces: ${updated_ifaces[*]}"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Generate Go client and server from spec.
x.gen.client-server() {
  { { {
    echo "Running generation"
    paths=".openapi.paths.yaml"

    # hack to get separate types generation
    sed "s/\$ref: '\#\//\$ref: '$SPEC_PATH\#\//g" $SPEC_PATH >"$paths"
    # yq e 'del(.components)' -i "$paths" # dont delete since recent oapi-codegen does some checks even if we are not generating types
    go build -o $BUILD_DIR/oapi-codegen cmd/oapi-codegen/main.go || [[ -n "$X_IGNORE_BUILD_ERRORS" ]] # templates are embedded, required rebuild

    oapi-codegen --config internal/models/oapi-codegen-types.yaml "$SPEC_PATH" || err "Failed types generation"
    oapi-codegen --config internal/rest/oapi-codegen-server.yaml --models-pkg models "$paths" || err "Failed server generation"
    oapi-codegen --config internal/client/oapi-codegen-client.yaml "$SPEC_PATH" || err "Failed client generation"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Generate mocks for specified interfaces.
x.gen.counterfeiter() {
  # This shouldn't pose any problems, the interface is the only input to counterfeiter.
  { { {
    tfidfpb_dir="internal/pb/python-ml-app-protos/tfidf/v1"
    envvar="internal/envvar/envvar.go"
    repos="$REPOS_DIR/repos.go"
    tfidfpb="$tfidfpb_dir/service_grpc.pb.go"

    declare -A ifaces
    ifaces=(
      ["Provider-$envvar"]="internal/envvar/envvartesting/provider.gen.go"
      ["User-$repos"]="$REPOS_DIR/repostesting/user.gen.go"
      ["Notification-$repos"]="$REPOS_DIR/repostesting/notification.gen.go"
      ["Project-$repos"]="$REPOS_DIR/repostesting/project.gen.go"
      ["Team-$repos"]="$REPOS_DIR/repostesting/team.gen.go"
      ["MovieGenreClient-$tfidfpb"]="$tfidfpb_dir/v1testing/movie_genre_client.gen.go"
      ["MovieGenreServer-$tfidfpb"]="$tfidfpb_dir/v1testing/movie_genre_server.gen.go"
    )

    local cache="$CACHE/counterfeiter"
    local updated_ifaces=()

    mkdir -p "$cache"

    for key in ${!ifaces[@]}; do
      input_path="${key#*-}"
      iface="${key%%-*}"
      iface_content="$(find_go_interface_content $iface $input_path)"

      cache_entry="$cache/$(echo "$key" | base64)"
      if diff "$cache_entry" <(echo "$iface_content") &>/dev/null && [[ $X_FORCE_REGEN -eq 0 ]]; then
        continue
      fi

      counterfeiter -o "${ifaces[$key]}" "$input_path" "$iface" 2>&1 &
      echo "$iface_content" >"$cache_entry"
      updated_ifaces+=("$key")
    done

    wait_without_error

    # counterfeiter is unaware of grpc's obscure mustEmbedUnimplemented***() for forward server compatibility
    if ! grep -q 'v1\.UnimplementedMovieGenreServer' $tfidfpb_dir/v1testing/movie_genre_server.gen.go; then
      sed -i '/type FakeMovieGenreServer struct {/a v1\.UnimplementedMovieGenreServer' $tfidfpb_dir/v1testing/movie_genre_server.gen.go
    fi

    if [[ ${#updated_ifaces[@]} -gt 0 ]]; then
      echo "Updated repo interfaces: ${updated_ifaces[*]}"
    fi

  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Generate the required servers/clients for relevant services.
x.gen.proto() {
  { { {
    local import_path="python-ml-app-protos/tfidf/v1"
    local filename="internal/python-ml-app-protos/tfidf/v1/service.proto"

    mkdir -p internal/pb
    # Plugins are no longer supported by protoc-gen-go.
    # Instead protoc-gen-go-grpc and the go package (in proto or via M flag) are required
    echo "Running generation"
    protoc \
      --go-grpc_out=internal/pb/. \
      --go_out=internal/pb/. \
      --go-grpc_opt=M${filename}=${import_path},paths=import \
      --go_opt=M${filename}=${import_path},paths=import \
      internal/python-ml-app-protos/tfidf/v1/service.proto || err "Failed proto generation"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Run frontend code generation.
x.gen.frontend() {
  xsetup.backup-gen
  { { {
    export PATH=frontend/node_modules/.bin:$PATH

    SCHEMA_OUT="frontend/src/types/schema.d.ts"
    orval_config="frontend/orval.config.ts"

    cache_all "$CACHE/frontend.md5" bin/project bin/.helpers.sh .env.$X_ENV $SPEC_PATH $orval_config frontend/scripts/ frontend/package.json && return 0

    config_template_setup frontend # no need to run if cached .env

    mkdir -p frontend/src/types
    rm -rf frontend/src/gen

    {
      node frontend/scripts/generate-client-validator.js
      # TODO allow custom validate.ts baked into fork
      rm -f frontend/src/client-validator/gen/validate.ts
      find frontend/src/client-validator/gen/ -type f -exec \
        sed -i "s/from '.\/validate'/from '..\/validate'/g" {} \;
    }

    {
      v="$(openapi-typescript --version)"
      openapi-typescript $SPEC_PATH --output "$SCHEMA_OUT" --path-params-as-types --prettier-config .prettierrc
      echo "/* Generated by openapi-typescript $v */
/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable */
// @ts-nocheck
export type schemas = components['schemas']
" | cat - "$SCHEMA_OUT" >/tmp/out && mv /tmp/out "$SCHEMA_OUT"
    } &
    orval --config $orval_config &

    wait_without_error
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Run e2e code generation.
x.gen.e2e() {
  xsetup.backup-gen
  { { {
    source .env.e2e

    export PATH=e2e/node_modules/.bin:$PATH

    orval_config="e2e/orval.config.ts"

    cache_all "$CACHE/e2e.md5" bin/project bin/.helpers.sh .env.$X_ENV $SPEC_PATH $orval_config e2e/package.json && return 0

    config_template_setup e2e # no need to run if cached .env

    rm -rf e2e/client/gen

    orval --config $orval_config
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Run all codegen and postgen commands for the project.
x.gen() {
  [[ -n $X_NO_GEN ]] && return
  xsetup.backup-gen # Modification of vars inside would be local to subshell (caused by pipeline)
  xsetup.drop-and-migrate-gen-db
  xsetup.build-tools
  { { {
    echo "Running code generation"

    x.lint.sql

    # TODO try use gnu parallel and exit when anyone fails
    # (with current setup the whole x.gen is executed and only then
    # the trap on SIGUSR1 coming from `err` is run - now temporarily using kill 0 in `err` instead on sending SIGUSR1)
    # shopt -s inherit_errexit
    go generate ./... &
    x.gen.gowrap &
    x.gen.proto &
    x.gen.xo &
    x.gen.sqlc &
    x.gen.jet &

    wait_without_error

    {
      x.gen.pregen
      x.gen.client-server
    } &
    x.gen.counterfeiter & # delay since it depends on generated output (xo, proto...)

    wait_without_error

    x.gen.postgen

    # restart is not robust
    # vscode will randomly lose connection when restarting
    # for pid in $(pidof gopls); do
    #   restart_pid $pid &
    # done

    x.gen.frontend &
    x.gen.e2e &

    wait_without_error
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Build code generation custom tools.
x.gen.build-tools() {
  { { {
    generate_structs_map # openapi-go requires structs already compiled

    out_dir=$BUILD_DIR

    mkdir -p $out_dir
    for cmd in jet oapi-codegen codegen; do
      echo "Building $cmd..."
      { go build -o $out_dir/$cmd cmd/$cmd/main.go || [[ -n "$X_IGNORE_BUILD_ERRORS" ]]; } &
    done

    wait
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Lint the entire project.
x.lint() {
  { { {
    x.lint.sql &
    x.lint.go &
    x.lint.frontend &
    wait
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Format Go files.
x.lint.go() {
  { { {
    files=$(find . \
      -not -path "**/$PROTO_DIR/*" \
      -not -path "**/$PG_REPO_GEN/*" \
      -not -path "**/testdata/*" \
      -not -path "**/node_modules/*" \
      -not -path "**/.venv/*" \
      -not -path "**/*.cache/*" \
      -not -path "**/vendor/*" \
      -not -path "**/*.gen.*" \
      -not -path "**/*.xo.go" \
      -name "*.go")
    goimports -w $files || echo "Linting failed"
    gofumpt -w $files || echo "Linting failed"
    golangci-lint run --config=.golangci.yml --fix &>/dev/null || true
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Format frontend files.
x.lint.frontend() {
  { { {
    cd frontend
    pnpm run lint:fix
    echo "Success"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Format SQL files.
x.lint.sql() {
  { { {
    SQL_DIRS=(
      "$REPOS_DIR"
      "db"
    )
    for slq_dir in ${SQL_DIRS[@]}; do
      pg_format \
        --spaces 2 \
        --wrap-limit 130 \
        --function-case 2 \
        --keyword-case 1 \
        --placeholder "sqlc\\.(arg|narg)\\(:?[^)]*\\)" \
        --inplace \
        --keep-newline \
        --comma-start \
        --nogrouping \
        $(find "$slq_dir" -name '*.*sql')
    done

    echo "Success"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Run required backend pre-test setup: services, database cleanup, codegen...
# Can be called independently, e.g. before running tests through an IDE.
x.test.backend.setup() {
  xsetup.backup-gen # Modification of vars inside would be local to subshell (caused by pipeline)
  { { {
    # NOTE: tests run independently in Go so we can't have a function be called and run
    # only once before any test starts
    run_shared_services up -d --build --remove-orphans --wait
    x.gen
    # no need to migrate, done on every test run internally
    drop_and_recreate_db $POSTGRES_TEST_DB
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
  xsetup.backup-gen.cleanup
}

# Test backend. Accepts `go test` parameters.
# Args: [...]
x.test.backend() {
  { { {
    x.test.xo

    yes y 2>/dev/null | POSTGRES_DB=$POSTGRES_TEST_DB x.migrate down || true
    POSTGRES_DB=$POSTGRES_TEST_DB x.migrate up # for post-migration scripts

    local cache_opt="-count=1"
    cache_all "$CACHE/go-test.md5" .env.$X_ENV db/ && cache_opt=""

    set -x
    APP_ENV="$X_ENV" go test -tags skipxo $cache_opt "$@" ./...
    set +x

  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Test frontend.
# Args: [...]
x.test.frontend() {
  # TODO accept vitest args
  config_template_setup frontend

  cd frontend
  pnpm run test:no-watch
  pnpm run test-types:no-watch
}

# Test frontend on file changes.
# Args: [...]
x.test.frontend.watch() {
  config_template_setup frontend

  cd frontend
  pnpm run test
}

# Run custom xo generation script tests. Accepts `go test` parameters.
# Args: [...]
x.test.xo() {
  { { {
    POSTGRES_DB="$POSTGRES_TEST_DB"
    GEN_POSTGRES_DB="$POSTGRES_TEST_DB"

    drop_and_recreate_db $POSTGRES_DB # FIXME wrong xo gen when tables with equal names in public schema exist as well

    echo "Running xo template tests"

    test_dir="$XO_TEMPLATES_DIR/tests"
    dockerdb_psql -d $POSTGRES_DB <"$test_dir/schema.sql" # need correct schema for xo gen
    echo "Schema loaded"

    rm -rf "$test_dir/got"
    mkdir -p "$test_dir/got"
    xo_schema -o "$test_dir/got" --debug \
      --schema xo_tests \
      --ignore "*.created_at" \
      --ignore "*.updated_at" || err "Failed xo xo_tests schema generation"

    files=$(find "$test_dir/got" \
      -name "*.go")
    goimports -w $files
    gofumpt -w $files

    APP_ENV="$X_ENV" go test -count=1 "$XO_TESTS_PKG" "$@" ||
      err "xo tests failed"

    if ! $test_dir/diff_check; then
      with_tty confirm "Do you want to update test snapshot with current changes?" && { # must redirect inside xlog pipeline
        rm -rf "$test_dir/snapshot"
        cp -r "$test_dir/got" "$test_dir/snapshot"
      }
    fi

  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Test backend on file changes. Accepts `go test` parameters.
# Args: [...]
x.test.backend.watch() {
  clear

  latency=2
  while true; do
    inotifywait \
      --monitor ./internal/**/* \
      --event=close_write \
      --format='%T %f' \
      --timefmt='%s' |
      while read -r event_time event_file 2>/dev/null || sleep $latency; do
        clear

        # NOTE: no count=1 default to allow caching
        { APP_ENV="$X_ENV" go test -tags skipxo "$@" ./...; } && echo "${GREEN}✓ All tests passing${OFF}"
      done
  done
}

# Run backend with hot-reloading.
x.run.backend-hr() {
  # TODO replace healthcheck with adhoc calls and bring services up in btaches
  # to prevent either bombarding with req or having to wait too long at startup.
  # see https://github.com/moby/moby/issues/33410
  run_shared_services up -d --build --remove-orphans --wait
  setup_swagger_ui
  # NOTE: building binary very unreliable, leads to bin not found.
  air \
    --build.pre_build_cmd "$pre_build_cmd" \
    --build.cmd "" \
    --build.bin "go run ./cmd/rest-server/ -env=.env.$X_ENV" \
    --build.include_ext "go" \
    --build.exclude_regex ".gen.go,_test.go" \
    --build.exclude_dir ".git,tmp,$PROTO_DIR,$PG_REPO_GEN,**/testdata,vendor,frontend,external,*.cache,$CACHE,$TOOLS_DIR" \
    --build.stop_watch "internal/rest/,internal/services/" \
    --build.delay 1000 \
    --build.exclude_unchanged "true" |
    sed -e "s/^/${BLUE}[Air]${OFF} /"
}

# Run frontend with hot-reloading.
x.run.frontend() {
  config_template_setup frontend
  cd frontend
  pnpm run dev |
    sed -e "s/^/${GREEN}[Vite]${OFF} /"
}

# Run all project services with hot reload enabled in dev mode.
x.run-dev() {
  run_hot_reload

  while true; do
    sleep 1000
  done

  # TODO fix won't kill children
  # next_allowed_run=$(date +%s)
  # latency=3
  # close_write event, else duplicated, tripl. events -> race condition
  # while true; do
  #   inotifywait \
  #     --monitor $SPEC_PATH \
  #     --event=close_write \
  #     --format='%T %f' \
  #     --timefmt='%s' |
  #     while read -r event_time event_file 2>/dev/null || sleep $latency; do
  #       if [[ $event_time -ge $next_allowed_run ]]; then
  #         next_allowed_run=$(date --date="${latency}sec" +%s)

  #         kill_descendants || true

  #         run_hot_reload
  #       fi
  #     done
  # done
}

x.backend.sync-deps() {
  go mod tidy
  go mod vendor
}

# Run project in production mode, i.e. dockerized and bundled.
x.run-dockerized() {
  # project run-dockerized --x-env=prod --x-no-gen
  run_shared_services up -d --build --wait
  x.db.recreate

  x.gen

  setup_swagger_ui

  DOCKER_BUILDKIT=1 BUILDKIT_PROGRESS=plain docker compose \
    --project-name "$PROJECT_PREFIX"_"$X_ENV" \
    -f docker-compose.yml \
    --env-file ".env.$X_ENV" \
    up -d --build --wait --force-recreate 2>&1 # https://github.com/docker/compose/issues/7346

  x.migrate up
}

# Remove running project containers, including shared ones between environments.
x.stop-project() {
  { { {
    run_shared_services down --remove-orphans
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Recreates docker volumes for Postgres, Redis, etc. Unsaved data will be lost.
x.recreate-shared-services() {
  run_shared_services up -d --build --force-recreate --wait
}

# Checks before release:
# - Magic keyword "STOPSHIP" not found in tracked files.
x.release() {
  { { {
    search_stopship "STOPSHIP" &
    GOWORK=off go mod verify & # (https://github.com/golang/go/issues/54372)

    wait_without_error
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Shows existing user api keys.
x.dev-utils.api-keys() {
  dockerdb psql --no-psqlrc -d $POSTGRES_DB -c "select email, api_key from user_api_keys left join users using (user_id);"
}

# Shows current database column comments used in xo codegen
# to ease further updates to comments in migrations.
x.dev-utils.show-column-comments() {
  local query="SELECT DISTINCT
  c.relname as table,
  a.attname::varchar AS column,
  COALESCE(col_description(format('%s.%s', n.nspname, c.relname)::regclass::oid, isc.ordinal_position), '') as column_comment
FROM pg_attribute a
  JOIN ONLY pg_class c ON c.oid = a.attrelid
  JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
  INNER JOIN information_schema.columns as isc on c.relname = isc.table_name and isc.column_name = a.attname
  LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid
    AND a.attnum = ANY(ct.conkey)
    AND ct.contype = 'p'
  LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid
    AND ad.adnum = a.attnum
WHERE a.attisdropped = false
  AND n.nspname = 'public'
  AND (true OR a.attnum > 0)
  AND col_description(format('%s.%s', n.nspname, c.relname)::regclass::oid, isc.ordinal_position) is not null;"

  dockerdb psql --no-psqlrc -d $POSTGRES_DB -c "$query" 2>/dev/null
}

# Setups a traefik container with predefined configuration in `install-dir`.
# Args: install-dir
x.setup.traefik() {
  { { {
    test -z "$1" && err "installation directory is required"

    x.setup.mkcert

    git clone --depth=1 https://github.com/danicc097/traefik-bootstrap.git "$1"
    docker network create traefik-net || true
    mkdir -p "$1"/traefik/certificates
    cp $CERTIFICATES_DIR/* "$1"/traefik/certificates
    cd "$1" || exit
    cp traefik/dynamic_conf.yaml.example traefik/dynamic_conf.yaml
    echo "Adding $PWD/certificates/"
    yq e ".tls.certificates += [{
    \"certFile\": \"$PWD/$CERTIFICATES_DIR/localhost.pem\",
    \"keyFile\": \"$PWD/$CERTIFICATES_DIR/localhost-key.pem\"
  }]" -i traefik/dynamic_conf.yaml

    ./compose-up
    cd - >/dev/null || exit
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Installs mkcert local development certificates.
x.setup.mkcert() {
  { { {
    cd "$CERTIFICATES_DIR" || exit
    echo "Setting up local certificates"
    mkcert --cert-file localhost.pem --key-file localhost-key.pem "localhost" "*.e2e.localhost" "*.local.localhost" "*.dev.localhost" "*.ci.localhost" "*.prod.localhost" "127.0.0.1" "::1" "host.docker.internal" 2>&1
    cd -
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

########################## migrations ##########################

# Wrapper for golang-migrate with predefined configuration.
x.migrate() {
  { { {
    migrate \
      -path $MIGRATIONS_DIR/ \
      -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:$EXPOSED_POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" \
      "$@" 2>&1

    if [[ "${*:1}" =~ (up)+ ]]; then
      echo "Running post-migration scripts"
      for file in $(find db/post-migration -maxdepth 1 -name '*.sql' | sort); do
        dockerdb_psql -U postgres -d "$POSTGRES_DB" <$file
      done
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Create a new migration file with the given `name`.
# Args: name
x.migrate.create() {
  { { {
    tmp="$*"
    tmp="${tmp// /_}"
    name="${tmp,,}"
    [[ -z $name ]] && err "Please provide a migration name"
    x.migrate create -ext sql -dir $MIGRATIONS_DIR/ -seq -digits 7 "$name"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

########################## db ##########################

# psql session for the current environment.
x.db.psql() {
  x.db.psql-db $POSTGRES_DB
}

# psql session for `database`.
# Args: database
x.db.psql-db() {
  docker exec -it postgres_db_"$PROJECT_PREFIX" psql -d $1
}

# Show active and max number of connections for the current environment.
x.db.conns() {
  { { {
    x.db.conns-db $POSTGRES_DB
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Show active and max number of connections for `database`.
# Args: database
x.db.conns-db() {
  { { {
    current_conns=$(dockerdb_psql -d $1 -c "SELECT count(*) FROM pg_stat_activity WHERE datname = '$1';")
    max_conns=$(dockerdb_psql -d $1 -c "SHOW max_connections;")
    echo "$current_conns/$max_conns active connections in '$1'"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}
# Create a new database in the current environment if it doesn't exist
# and stops its running processes if any.
x.db.recreate() {
  { { {
    create_db_if_not_exists $POSTGRES_DB
    stop_db_processes $POSTGRES_DB
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Drop and recreate the database in the current environment.
x.db.drop() {
  [[ $x_env = "prod" && "$POSTGRES_DB" != "$GEN_POSTGRES_DB" ]] && with_tty confirm "This will drop production database data. Continue?"
  { { {
    drop_and_recreate_db "$POSTGRES_DB"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Drop and recreate the database used for code generation up to N, N-1 revisions or none.
# Args: {up|up-1|drop}
x.db.gen() {
  { { {
    latest_rev=$(find $MIGRATIONS_DIR/*.sql -maxdepth 0 | wc -l)
    second_latest_rev=$(((latest_rev - 2) / 2)) # up+down
    POSTGRES_DB=$GEN_POSTGRES_DB x.db.drop
    # TODO should be able to autocomplete x function nested cases
    # if it's called as $CMD
    case $1 in
    up)
      POSTGRES_DB=$GEN_POSTGRES_DB x.migrate up
      ;;
    up-1)
      POSTGRES_DB=$GEN_POSTGRES_DB x.migrate up $second_latest_rev
      ;;
    drop) ;;
    *)
      err "Valid options {up|up-1|drop}, got: $1"
      ;;
    esac
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Seed gen database.
x.db.gen.initial-data() {
  { { {
    POSTGRES_DB=$GEN_POSTGRES_DB x.db.initial-data
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Seed database.
x.db.initial-data() {
  { { {
    x.db.drop
    x.migrate up
    echo "Loading initial data to $POSTGRES_DB"
    # dockerdb_psql -d $POSTGRES_DB <"./db/initial_data_$x_env.pgsql"
    go run cmd/initial-data/main.go -env .env.$X_ENV
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Backup the database for the current environment.
x.db.dump() {
  { { {
    local dump_prefix="dump_${X_ENV}_"
    running_dumps=$(dockerdb_psql -P pager=off -U postgres -d "postgres_$x_env" \
      -c "SELECT pid FROM pg_stat_activity WHERE application_name = 'pg_dump';")
    if [[ "$running_dumps" != "" ]]; then
      err "pg_dump is already running, aborting new dump"
    fi

    mkdir -p "$DUMPS_DIR"
    schema_v=$(dockerdb_psql -P pager=off -U postgres -d "postgres_$x_env" \
      -c "SELECT version FROM schema_migrations;")
    dump_file="${dump_prefix}$(date +%Y-%m-%dT%H-%M-%S)_version${schema_v}.gz"

    echo "Dumping database to $dump_file"
    dockerdb pg_dump -U postgres -d "postgres_$x_env" |
      gzip >"$DUMPS_DIR/$dump_file"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

# Restore the database with the latest dump or `file` for the current environment.
# Args: [file]
x.db.restore() {
  dump_file="$1"
  local dump_prefix="dump_${X_ENV}_"
  if [[ -n $dump_file ]]; then
    [[ ! -f $dump_file ]] && err "$dump_file does not exist"
    [[ "$dump_file" != *"$dump_prefix"* ]] && confirm "${RED}Dump doesn't match prefix '$dump_prefix'. Continue?${OFF}"
  else
    mkdir -p "$DUMPS_DIR"
    latest_dump_file=$(find "$DUMPS_DIR"/ -name "$dump_prefix*.gz" | sort -r | head -n 1)
    if [[ -z "$latest_dump_file" ]]; then
      err "No $dump_prefix* file found in $DUMPS_DIR"
    fi
    dump_file="$latest_dump_file"
  fi

  confirm "Do you want to restore ${YELLOW}$dump_file${OFF} in the ${RED}$x_env${OFF} environment?"

  x.db.drop
  gunzip -c "$dump_file" | dockerdb_psql -U postgres -d "postgres_$x_env"
  # sanity check, but probably better to do it before restoring...
  dump_schema_v=$(dockerdb_psql -P pager=off -U postgres -d "postgres_$x_env" -c "SELECT version FROM schema_migrations;")
  file_schema_v=$(echo "$dump_file" | sed -E 's/.*_version([0-9]+)\..*/\1/')
  echo "Migration revision: $dump_schema_v"
  if [[ "$dump_schema_v" != "$file_schema_v" ]]; then
    err "Schema version mismatch: dump $dump_schema_v != file $file_schema_v"
  fi
}

########################## e2e ##########################

# Run E2E tests.
x.e2e.run() {
  { { {
    source .env.e2e

    x.gen.e2e

    name="$PROJECT_PREFIX-e2e"
    cd e2e
    DOCKER_BUILDKIT=1 BUILDKIT_PROGRESS=plain docker build -t "$name" .
    cd - >/dev/null

    # need symlink resolution for data

    test -t 0 && opts="-t"
    docker run -i $opts --rm \
      --ipc=host \
      --network host \
      -v "$(pwd)/cmd/oidc-server/data/:/cmd/oidc-server/data/" \
      -v "$(pwd)/e2e:/e2e/" \
      "$name" \
      bash -c "playwright test"
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

########################## openapi ##########################

# Run a diff against the previous OpenAPI spec in the main branch.
# Can also be used to generate changelogs when upgrading major versions.
x.diff-oas() {
  { { {
    base_spec="/tmp/openapi.yaml"
    git show "main:$SPEC_PATH" >"$base_spec"

    tmp="$(yq .info.version "$base_spec")"
    base_v="${tmp%%.*}"
    tmp=$(yq .info.version "$SPEC_PATH")
    rev_v="${tmp%%.*}"
    ((rev_v != base_v)) &&
      echo "${YELLOW}Revision mismatch $rev_v and $base_v, skipping diff.${OFF}" && return

    args="-format text -breaking-only -fail-on-diff -exclude-description -exclude-examples"
    if oasdiff $args -base "$base_spec" -revision $SPEC_PATH; then
      echo "${GREEN}No breaking changes found in $SPEC_PATH${OFF}"
    else
      echo "${RED}Breaking changes found in $SPEC_PATH${OFF}"
      return 1
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

########################### helpers ###########################

# IMPORTANT: bug in declare -F returns line number of last nested function, if any.
# extracting function here instead...
run_hot_reload() {
  x.run.backend-hr &
  pids="$pids $!"
  x.run.frontend &
  pids="$pids $!"
}

run_shared_services() {
  docker network create traefik-net 2>/dev/null || true

  local extra_services
  if [[ $x_env != "prod" ]]; then
    extra_services="-f docker-compose.oidc.yml"
  fi
  cd docker
  DOCKER_BUILDKIT=1 BUILDKIT_PROGRESS=plain docker compose \
    -p "$PROJECT_PREFIX" \
    -f docker-compose.shared.yml \
    $extra_services \
    --env-file ../.env."$X_ENV" \
    "$@" 2>&1 # https://github.com/docker/compose/issues/7346
  cd - >/dev/null
}

# generate db structs for use with swaggest/openapi-go.
# no need for ast parsing since all code is predictable
# NOTE: type grouping not supported.
generate_structs_map() {
  local structs=()
  local enums=()
  find_go_structs structs "$PG_REPO_GEN/db"
  find_go_enums enums "$PG_REPO_GEN/db"
  find_deleted_pkg_schemas structs enums Db
  for struct in ${structs[@]}; do
    map_fields+=("\"Db$struct\": new(db.$struct),") # swaggest requires pointer to struct
  done
  map_fields+=("
  //
  ")

  structs=()
  enums=()
  find_go_structs structs "internal/rest/models.go"
  find_go_enums enums "internal/rest/models.go"
  find_deleted_pkg_schemas structs enums Rest
  for struct in ${structs[@]}; do
    map_fields+=("\"Rest$struct\": new(rest.$struct),")
  done
  map_fields+=("
  //
  ")

  out="internal/codegen/structs.gen.go"
  cat <<EOF >$out
// Code generated by project. DO NOT EDIT.

package codegen

import (
  db "$GOMOD_PKG/$PG_REPO_GEN/db"
	rest "$GOMOD_PKG/internal/rest"
  )
    var PublicStructs = map[string]any{
$(printf "%s\n" "${map_fields[@]}")
  }
EOF

  gofumpt -w $out
}

sync_db_enums_with_spec() {
  db_search_paths=("public" "other_schema")

  search_path_str=$(printf "'%s'," "${db_search_paths[@]}")
  search_path_str=${search_path_str%,}
  enum_query=$(printf "
    SELECT t.typname AS enum_name
    FROM pg_type t
    INNER JOIN pg_namespace n ON n.oid = t.typnamespace
    WHERE t.typtype = 'e' AND n.nspname IN (%s);" "$search_path_str")

  mapfile -t db_enum_names < <(dockerdb_psql -d $POSTGRES_DB -c "$enum_query" 2>/dev/null)
  for db_enum_name in "${db_enum_names[@]}"; do
    local schema_name="$(to_pascal $db_enum_name)"
    local enum_values=()
    mapfile -t enum_values < <(dockerdb_psql -d $POSTGRES_DB -c "SELECT unnest(enum_range(NULL::\"$db_enum_name\"));" 2>/dev/null)

    local schema_path=".components.schemas.$schema_name"
    if yq -e "$schema_path" $SPEC_PATH &>/dev/null; then # -e exits 1 if no match
      if ! yq -e "$schema_path | has(\"x-generated\")" $SPEC_PATH &>/dev/null; then
        err "Clashing schema name '$schema_name'. Please remove it before continuing."
      fi
    fi

    src_comment="database enum '$db_enum_name'"
    replace_enum_in_spec "$schema_name" enum_values "$src_comment"

  done

}

setup_swagger_ui() {
  go run cmd/swagger-ui-setup/main.go -env=".env.$X_ENV" -swagger-ui-dir=$SWAGGER_UI_DIR

  local bundle_spec="$SWAGGER_UI_DIR/openapi.yaml"
  cp $SPEC_PATH $bundle_spec
  yq 'explode(.)' -i $bundle_spec
  sed -i 's/!!merge //' $bundle_spec
}

replace_enum_in_spec() {
  local enum="$1"
  local -n __arr="$2"
  local src_comment="$3"

  local schema_path=".components.schemas.$enum"

  local __enums
  __enums=$(printf ",\"%s\"" "${__arr[@]}")
  __enums="[${__enums:1}]"

  echo "Replacing '$enum' enum in $SPEC_PATH with values from $src_comment"
  __enums=$__enums yq e "
    $schema_path.type = \"string\" |
    $schema_path.enum = env(__enums) |
    $schema_path.x-generated = \"-\" |
    ($schema_path | key) line_comment=\"Generated from $src_comment. DO NOT EDIT.\"" -i $SPEC_PATH
}

awk_remove_go_comments='
     /\/\*/ { f=1 } # set flag that is a block comment

     f==0 && !/^\s*(\/\/|\/\*)/ {
        print  # print non-commented lines
     }
     /\*\// { f=0 } # reset flag at end of comment
'

find_go_structs() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/[\s]*type\(.*\)struct.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No structs found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

find_go_interfaces() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/[\s]*type\(.*\)interface.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    err "No interfaces found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

find_go_interface_content() {
  local interface="$1"
  local file="$2"
  awk "/^type $interface /{flag=1; print; next} flag && /^}/{flag=0} flag" $file |
    sed -e '1d' |
    awk "$awk_remove_go_comments"
}

find_go_enums() {
  local -n __arr="$1"
  local pkg="$2"
  mapfile -t __arr < <(find $pkg -maxdepth 1 -name "*.go" -exec awk "$awk_remove_go_comments" {} \; |
    sed -ne 's/.*type[[:space:]]\+\([^=[:space:]]\+\)[[:space:]]\+string.*/\1/p')
  if [[ ${#__arr[@]} -eq 0 ]]; then
    echo "No enums found in package $pkg"
  fi
  mapfile -t __arr < <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

find_deleted_pkg_schemas() {
  local -n __structs="$1"
  local -n __enums="$2"
  local pkg="$3"
  local pkg_prefix=$(to_pascal $pkg)
  echo "Finding deleted structs or enums from package '$pkg_prefix'..."
  mapfile -t spec_schemas < <(yq eval '.components.schemas[] | key' "$SPEC_PATH" | grep -E "^${pkg_prefix}" || true)

  local found=0
  for spec_schema in ${spec_schemas[*]}; do
    for struct in ${__structs[@]}; do
      if [[ $spec_schema == "${pkg_prefix}$struct" ]]; then
        found=1
      fi
    done
    for enum in ${__enums[@]}; do
      if [[ $spec_schema == "${pkg_prefix}$enum" ]]; then
        found=1
      fi
    done
    ((found == 0)) && echo "${YELLOW}[WARNING] $SPEC_PATH schema $spec_schema no longer exists in package '$pkg'. Remove if necessary.${OFF}"
    found=0
  done
}

update_spec_with_structs() {
  vext="x-postgen-struct"
  struct_names=$(yq e ".components.schemas[] | select(has(\"$vext\")).$vext" $SPEC_PATH)
  schema_names=$(yq e ".components.schemas[] | select(has(\"$vext\")) | key" $SPEC_PATH)
  mapfile -t struct_names <<<$struct_names
  mapfile -t schema_names <<<$schema_names

  # openapi-go will generate a RestXYZ if we have a XYZ: <...> x-postgen-struct: RestXYZ
  # we need to detect these early, because it will sync a RestXYZ instead of XYZ
  declare -A schemas
  for i in ${!struct_names[@]}; do
    schemas["${struct_names[$i]}"]="${schema_names[$i]}" # keep track of custom structs per schema name
  done

  struct_names_list=$(join_by "," ${struct_names[*]})
  ((${#struct_names_list} == 0)) && return

  # NOTE: maybe https://github.com/pkujhd/goloader allows reloading packages at runtime
  # otherwise maybe yaegi can be used for openapi-go without much fuzz
  # Only implement the above if there are cases where we cant have compilable state when
  # building gen-schema at this step, or the workaround is too tedious

  local gen_schema_spec="/tmp/openapi.yaml"
  # always compile and run since we need new PublicStructs that were just changed
  codegen gen-schema --struct-names $struct_names_list | yq '
    with_entries(select(.key == "components"))' \
    >$gen_schema_spec || [[ -n "$X_IGNORE_BUILD_ERRORS" ]]

  # replace every schema back into the spec
  for schema in $(yq '.components.schemas[] | key' $gen_schema_spec); do
    schema_name="${schemas[$schema]:-$schema}"

    yq eval-all "(
        select(fi == 1).components.schemas.$schema
        ) as \$schema
        | select(fi == 0)
        | .components.schemas.$schema_name = \$schema
        | (.components.schemas.$schema_name | key) line_comment=\"Generated from internal structs. DO NOT EDIT.\"
      " "$SPEC_PATH" $gen_schema_spec >/tmp/final-spec

    if [[ "$schema_name" != "$schema" ]]; then
      echo "Deleting schema \"$schema\". Replace references with \"$schema_name\""
      yq e "del(.components.schemas.$schema)" -i /tmp/final-spec
    fi

    mv /tmp/final-spec "$SPEC_PATH" # need to update at each iteration since next depends on it
  done
}

remove_schemas_marked_to_delete() {
  local paths_arr paths
  paths_arr=$(yq e '..
      | select(has("x-TO-BE-DELETED"))
      | path
      | with(.[] | select(contains(".") or contains("/") or contains("{")); . = "\"" + . + "\"")
      | join(".")
      | . = "." + .
    ' $SPEC_PATH)

  paths=$(join_by "," "${paths_arr[@]}")
  yq e "del($paths)" -i "$SPEC_PATH"
}

xo_schema() {
  # xo cannot use db files as input, needs an up-to-date schema
  # not recreating db on every gen can lead to plain wrong generation based on an old dev schema.
  # Also use a unique db to prevent cosmic accidents
  xo schema "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:$EXPOSED_POSTGRES_PORT/$GEN_POSTGRES_DB?sslmode=disable" \
    --src "$XO_TEMPLATES_DIR" \
    "$@"
}

# updates dynamic config with env vars for
config_template_setup() {
  local dir="$1"
  export ENV_REPLACE_GLOB=$dir/config.json
  # ensure config has all k:v as "<KEY>": "$<KEY>"
  # this has to always be run at startup
  jq \
    'to_entries | map_values({ (.key) : ("$" + .key) }) | reduce .[] as $item ({}; . + $item)' \
    $dir/config.template.json >/tmp/$dir-config.tmp.json && mv /tmp/$dir-config.tmp.json "$ENV_REPLACE_GLOB"
  envvars=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,)
  frontend/nginx/replace-envvars.sh "$envvars"
}
