#!/bin/bash

VERSION=3.19.4
wget https://github.com/protocolbuffers/protobuf/releases/download/v"$VERSION"/protoc-"$VERSION"-linux-x86_64.zip
unzip protoc-"$VERSION"-linux-x86_64.zip
mkdir -p ~/bin
mv bin/protoc ~/bin
