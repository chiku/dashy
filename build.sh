#/bin/bash

set -euo pipefail
IFS=$'\n\t'

make clean
npm run clean

make test
npm run test

make compile
npm run compile
