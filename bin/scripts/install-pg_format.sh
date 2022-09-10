#!/bin/bash

VERSION="5.2"
wget https://github.com/darold/pgFormatter/archive/refs/tags/v"$VERSION".tar.gz
tar xzf v"$VERSION".tar.gz
cd pgFormatter-"$VERSION"/ || exit 1
perl Makefile.PL
sudo make install
cd ..
rm -rf pgFormatter-"$VERSION"/
rm v"$VERSION".tar.gz
