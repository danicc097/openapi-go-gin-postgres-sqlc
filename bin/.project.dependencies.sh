#!/bin/bash

source "${BASH_SOURCE%/*}/.helpers.sh"

check.bin.bash() {
  { { {
    vers=${BASH_VERSION:0:1}
    minver=4
    { ((vers >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        exit 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.column() {
  { { {
    local vers
    vers=$(column --version)
    minver="util-linux"
    { [[ "$vers" = *$minver* ]] &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.column() {
  { { {
    wget https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.36/util-linux-2.36.2.tar.gz
    tar -xf util-linux-2.36.2.tar.gz
    cd util-linux-2.36.2 || exit 1
    ./configure
    make column
    cp .libs/column ./bin/tools/
    cd ..
    rm -rf util-linux-2.*
    column --version
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.protoc() {
  { { {
    local vers
    vers=$(protoc --version)
    minver="libprotoc 3"
    { [[ "$vers" = *$minver* ]] &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: $minver"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.protoc() {
  { { {
    VERSION=3.19.4
    id="protoc-$VERSION-linux-x86_64"
    wget https://github.com/protocolbuffers/protobuf/releases/download/v"$VERSION"/$id.zip
    unzip -q "$id".zip -d "$id"
    mv "$id"/bin/protoc ./bin/tools/
    rm -rf "$id"
    rm -f "$id".zip
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.curl() {
  { { {
    local -a versa
    mapfile versa < <(curl -V 2>&1)
    { [[ "${versa[0]}" =~ ^.*(libcurl).* ]] &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. libcurl Required"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.go() {
  { { {
    local vers
    vers=$(go version)
    minver=18
    { [[ "$vers" =~ ^[^\ ]+\ [^\ ]+\ go1\.([^\ \.]+) ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.docker() {
  { { {
    local vers
    vers=$(docker --version)
    minver=2
    { [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.docker-compose() {
  { { {
    local vers
    vers=$(docker compose version)
    minver=2
    {
      [[ "$vers" =~ [\ ]+[v]?([0-9]+)[\.]{1} ]] &&
        ((BASH_REMATCH[1] >= minver)) &&
        printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
    } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.pg_format() {
  { { {
    local vers
    vers=$(pg_format --version)
    minver=5
    {
      [[ "$vers" =~ [\ ]+([0-9]+)[\.]{1} ]] &&
        ((BASH_REMATCH[1] >= minver)) &&
        printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
    } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.pg_format() {
  { { {
    sudo apt-get install pgformatter
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.parallel() {
  { { {
    local vers
    vers=$(parallel --version)
    {
      [[ "$vers" =~ (GNU parallel )([0-9]+) ]] &&
        printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[2]}"
    } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. Check install-parallel.sh"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.parallel() {
  { { {
    sudo apt-get install parallel
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.direnv() {
  { { {
    local vers
    vers=$(direnv --version)
    minver=2
    { [[ "$vers" =~ ([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.direnv() {
  { { {
    sudo apt-get install direnv
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.pnpm() {
  { { {
    local vers
    vers=$(pnpm --version)
    minver=7
    { [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.pnpm() {
  { { {
    sudo npm i -g pnpm@7.6.0
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.node() {
  { { {
    local vers
    vers=$(node --version)
    minver=16
    { [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] &&
      ((BASH_REMATCH[1] >= minver)) &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.node() {
  { { {
    sudo npm i -g node@7.6.0
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.sponge() {
  { { {
    { [[ $(command -v sponge) =~ /sponge$ ]] &&
      printf "%s ✅\n" "${FUNCNAME[0]##*.}"; } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check."
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.sponge() {
  { { {
    sudo apt-get install moreutils
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.mkcert() {
  { { {
    local vers
    vers=$(mkcert --version)
    minver=1
    {
      [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] &&
        ((BASH_REMATCH[1] >= minver)) &&
        printf "%s ✅\n" "${FUNCNAME[0]##*.}: ${BASH_REMATCH[1]}"
    } ||
      {
        echo "${RED}Failed ${FUNCNAME[0]##*.} check. (minimum version: $minver)${OFF}"
        echo "Current version: $vers"
        return 1
      }
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.mkcert() {
  { { {
    VERSION="1.4.3"
    sudo apt-get install libnss3-tools
    wget https://github.com/FiloSottile/mkcert/releases/download/v"$VERSION"/mkcert-v"$VERSION"-linux-amd64 -O mkcert
    chmod +x mkcert
    mv mkcert ./bin/tools/
    cd $CERTIFICATES_DIR || exit
    echo "Setting up local certificates"
    mkcert --cert-file localhost.pem --key-file localhost-key.pem localhost "*.dev.localhost" "*.ci.localhost" "*.prod.localhost" 127.0.0.1 ::1 host.docker.internal
    cd ..
    mkcert -install
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}
