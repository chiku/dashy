#/bin/sh

set -e

setup_environment() {
  export PATH=$PATH:$GOPATH/bin
}

setup_os_prerequisites() {
  sudo apt-get install -y nodejs golang
}

setup_npm_prerequisites() {
  npm install
}

setup_godep() {
  go get -t github.com/tools/godep
}

setup_gulp() {
  ./node_modules/.bin/gulp
}

run() {
  commnd=$1
  echo "--- Start $commnd ---"
  "$commnd"
  echo "--- End $commnd ---"
}

main() {
  setup_environment
  run setup_os_prerequisites
  run setup_npm_prerequisites
  run setup_godep
  run setup_gulp
}

main
