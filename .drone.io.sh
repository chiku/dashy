#/bin/bash

set -e

setup_environment() {
  export GOROOT=$HOME/go
  export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
}

install_os_packages() {
  sudo apt-get update
  sudo apt-get remove --purge nodejs npm
  sudo apt-get install -y curl wget
  curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
  sudo apt-get update
  sudo apt-get install -y nodejs
}

install_golang() {
  mkdir -p $GOROOT
  pushd $HOME
  file=go1.5.2.linux-amd64.tar.gz
  wget --continue "https://storage.googleapis.com/golang/$file" -O "$HOME/$file"
  tar -zxvf $file
  popd
}

install_node_deps() {
  npm install
}

install_golang_deps() {
  go get github.com/tools/godep
}

build() {
  ./node_modules/.bin/gulp
}

run() {
  commnd=$1
  echo "--- Start $commnd ---"
  "$commnd"
  echo "--- End $commnd ---"
}

main() {
  run setup_environment
  run install_os_packages
  run install_golang
  run install_node_deps
  run install_golang_deps
  run build
}

main
