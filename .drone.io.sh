#/bin/sh

install_prerequisites() {
  sudo apt-get install -y nodejs golang
}

gulp() {
  ./node_modules/.bin/gulp
}

run() {
  commnd=$1
  echo "--- Start $commnd ---"
  "$commnd"
  echo "--- End $commnd ---"
}


set -e

run install_prerequisites
run gulp
