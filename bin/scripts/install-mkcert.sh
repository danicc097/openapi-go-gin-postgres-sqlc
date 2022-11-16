#!/bin/bash

VERSION="1.4.3"
sudo apt-get install libnss3-tools
wget https://github.com/FiloSottile/mkcert/releases/download/v"$VERSION"/mkcert-v"$VERSION"-linux-amd64 -O mkcert
chmod +x mkcert
sudo mv mkcert /usr/bin/
source ~/.bashrc
cd certificates || exit
mkcert --cert-file localhost.pem --key-file localhost-key.pem localhost "*.dev.localhost" "*.ci.localhost" "*.prod.localhost" 127.0.0.1 ::1 host.docker.internal
cd ..
mkcert -install
