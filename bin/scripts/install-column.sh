#!/bin/bash

wget https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.36/util-linux-2.36.2.tar.gz
tar -xf util-linux-2.36.2.tar.gz
cd util-linux-2.36.2 || exit 1
./configure
make column
mkdir -p ~/bin
cp .libs/column ~/bin/
cd ..
rm -rf util-linux-2.*
column --version
echo "Ensure your ~/bin folder is in PATH."
