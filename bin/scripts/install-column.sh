#!/bin/bash

wget https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.35/util-linux-2.35.1.tar.gz
tar -xf util-linux-2.35.1.tar.gz
cd util-linux-2.35.1 || exit 1
./configure
make column
mkdir -p ~/bin
cp .libs/column ~/bin/
cd ..
rm -rf util-linux-2.*
column --version
echo "Ensure your ~/bin folder is in PATH."
