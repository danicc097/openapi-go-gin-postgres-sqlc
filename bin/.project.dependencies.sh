#!/usr/bin/env bash

source ".helpers.sh"

report_failure() {
  local info="$1"
  echo "
${RED}Failed ${FUNCNAME[1]##*.} check.${OFF}
Minimum version: $minver
Current version: $vers
$info
"
  exit 1
}

report_success() {
  printf "%s ✅\n" "${FUNCNAME[1]##*.}: $minver"
  exit 0
}

check.bin.bash() {
  { { {
    vers=${BASH_VERSION:0:1}
    minver=5
    if ((vers >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.column() {
  { { {
    vers=$(column --version)
    minver="util-linux"
    if [[ "$vers" = *$minver* ]]; then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.column() {
  { { {
    wget --tries=10 https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.36/util-linux-2.36.2.tar.gz
    tar -xf util-linux-2.36.2.tar.gz
    cd util-linux-2.36.2 || exit 1
    ./configure
    make column
    cp .libs/column ../bin/tools/
    cd ..
    rm -rf util-linux-2.*
    column --version
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.protoc() {
  { { {
    vers=$(protoc --version)
    minver="libprotoc 3"
    if [[ "$vers" = *$minver* ]]; then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.protoc() {
  { { {
    VERSION=3.19.4
    id="protoc-$VERSION-linux-x86_64"
    wget --tries=10 https://github.com/protocolbuffers/protobuf/releases/download/v"$VERSION"/$id.zip
    unzip -q "$id".zip -d "$id"
    mv "$id"/bin/protoc ./bin/tools/
    rm -rf "$id"
    rm -f "$id".zip
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.curl() {
  { { {
    local -a vers_arr
    mapfile vers_arr < <(curl -V 2>&1)
    minver="libcurl"
    vers="${vers_arr[0]}"
    if [[ $vers = *$minver* ]]; then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.go() {
  { { {
    vers=$(go version)
    minver=18
    if [[ "$vers" =~ ^[^\ ]+\ [^\ ]+\ go1\.([^\ \.]+) ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.docker() {
  { { {
    vers=$(docker --version)
    minver=2
    if [[ "$vers" =~ version[\ ]+([^\ \.]+) ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.docker-compose() {
  { { {
    vers=$(docker compose version)
    minver=2
    if [[ "$vers" =~ [\ ]+[v]?([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.pg_format() {
  { { {
    vers=$(pg_format --version)
    minver=5
    if [[ "$vers" =~ [\ ]+([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.pg_format() {
  { { {
    local version="5.2"
    wget --tries=10 https://github.com/darold/pgFormatter/archive/refs/tags/v"$version".tar.gz
    tar xzf v"$version".tar.gz
    cd pgFormatter-"$version"/ || exit 1
    perl Makefile.PL
    sudo make install
    cd ..
    sudo rm -rf pgFormatter-"$version"/
    sudo rm v"$version".tar.gz
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.ripgrep() {
  { { {
    vers=$(rg --version)
    minver=11
    if [[ "$vers" =~ (ripgrep )([0-9]+) ]] && ((BASH_REMATCH[2] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.ripgrep() {
  { { {
    sudo apt-get install ripgrep -y
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.parallel() {
  { { {
    vers=$(parallel --version)
    if [[ "$vers" =~ (GNU parallel )([0-9]+) ]]; then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.parallel() {
  { { {
    sudo apt-get install parallel -y
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.direnv() {
  { { {
    vers=$(direnv --version)
    minver=2
    if [[ "$vers" =~ ([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.direnv() {
  { { {
    sudo apt-get install direnv -y
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.pnpm() {
  { { {
    vers=$(pnpm --version)k
    minver=8
    if [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.pnpm() {
  { { {
    npm install --prefix ./bin/tools/ -g pnpm@8.3.1
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.node() {
  { { {
    vers=$(node --version)
    minver=16
    if [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.node() {
  { { {
    sudo npm i -g node@7.6.0
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.sponge() {
  { { {
    vers=$(command -v sponge)
    minver="-"
    if [[ $vers = */sponge* ]]; then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.sponge() {
  { { {
    sudo apt-get install moreutils -y
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

check.bin.mkcert() {
  { { {
    vers=$(mkcert --version)
    minver=1
    if [[ "$vers" =~ [v]?([0-9]+)[\.]{1} ]] && ((BASH_REMATCH[1] >= minver)); then
      report_success
    else
      report_failure "$@"
    fi
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}

install.bin.mkcert() {
  { { {
    VERSION="1.4.3"
    sudo apt-get install libnss3-tools -y
    wget --tries=10 https://github.com/FiloSottile/mkcert/releases/download/v"$VERSION"/mkcert-v"$VERSION"-linux-amd64 -O mkcert
    chmod +x mkcert
    mv mkcert ./bin/tools/
    mkcert -install
  } 2>&4 | xlog >&3; } 4>&1 | xerr >&3; } 3>&1
}
