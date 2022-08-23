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
  { [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
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

check.java() {
  local vers
  vers=$(java -version 2>&1 | head -1)
  vers=${vers#*version \"}
  vers=${vers%%\"*}
  { [[ $vers =~ ^(1\.[89]|9\.|[1-9][0-9]+) ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}
