#!/bin/bash

check.bash() {
  minver=4
  { ((${BASH_VERSION:0:1} >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_VERSION:0:1}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.column() {
  local vers
  vers=$(column --version)
  minver="util-linux"
  { [[ "$vers" = *$minver* ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check.${OFF}"
      echo "${YELLOW}Run install-column.sh to get the $minver version.${OFF}"
      return 1
    }
}

check.protoc() {
  local vers
  vers=$(protoc --version)
  minver="libprotoc 3"
  { [[ "$vers" = *$minver* ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check.${OFF}"
      echo "${YELLOW}Run install-protoc.sh to get $minver.${OFF}"
      return 1
    }
}

check.curl() {
  local -a versa
  mapfile versa < <(curl -V 2>&1)
  { [[ "${versa[0]}" =~ ^[^\ ]+\ ([^\ ]+) ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.go() {
  local vers
  vers=$(go version)
  minver=18
  { [[ "$vers" =~ ^[^\ ]+\ [^\ ]+\ go1\.([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.yq() {
  local vers
  vers=$(yq --version)
  minver=4
  { [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] && [[ "$vers" = *mikefarah/yq* ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver of https://github.com/mikefarah/yq/)${OFF}"
      return 1
    }
}

check.docker() {
  local vers
  vers=$(docker --version)
  minver=2
  { [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.docker-compose() {
  local vers
  vers=$(docker-compose version)
  minver=2
  {
    [[ "$vers" =~ [\ ]+v([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
  } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.pg_format() {
  local vers
  vers=$(pg_format --version)
  minver=5
  {
    [[ "$vers" =~ [\ ]+([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
  } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.direnv() {
  local vers
  vers=$(direnv --version)
  minver=2
  { [[ "$vers" =~ ([0-9]+)[\.]{1} ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}
