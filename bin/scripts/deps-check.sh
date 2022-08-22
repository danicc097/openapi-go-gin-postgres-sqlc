#!/bin/bash

check.build-deps() {
  local -i fails
  check.bash || (echo "${RED}Failed bash check${OFF}" && ((fails++)))
  check.go || (echo "${RED}Failed go check${OFF}" && ((fails++)))
  check.java || (echo "${RED}Failed java check${OFF}" && ((fails++)))
  check.curl || (echo "${RED}Failed curl check${OFF}" && ((fails++)))
  check.docker || (echo "${RED}Failed docker check${OFF}" && ((fails++)))
  check.docker-compose || (echo "${RED}Failed docker-compose check${OFF}" && ((fails++)))
  check.direnv || (echo "${RED}Failed direnv check${OFF}" && ((fails++)))
  check.yq || (echo "${RED}Failed yq check${OFF}" && ((fails++)))
  ((fails == 0)) && echo "${GREEN}🎉 All build dependencies met${OFF}"
}

check.bash() {
  ((${BASH_VERSION:0:1} >= 4)) &&
    printf "%-30s ✅\n" "Bash: ${BASH_VERSION:0:1}"
}

check.curl() {
  local -a versa
  mapfile versa < <(curl -V 2>&1)
  if [[ "${versa[0]}" =~ ^[^\ ]+\ ([^\ ]+) ]]; then
    printf "%-30s ✅\n" "Curl: ${BASH_REMATCH[1]}"
    return 0
  else
    return 1
  fi
}

check.go() {
  local vers
  vers=$(go version)
  [[ "$vers" =~ ^[^\ ]+\ [^\ ]+\ go1\.([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= 18)) &&
    printf "%-30s ✅\n" "Go: ${BASH_REMATCH[1]}"
}

check.yq() {
  local vers
  vers=$(yq --version)
  [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= 4)) &&
    printf "%-30s ✅\n" "yq: ${BASH_REMATCH[1]}"
}

check.docker() {
  local vers
  vers=$(docker --version)
  [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] &&
    ((BASH_REMATCH[1] >= 20)) &&
    printf "%-30s ✅\n" "Docker: ${BASH_REMATCH[1]}"
}

check.docker-compose() {
  local vers
  vers=$(docker-compose version)
  [[ "$vers" =~ [\ ]+v([0-9]+)[\.]{1} ]] &&
    ((BASH_REMATCH[1] >= 2)) &&
    printf "%-30s ✅\n" "Docker Compose: ${BASH_REMATCH[1]}"
}

check.direnv() {
  local vers
  vers=$(direnv --version)
  [[ "$vers" =~ ([0-9]+)[\.]{1} ]] &&
    ((BASH_REMATCH[1] >= 2)) &&
    printf "%-30s ✅\n" "direnv: ${BASH_REMATCH[1]}"
}

check.java() {
  local jvers
  jvers=$(java -version 2>&1 | head -1)
  jvers=${jvers#*version \"}
  jvers=${jvers%%\"*}
  if [[ 
    $jvers =~ ^(1\.[89]|9\.|[1-9][0-9]+) ]]; then
    printf "%-30s ✅\n" "Java: $jvers"
    return 0
  else
    return 1
  fi
}
