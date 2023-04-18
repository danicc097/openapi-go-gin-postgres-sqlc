#!/bin/bash

source "${BASH_SOURCE%/*}/.helpers.sh"

ensure_pwd_is_top_level

check.all() {
  while IFS= read -r line; do
    [[ $line =~ ^declare\ -f\ check\.bin\. ]] && BIN_CHECKS+=("${line##declare -f check.bin.}")
    [[ $line =~ ^declare\ -f\ install\.bin\. ]] && BIN_INSTALLS+=("${line##declare -f install.bin.}")
  done < <(declare -F)

  for bin in "${BIN_CHECKS[@]}"; do
    "check.bin.$bin" && continue

    if ! element_in_array "$bin" "${BIN_INSTALLS[@]}"; then
      exit 1
    fi

    confirm "Do you want to install $bin now?" || exit 1

    "install.bin.$bin" || err "$bin installation failed"
  done
}

check.bin.bash() {
  minver=4
  { ((${BASH_VERSION:0:1} >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_VERSION:0:1}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

check.bin.column() {
  local vers
  vers=$(column --version)
  minver="util-linux"
  { [[ "$vers" = *$minver* ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check.${OFF}"
      return 1
    }
}

install.bin.column() {
  wget https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.36/util-linux-2.36.2.tar.gz
  tar -xf util-linux-2.36.2.tar.gz
  cd util-linux-2.36.2 || exit 1
  ./configure
  make column
  cp .libs/column ./bin/tools/
  cd ..
  rm -rf util-linux-2.*
  column --version
}

check.bin.protoc() {
  local vers
  vers=$(protoc --version)
  minver="libprotoc 3"
  { [[ "$vers" = *$minver* ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check.${OFF}"
      return 1
    }
}

install.bin.protoc() {
  VERSION=3.19.4
  id="protoc-$VERSION-linux-x86_64"
  wget https://github.com/protocolbuffers/protobuf/releases/download/v"$VERSION"/$id.zip
  unzip -q "$id".zip -d "$id"
  mv "$id"/bin/protoc ./bin/tools/
  rm -rf "$id"
  rm -f "$id".zip
}

check.bin.curl() {
  local -a versa
  mapfile versa < <(curl -V 2>&1)
  { [[ "${versa[0]}" =~ ^.*(libcurl).* ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check."
      return 1
    }
}

check.bin.go() {
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

check.bin.yq() {
  local vers
  vers=$(yq --version)
  minver=4
  { [[ "$vers" =~ version[\ ]+[v]?([^\ \.]+) ]] && [[ "$vers" = *mikefarah/yq* ]] &&
    ((BASH_REMATCH[1] >= minver)) &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver of https://github.com/mikefarah/yq/)${OFF}"
      return 1
    }
}

check.bin.docker() {
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

check.bin.docker-compose() {
  local vers
  vers=$(docker compose version)
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

check.bin.pg_format() {
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

install.bin.pg_format() {
  sudo apt install pgformatter
}

check.bin.parallel() {
  local vers
  vers=$(parallel --version)
  {
    [[ "$vers" =~ (GNU parallel )([0-9]+) ]] &&
      printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[2]}"
  } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. Check install-parallel.sh"
      return 1
    }
}

install.bin.parallel() {
  sudo apt install parallel
}

check.bin.direnv() {
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

install.bin.direnv() {
  sudo apt install direnv
}

check.bin.sponge() {
  { [[ $(command -v sponge) =~ /sponge$ ]] &&
    printf "%-40s ✅\n" "${FUNCNAME[0]##*.}"; } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check."
      return 1
    }
}

install.bin.sponge() {
  sudo apt install sponge
}
check.bin.mkcert() {
  local vers
  vers=$(mkcert --version)
  minver=1
  {
    [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%-40s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
  } ||
    {
      echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver${OFF})"
      return 1
    }
}

install.bin.mkcert() {
  VERSION="1.4.3"
  sudo apt-get install libnss3-tools
  wget https://github.com/FiloSottile/mkcert/releases/download/v"$VERSION"/mkcert-v"$VERSION"-linux-amd64 -O mkcert
  chmod +x mkcert
  mv mkcert bin/tools/
  cd certificates || exit
  echo "Setting up local certificates"
  mkcert --cert-file localhost.pem --key-file localhost-key.pem localhost "*.dev.localhost" "*.ci.localhost" "*.prod.localhost" 127.0.0.1 ::1 host.docker.internal
  cd ..
  mkcert -install
}
